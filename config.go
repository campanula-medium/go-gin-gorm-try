package main

import (
	"fmt"
	"github.com/spf13/viper"
)

var GlobalConfig Config

func InitConfig() Config {
	v := viper.New()
	v.SetConfigFile("config.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = v.Unmarshal(&GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("config convert error: %s \n", err))
	}

	return GlobalConfig
}

type Config struct {
	Service Service
	Db      Db
	Consul  Consul
}

type Service struct {
	Port string
}

type Db struct {
	Dsn string
}

type Consul struct {
	Address      string
	Token        string
	Name         string
	LocalAddress string
	LocalPort    int
	CheckUrl     string
}
