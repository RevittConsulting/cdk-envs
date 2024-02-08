package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func LoadConfig() (*Config, error) {
	path := getConfigPath(os.Getenv("APP_CONFIG"))
	v, err := loadConfig(path, "yaml")
	if err != nil {
		return nil, err
	}
	cfg, err := parseConfig(v)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return cfg, nil
}

func loadConfig(filename, fileType string) (*viper.Viper, error) {
	viper.SetConfigName(filename)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config file: %w", err)
		}
	}

	return viper.GetViper(), nil
}

func getConfigPath(env string) string {
	if env == "cardona" {
		return "config/hermez-cardona"
	} else {
		return "config/hermez-dev"
	}
}
