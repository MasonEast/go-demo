package htmlRender

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	
}

func Render(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Home",
		})
	})
}