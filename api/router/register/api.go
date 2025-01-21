package register

import "github.com/gin-gonic/gin"

func RegisterAPI(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping")
	}
}
