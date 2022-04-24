package upload

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload (r *gin.Engine) {
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		dst := "./" + file.Filename

		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploadde!", file.Filename))
	})
} 