package main

import (
	"github.com/gin-gonic/gin"
)

type ControllerConfig struct {
	Host string
}

func InitController(config Service, consul Consul) {
	engine := gin.Default()
	engine.GET(consul.CheckUrl, func(context *gin.Context) {
		context.Status(200)
	})
	router(engine)
	engine.Run(":" + config.Port)
}

func router(engine *gin.Engine) {
	studentGroup(engine)
}

func studentGroup(engine *gin.Engine) {
	group := engine.Group("/student")
	group.GET("/", GetStudent)
	group.POST("/", AddStudent)
	group.PUT("/:id", UpdateStudent)
	group.DELETE("/:id", DeleteStudent)
}
