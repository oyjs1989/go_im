package router

import (
	"im_go/docs"
	"im_go/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// ...
	r := gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// server
	r.GET("/", service.GetHome)
	r.GET("/index", service.GetIndex)
	r.GET("/user", service.GetUser)

	return r
}
