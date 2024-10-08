package router

import (
	"bdanmu/api/router/middleware"
	"bdanmu/api/router/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() *gin.Engine {
	r = gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ws", ws.RegisterClient)
	r.Run(":8080")
	return r
}
