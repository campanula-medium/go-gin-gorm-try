package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {

	v := viper.New()
	v.SetConfigFile("config.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	InitDao(DaoConfig{Args: v.Get("db.dsn").(string)})
	//RegService()
	InitController(ControllerConfig{Host: v.Get("server.port").(string)})
}

func RegService() {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.0.167:8500"
	config.Token = "7a282761-0a4a-cff0-f3d0-9a08b631612d"
	reg := consulapi.AgentServiceRegistration{}
	reg.Name = "go-consul-service" //注册service的名字
	reg.Address = "192.168.10.44"  //注册service的ip
	reg.Port = 8888                //注册service的端口
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}          //创建consul的检查器
	check.Interval = "5s"                           //设置consul心跳检查时间间隔
	check.HTTP = "http://192.168.10.44:8888/health" //设置检查使用的url

	reg.Check = &check

	client, err := consulapi.NewClient(config) //创建客户端
	if err != nil {
		log.Fatal(err)
	}
	err = client.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}
