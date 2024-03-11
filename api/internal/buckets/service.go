package buckets

import (
	"encoding/hex"
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/buckets/db/mdbx"
	"github.com/RevittConsulting/cdk-envs/internal/types"
	"github.com/RevittConsulting/cdk-envs/pkg/utils"
	"io/fs"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

type IDatabase interface {
	Open(path string) error
	Close() error
	ListBuckets() ([]string, error)
	CountKeys(bucketName string) (uint64, error)
	CountKeysOfLength(bucketName string, length uint64) (uint64, []string, error)
	FindByKey(bucketName string, key []byte) ([]byte, error)
	FindByValue(bucketName string, value []byte) ([][]byte, error)
	Read(bucketName string, take, offset uint64) ([]types.KeyValuePair, error)
}

type HttpService struct {
	Config *config.BucketsConfig
	Db     IDatabase
}

func NewService(Config *config.BucketsConfig, Db IDatabase) *HttpService {
	return &HttpService{
		Config: Config,
		Db:     Db,
	}
}

func (s *HttpService) ChangeDB(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("database file does not exist at path: %s", path)
	}

	if err := s.Db.Close(); err != nil {
		return err
	}

	newEnv := mdbx.New()
	if err := newEnv.Open(path); err != nil {
		return err
	}

	s.Db = newEnv

	return nil
}

func (s *HttpService) ListDataSource() ([]string, error) {
	dataDir := os.Getenv("DATA_DIR")

	var files []string

	err := filepath.Walk(dataDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".dat" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (s *HttpService) ListBuckets() ([]string, error) {
	buckets, err := s.Db.ListBuckets()
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

func (s *HttpService) KeysCount(name string) (uint64, error) {
	count, err := s.Db.CountKeys(name)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *HttpService) GetPage(name string, num int, pageLen int) ([]types.KeyValuePairString, error) {
	foundData, err := s.Db.Read(name, uint64(pageLen), uint64(num))
	if err != nil {
		return nil, err
	}

	data := make([]types.KeyValuePairString, 0)

	for _, kv := range foundData {
		data = append(data, kv.HexKeyHexValue())
	}

	return data, nil
}

func (s *HttpService) KeysCountLength(name string, length uint64) (uint64, []string, error) {
	count, keys, err := s.Db.CountKeysOfLength(name, length)
	if err != nil {
		return 0, nil, err
	}
	return count, keys, nil
}

func (s *HttpService) LookupByKey(bucketName string, searchKey string) ([]byte, error) {
	var foundValue []byte
	if strings.HasPrefix(searchKey, "0x") {
		searchKey = searchKey[2:]
		bytes, err := hex.DecodeString(searchKey)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		foundValue, _ = s.Db.FindByKey(bucketName, bytes)
	} else {
		num := new(big.Int)
		num, ok := num.SetString(searchKey, 16)
		if !ok {
			return nil, fmt.Errorf("error parsing the number")
		}
		//num, err := strconv.ParseUint(searchKey, 10, 64)
		//if err != nil {
		//	fmt.Println("Error:", err)
		//	return nil, err
		//}
		//skuint := utils.Uint64ToBytes(num)
		bytes := utils.BigIntToBytes(num)
		foundValue, _ = s.Db.FindByKey(bucketName, bytes)
	}
	return foundValue, nil
}

func (s *HttpService) SearchByValue(bucketName string, num uint64) ([]string, error) {
	foundKeys, _ := s.Db.FindByValue(bucketName, utils.Uint64ToBytes(num))
	hexKeys := make([]string, 0)

	for _, key := range foundKeys {
		hexKeys = append(hexKeys, hex.EncodeToString(key))
	}
	return hexKeys, nil
}
