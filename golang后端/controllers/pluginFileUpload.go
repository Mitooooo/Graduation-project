package controllers

import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Fileupload(c *gin.Context) {
	//得到上传的文件
	file, header, err := c.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename
	region := c.PostForm("region")
	name := c.PostForm("name")
	desc := c.PostForm("describe")
	harm := c.PostForm("harm")
	fmt.Println(region)
	if region == "applicationOfWeb" || region == "host" {
		//verificationScript
		path := "verificationScript/" + region + "/"
		fmt.Println(file, err, filename)
		path = path + filename
		out, err := os.Create(path)
		flag := models.InsertPocInfo(name, region, desc, harm, path)
		if flag {
			fmt.Println("插入成功")
		}
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusCreated, "upload successful")
	} else {
		c.String(500, "有误")
		return
	}
}
