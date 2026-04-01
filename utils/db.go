package utils

import (
	u "jin_quickly/models"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

// 2. 连接数据库
func InitDB() {
	//dns := "root:root123@tcp(127.0.0.1:3306)/register?charset=utf8&parseTime=True&loc=Local"
	dsn := "app:app123456@tcp(127.0.0.1:3306)/register?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "app:app123456@tcp(127.0.0.1:3306)/register?charset=utf8mb4&parseTime=True&loc=Local&allowPublicKeyRetrieval=true"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&u.User{})
}
