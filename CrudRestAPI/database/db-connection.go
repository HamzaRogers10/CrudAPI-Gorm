package database

import (
	"CrudRestAPI/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var urlDSN = "root:Hamza@10@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
var err error

func Migration() {
	DB, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	User := models.User{}
	err := DB.AutoMigrate(&User)
	if err != nil {
		return
	}
}
