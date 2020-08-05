package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type DaoConfig struct {
	Args string
}

func InitDao(config DaoConfig) {
	db, _ = gorm.Open("mysql", config.Args)
}

func getDao() *gorm.DB {
	return db
}
