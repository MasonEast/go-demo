package main

import (
	"github.com/gin-gonic/gin"

	"gindemo/htmlRender"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "服务启动成功"})
	})

	htmlRender.Render(r)

	r.Run(":8081")
}