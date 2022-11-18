package routes

import (
	"GinTemplate/api"
	"GinTemplate/routes/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	//跨域
	app.Use(middleware.Cors())

	app.POST("/api/v1/admin/login", api.Login)

	r := app.Group("/api/v1")

	r.Use(api.Auth)

	r.POST("/admin/logout", api.Logout)

}
