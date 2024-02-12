package buckets

import (
	"encoding/hex"
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/types"
	"github.com/RevittConsulting/cdk-envs/pkg/utils"
	"strconv"
	"strings"
)

type IDatabase interface {
	Close() error
	ListBuckets() ([]string, error)
	CountKeys(bucketName string) (uint64, error)
	CountKeysOfLength(bucketName string, length uint64) (uint64, []string, error)
	FindByKey(bucketName string, key []byte) ([]byte, error)
	FindByValue(bucketName string, value []byte) ([][]byte, error)
	Read(bucketName string, take, offset uint64) ([]types.KeyValuePair, error)
}

type Service struct {
	Config *config.BucketsConfig
	Db     IDatabase
}

func NewService(Config *config.BucketsConfig, Db IDatabase) *Service {
	return &Service{
		Config: Config,
		Db:     Db,
	}
}

func (s Service) ListBuckets() ([]string, error) {
	buckets, err := s.Db.ListBuckets()
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

func (s Service) GetPage(name string, num int, pageLen int) ([]types.KeyValuePairString, error) {
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

func (s Service) KeysCount(name string) (uint64, error) {
	count, err := s.Db.CountKeys(name)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s Service) KeysCountLength(name string, length uint64) (uint64, []string, error) {
	count, keys, err := s.Db.CountKeysOfLength(name, length)
	if err != nil {
		return 0, nil, err
	}
	return count, keys, nil
}

func (s Service) LookupByKey(bucketName string, searchKey string) ([]byte, error) {
	var foundValue []byte
	if strings.HasPrefix(searchKey, "0x") {
		searchKey = searchKey[2:]
		str, err := hex.DecodeString(searchKey)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		foundValue, _ = s.Db.FindByKey(bucketName, str)
	} else {
		num, err := strconv.ParseUint(searchKey, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}
		skuint := utils.Uint64ToBytes(num)
		foundValue, _ = s.Db.FindByKey(bucketName, skuint)
	}
	return foundValue, nil
}

func (s Service) SearchByValue(bucketName string, num uint64) ([]string, error) {
	foundKeys, _ := s.Db.FindByValue(bucketName, utils.Uint64ToBytes(num))
	hexKeys := make([]string, 0)

	for _, key := range foundKeys {
		hexKeys = append(hexKeys, hex.EncodeToString(key))
	}
	return hexKeys, nil
}
