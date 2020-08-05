package main

func main() {

	InitDao(DaoConfig{Args: "root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"})
	InitController(ControllerConfig{Host: "8888"})
}
