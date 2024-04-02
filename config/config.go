package config

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigParams struct {
	PORT       int    `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DBHOST"`
	DBName     string `mapstructure:"DBNAME"`
	DBUser     string `mapstructure:"DBUSER"`
	DBPort     string `mapstructure:"DBPORT"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	SECRETKEY  string `mapstructure:"SECRETKEY"`
}

func Configure() *ConfigParams {
	var cfg ConfigParams
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error 1: ", err.Error())
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("error 2: ", err.Error())
	}
	return &cfg
}
