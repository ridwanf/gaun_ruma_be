package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	SIGNING_KEY string
}

func GetConfig(path string) (config Configuration, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
