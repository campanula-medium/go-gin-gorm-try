package main

import (
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {

	InitDao(DaoConfig{Args: "root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"})
	RegService()
	InitController(ControllerConfig{Host: "8888"})
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
