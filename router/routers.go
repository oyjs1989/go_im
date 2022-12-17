package router

import (
	"im_go/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// ...
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user", service.GetUser)
	return r
}
