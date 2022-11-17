package api

import (
	"GinTemplate/models"
	"GinTemplate/routes/middleware"
	"GinTemplate/service"
	"github.com/gin-gonic/gin"
	"log"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		service.Response(c, 401, nil)
		c.Abort()
		return
	}

	verify, err := middleware.VerifyToken(token)
	if err != nil {
		service.Response(c, 10001, nil)
		c.Abort()
		return
	}

	claims := models.Claims{}
	if err := verify.Claims(&claims); err == nil {
		c.Set("gin.jwt.claims", claims)
	}

	c.Next()
}

func Login(c *gin.Context) {
	claims := models.Claims{}
	if err := c.ShouldBindJSON(&claims); err != nil {
		service.Response(c, 10010, nil)
		return
	}

	if !service.FindPasswd(&claims) {
		service.Response(c, 20002, nil)
		return
	}

	if jwtToken, err := middleware.GenerateToken(claims); err == nil {
		service.Response(c, 200, jwtToken)
		return
	}

	service.Response(c, 20010, nil)
}

func Logout(c *gin.Context) {
	claims, _ := c.Get("gin.jwt.claims")
	log.Println(c.FullPath(), claims.(models.Claims))
}
