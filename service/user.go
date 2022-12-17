package service

import (
	"im_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := models.GetOne()
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
