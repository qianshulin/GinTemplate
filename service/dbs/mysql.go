package dbs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Mysqldb *gorm.DB

func InitMySql(dsn string) {
	var err error
	Mysqldb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Mysql连接错误", err)
	}
}
