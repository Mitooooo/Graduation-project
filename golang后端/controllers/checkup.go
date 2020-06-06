package controllers

import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Checkup(c *gin.Context) {
	var loginStatus  string
	username := c.Query("username")
	pwd := c.Query("password")
	if models.Checkup(username,pwd) {
		fmt.Println("用户名密码正确")
		loginStatus = "用户名密码正确"
	}else{
		fmt.Println("用户名密码错误")
		loginStatus = "用户名密码错误"
	}
	c.String(http.StatusOK,loginStatus)
}