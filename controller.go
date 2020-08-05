package main

import (
	"github.com/gin-gonic/gin"
)

type ControllerConfig struct {
	Host string
}

func InitController(config ControllerConfig) {
	engine := gin.Default()
	router(engine)
	engine.Run(":" + config.Host)
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
