package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.ReadInConfig()
}

func Config() *viper.Viper {
	return viper.GetViper()
}
