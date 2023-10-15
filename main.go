package main

import (
	"GinTemplate/conf"
	"GinTemplate/routes"
	"GinTemplate/routes/middleware"
	"GinTemplate/service/dbs"
	"flag"

	"log"

	"github.com/gin-gonic/gin"
)

var (
	configFile = flag.String("f", "config.yaml", "The Config File")
)

func main() {
	flag.Parse()

	c, err := conf.Init(*configFile)
	if err != nil {
		log.Panicln(err)
	}

	if !c.Dev {
		middleware.InitLog()
		gin.SetMode(gin.ReleaseMode)
	}

	middleware.InitJWT(c.Auth.AccessSecret, c.Auth.AccessExpire)

	app := gin.Default()

	routes.InitRouter(app)

	//dbs.InitMongo(c.Mongo.DataSource, c.Mongo.Database)
	dbs.InitMySql(c.MySql)

	if err := app.Run(c.Host); err != nil {
		log.Panicln(err)
	}
}
