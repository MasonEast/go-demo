package httppusher

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func Push(r *gin.Engine) {
	r.SetHTMLTemplate(html)
	r.Static("assets", "./assets")

	r.GET("/push", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("fail to push: %v", err)
			}
		}

		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
}