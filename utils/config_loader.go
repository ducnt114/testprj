package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
