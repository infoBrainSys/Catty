package config

import (
	"github.com/spf13/viper"
)

func Config() (v *viper.Viper) {
	v = viper.New()
	v.SetConfigFile("config/config.yaml")
	v.SetConfigType("yaml")
	v.ReadInConfig()
	return
}
