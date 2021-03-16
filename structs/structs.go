package structs

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}

type Book struct {
	gorm.Model
	Title        string
	Author       string
	Release_Date int
}
