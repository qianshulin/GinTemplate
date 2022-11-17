package middleware

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func InitLog() {
	g, _ := rotatelogs.New("log/gin_%Y%m%d.log")
	gin.DefaultWriter = io.MultiWriter(g)

	l, _ := rotatelogs.New("log/bin_%Y%m%d.log")
	log.SetPrefix("[BIN] ")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(l)
}
