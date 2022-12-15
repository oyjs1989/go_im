package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetHome(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/index")
}
