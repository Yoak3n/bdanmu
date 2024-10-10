package router

import (
	"bdanmu/api/router/middleware"
	"bdanmu/api/router/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ws", ws.RegisterClient)
	r.LoadHTMLFiles("resource/index.html")

	r.Static("/assets", "./resource/assets")
	r.GET("/", renderIndex)
	r.Run(":8080")
}
func renderIndex(c *gin.Context) {

	c.HTML(200, "index.html", "")
}
