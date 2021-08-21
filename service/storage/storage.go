package storage

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type City struct {
	Code   int		`gorm:"primary_key;not_null;auto_increment"`
	CnName string	`gorm:"not_null"`
	CeName string
	Cpre   string
}

type Connection struct {
	db *gorm.DB
}

func (conn *Connection) Open() {

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
		fmt.Println("连接数据库失败！",err)
	}
	conn.db = db
}



type Storage interface {
	Update() error
}

type Model interface {

}

func (conn *Connection) Update(city []City) error {
	return conn.db.Create(&city).Error
}

func Update(city []City) {
	var conn Connection
	conn.Open()
	err := conn.Update(city)
	fmt.Println(err)
}