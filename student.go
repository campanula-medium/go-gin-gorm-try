package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type student struct {
	Id   int64  `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
	Age  int    `json:"age" gorm:"age"`
	No   string `json:"no" gorm:"no"`
}

func (s student) TableName() string {
	return "student"
}

func AddStudent(c *gin.Context) {
	var student student
	c.BindJSON(&student)
	getDao().Create(&student)
	c.Status(http.StatusOK)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student student
	c.BindJSON(&student)
	student.Id, _ = strconv.ParseInt(id, 10, 64)
	getDao().Save(&student)
	c.Status(http.StatusOK)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)
	getDao().Where(" id = ? ", idInt).Delete(&student{})
	c.Status(http.StatusOK)
}

func GetStudent(c *gin.Context) {
	students := []student{}
	getDao().Find(&students)
	c.JSON(http.StatusOK, students)
}
