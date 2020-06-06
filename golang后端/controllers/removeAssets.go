package controllers

import (
	"gin/models"
	"github.com/gin-gonic/gin"
)
func RemoveAssets( c *gin.Context){
	assetsName := c.PostForm("assetsName")
	flag := models.RemoveAssets(assetsName)
	if flag{
		c.JSON(200, gin.H{
			"removeStatus":  "success",
			"message": "移除成功",
		})
	}else {
		c.JSON(201,gin.H{
			"removeStatus": "fail",
			"message" : "移除失败",
		})
	}
}