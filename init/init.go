package main

import (
	"github.com/acaee/metro/service/storage"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {
	dsn := "root:123456@(192.168.0.110)/metro?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println("连接数据库失败：",db)
	}
	db.AutoMigrate(storage.City{})
	
}

