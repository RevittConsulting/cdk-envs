package buckets

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/RevittConsulting/chain-dev-utils/internal/buckets/db/mdbx"
	"github.com/RevittConsulting/chain-dev-utils/internal/types"
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
	CountValuesOfLength(bucketName string, length uint64) (uint64, []string, error)
	FindByKey(bucketName string, key []byte) ([][]byte, error)
	FindByValue(bucketName string, value []byte) ([][]byte, error)
	Read(bucketName string, take, offset uint64) ([]types.KeyValuePair, error)
}

type HttpService struct {
	Config *Config
	Db     IDatabase
}

func NewService(Config *Config, Db IDatabase) *HttpService {
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

func (s *HttpService) KeysCount(ctx context.Context) (uint64, error) {
	bucketName := ctx.Value(BucketNameCTX).(string)
	count, err := s.Db.CountKeys(bucketName)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *HttpService) GetPage(ctx context.Context, num int, pageLen int) ([]types.KeyValuePairString, error) {
	bucketName := ctx.Value(BucketNameCTX).(string)
	foundData, err := s.Db.Read(bucketName, uint64(pageLen), uint64(num))
	if err != nil {
		return nil, err
	}

	data := make([]types.KeyValuePairString, 0)

	for _, kv := range foundData {
		data = append(data, kv.HexKeyHexValue())
	}

	return data, nil
}

func (s *HttpService) KeysCountLength(ctx context.Context, length uint64) (uint64, []string, error) {
	bucketName := ctx.Value(BucketNameCTX).(string)
	count, keys, err := s.Db.CountKeysOfLength(bucketName, length)
	if err != nil {
		return 0, nil, err
	}
	return count, keys, nil
}

func (s *HttpService) ValuesCountLength(ctx context.Context, length uint64) (uint64, []string, error) {
	bucketName := ctx.Value(BucketNameCTX).(string)
	count, values, err := s.Db.CountValuesOfLength(bucketName, length)
	if err != nil {
		return 0, nil, err
	}
	return count, values, nil
}

func (s *HttpService) SearchByKey(ctx context.Context, searchKey string) ([]string, error) {
	bucketName := ctx.Value(BucketNameCTX).(string)
	var key []byte
	var err error

	searchKey = strings.TrimPrefix(searchKey, "0x")

	key, err = hex.DecodeString(searchKey)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	foundValues, err := s.Db.FindByKey(bucketName, key)
	if err != nil {
		return nil, err
	}

	hexValues := make([]string, 0)
	for _, value := range foundValues {
		hexValues = append(hexValues, hex.EncodeToString(value))
	}

	return hexValues, nil
}

func (s *HttpService) SearchByValue(ctx context.Context, value string) ([]string, error) {
	bucketName := ctx.Value(BucketNameCTX).(string)
	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(value, 16)
	if !ok {
		return nil, fmt.Errorf("error parsing the number")
	}
	bytes := bigInt.Bytes()

	foundKeys, _ := s.Db.FindByValue(bucketName, bytes)
	hexKeys := make([]string, 0)

	for _, key := range foundKeys {
		hexKeys = append(hexKeys, hex.EncodeToString(key))
	}
	return hexKeys, nil
}
