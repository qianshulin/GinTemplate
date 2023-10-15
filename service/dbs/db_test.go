package dbs

import (
	"GinTemplate/conf"
	"flag"
	"fmt"
	"log"
	"testing"
)

func init() {
	var configFile = flag.String("f", "D:\\work\\go\\GinTemplate\\config.yaml", "The Config File")
	c, err := conf.Init(*configFile)
	if err != nil {
		log.Panic(err)
	}
	//InitMongo(c.Mongo.DataSource, c.Mongo.Database)
	InitMySql(c.MySql)
}

func TestDemo(t *testing.T) {
	var data *interface{}
	err := Mysqldb.Table("etc_users").Where("id = 1").First(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
