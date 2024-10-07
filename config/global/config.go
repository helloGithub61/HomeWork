package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config = viper.New()

func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.WatchConfig()
	Config.AddConfigPath("./")
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file, ", err)
	}
}
