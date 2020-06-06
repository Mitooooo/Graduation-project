package api

import (
	"fmt"
	"gin-vue/pkg/e"
	"gin/models"
	"gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var code int64
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)
	data := make(map[string]interface{})
	isExist := models.Checkup(username, password)
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			c.Request.Header.Add("token", token)
			data["token"] = token
			code = e.SUCCESS
		}
		fmt.Println("登陆成功")

	} else {
		code = e.ERROR_AUTH
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(int(code)),
		"data": data,
	})
}
