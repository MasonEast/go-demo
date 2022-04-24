package main

import (
	"github.com/gin-gonic/gin"

	"gindemo/binding"
	"gindemo/htmlRender"
	"gindemo/httppusher"
	"gindemo/jsonp"
	"gindemo/upload"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "服务启动成功"})
	})

	htmlRender.Render(r)
	httppusher.Push(r)
	jsonp.Jsonp(r)
	binding.Binding(r)
	upload.Upload(r)

	r.Run(":8081")
}