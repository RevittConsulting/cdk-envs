package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	v, err := loadConfig("chains", "yaml")
	if err != nil {
		return nil, err
	}

	err = parseConfig(cfg, v)
	if err != nil {
		return nil, err
	}

	chainCfg := &Chains{}
	err = parseConfig(chainCfg, v)
	if err != nil {
		return nil, err
	}
	cfg.Chains = chainCfg

	return cfg, nil
}

func parseConfig(cfg interface{}, v *viper.Viper) error {
	if err := v.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}
	return nil
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
