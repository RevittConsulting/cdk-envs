package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", 8080)
	viper.SetDefault("dbfile", "mdbx.dat")
	viper.SetDefault("shutdowntime", 0)

	viper.SetDefault("rpc.url", "https://sepolia.drpc.org")
	viper.SetDefault("rpc.url2", "https://rpc2.sepolia.org")
	viper.SetDefault("rpc.cardonaurl", "https://rpc.cardona.zkevm-rpc.com/")
	viper.SetDefault("rpc.pollinginterval", "5s")
}
