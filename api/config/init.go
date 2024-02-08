package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", ":8080")
}
