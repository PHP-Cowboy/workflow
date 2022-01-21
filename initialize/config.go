package initialize

import (
	"fmt"
	"workflow/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	var configFileName string
	configFileNamePrefix := "config"
	configFileName = fmt.Sprintf("%s-pro.yaml", configFileNamePrefix)

	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
}
