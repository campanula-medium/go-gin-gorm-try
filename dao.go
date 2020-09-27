package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDao(config Db) {
	db, _ = gorm.Open(mysql.Open(config.Dsn), nil)
}

func Add(student *student) error {
	return db.Create(student).Error
}

func Update(student *student) error {
	return db.Save(student).Error
}

func DelById(id int64) error {
	return db.Where(" id = ? ", id).Delete(&student{}).Error
}

func List(students *[]student) error {
	return db.Find(students).Error
}
