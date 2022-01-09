package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configure struct {
	Db struct{
		EngineName string
		UserName string
		PassWord string
		Host string
		Port string
		DbName string
		Charset string
	}
}

var Conf = &Configure{}

func init()  {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	viper.Unmarshal(Conf)
}
