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
	viper.SetDefault("rpc.zkevm", "https://zkevm-rpc.com/")
	viper.SetDefault("rpc.cardonaurl", "https://rpc.cardona.zkevm-rpc.com/")
	viper.SetDefault("rpc.pollinginterval", "5s")

	viper.SetDefault("l1contracts.sequencedbatchtopic", "0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce")
	viper.SetDefault("l1contracts.verificationtopic", "0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe")
	viper.SetDefault("l1contracts.updatel1infotreetopic", "0xda61aa7823fcd807e37b95aabcbe17f03a6f3efd514176444dae191d27fd66b3")
	viper.SetDefault("l1contracts.initialsequencedbatchestopic", "0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f")
}
