package main

import (
	"github.com/gin-gonic/gin"

	"gindemo/basicauth"
	"gindemo/cookie"
	"gindemo/htmlRender"
	"gindemo/httppusher"
	"gindemo/jsonp"
	"gindemo/log"
	"gindemo/upload"
	"gindemo/validator"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "服务启动成功"})
	})

	htmlRender.Render(r)
	httppusher.Push(r)
	jsonp.Jsonp(r)
	upload.Upload(r)
	basicauth.BasicAuth(r)
	log.Log(r)
	validator.Validator(r)
	cookie.Cookie(r)

	r.Run(":8081")
}