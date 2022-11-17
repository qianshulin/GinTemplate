package service

import (
	"GinTemplate/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, &models.Response{
		Code: code,
		Data: data,
		Msg:  models.ErrCode[code],
	})
}
