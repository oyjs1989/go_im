package service

import (
	"im_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 摘要 比如获取用户列表
// @Schemes
// @Description 这里写描述 get users
// @Tags ops 标签 比如同类功能使用同一个标签
// @Param Id query int true "Id"     参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} Model.PageResult GetUser 返回结果 200 类型（object就是结构体） 类型 注释
// @Router /user [get]
func GetUser(c *gin.Context) {
	user := models.GetOne()
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
