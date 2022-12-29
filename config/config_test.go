package config_test

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
	v.SetConfigType("yaml")
	v.ReadInConfig()

	fmt.Println(v.GetString("DatabaseCfg.DriverName"))
}
