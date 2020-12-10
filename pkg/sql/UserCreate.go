package sql

import (
	"go-mission/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateUser(data *model.User) {
	dsn := "root:password@tcp(127.0.0.1:5506)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	user := model.User{Name: data.Name, Token: data.Token}
	db.Create(&user)
}
