package main

func main() {
	config := InitConfig()

	InitDao(config.Db)
	InitConsul(config.Consul)
	InitController(config.Service, config.Consul)
}
