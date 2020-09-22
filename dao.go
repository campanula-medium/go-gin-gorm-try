package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type DaoConfig struct {
	Args string
}

func InitDao(config DaoConfig) {
	db, _ = gorm.Open(mysql.Open(config.Args), nil)
}

func getDao() *gorm.DB {
	return db
}
