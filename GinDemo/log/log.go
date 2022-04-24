package log

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Log (r *gin.Engine){
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}