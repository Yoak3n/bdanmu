package router

import (
	"bdanmu/api/router/ws"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	r.GET("/ws", ws.RegisterClient)
	r.Run(":8080")
}
