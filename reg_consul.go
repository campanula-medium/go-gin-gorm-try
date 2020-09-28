package main

import (
	consulapi "github.com/hashicorp/consul/api"
	"strconv"
)

func InitConsul(consul Consul) {
	config := consulapi.DefaultConfig()
	config.Address = consul.Address
	config.Token = consul.Token
	reg := consulapi.AgentServiceRegistration{}
	reg.Name = consul.Name            //注册service的名字
	reg.Address = consul.LocalAddress //注册service的ip
	reg.Port = consul.LocalPort       //注册service的端口
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}                                                                //创建consul的检查器
	check.Interval = "5s"                                                                                 //设置consul心跳检查时间间隔
	check.HTTP = "http://" + consul.LocalAddress + ":" + strconv.Itoa(consul.LocalPort) + consul.CheckUrl //设置检查使用的url

	reg.Check = &check

	client, _ := consulapi.NewClient(config) //创建客户端
	client.Agent().ServiceRegister(&reg)

}
