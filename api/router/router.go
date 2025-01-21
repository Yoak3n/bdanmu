package router

import (
	"bdanmu/api/router/middleware"
	"bdanmu/api/router/register"
	"bdanmu/api/router/ws"
	"bdanmu/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ws", ws.RegisterClient)
	api := r.Group("/api")
	register.RegisterAPI(api)
	r.LoadHTMLFiles("resource/index.html")

	r.Static("/assets", "./resource/assets")
	r.GET("/", renderIndex)
	addr := fmt.Sprintf("0.0.0.0:%d", config.Conf.Port)
	r.Run(addr)
}
func renderIndex(c *gin.Context) {

	c.HTML(200, "index.html", "")
}
