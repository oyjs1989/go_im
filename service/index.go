package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title IvanApi Swagger 标题
// @version 1.0 版本
// @description IvanApi Service 描述
// @BasePath /api/v1  基础路径
// @query.collection.format multi
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// @GetHome
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetHome(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/index")
}
