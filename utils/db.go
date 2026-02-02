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
	dns := "root:@(127.0.0.1:3306)/register?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&u.User{})
}
