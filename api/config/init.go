package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", 80)

	viper.SetDefault("rpc.url", "https://sepolia.drpc.org")
	viper.SetDefault("rpc.url2", "https://rpc2.sepolia.org")
	viper.SetDefault("rpc.pollinginterval", "5s")
}
