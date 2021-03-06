package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect")
	}

	db.AutoMigrate(structs.Person{})
	db.AutoMigrate(structs.Book{})
	return db
}
