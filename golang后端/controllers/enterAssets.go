package controllers

import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EnterAssets(c *gin.Context) {
	assetsName := c.PostForm("assetsName")
	assetsAddress := c.PostForm("assetsAddress")

	if assetsName == "" || assetsAddress == "" {
		message := "不允许提交空串"
		c.JSON(401, gin.H{
			"insertStatus": "fail",
			"message":      message,
		})
		fmt.Println(message)
	} else {
		if status(assetsAddress) {
			var flag bool
			flag = models.EnterAssets(assetsName, assetsAddress)
			fmt.Println(flag)
			if flag {
				fmt.Println("资产表插入成功")
				message := "插入成功"
				if models.EnterAssetsWithPrimaryInfoCol(assetsName, assetsAddress, "2020-00-00 00:00:00", "2020-00-00 00:00:00") {
					fmt.Println("PrimaryInfoCol insert success")
					if models.EnterAssetsWithSecondaryInfoCol(assetsName, assetsAddress) {
						fmt.Println("SecondaryInfoCol insert success")
					}
				}
				c.JSON(200, gin.H{
					"insertStatus": "success",
					"message":      message,
				})
			} else {
				//c.String(401,"请勿插入重复数据或非法数据")
				message := "请勿插入重复数据或非法数据"
				c.JSON(201, gin.H{
					"insertStatus": "fail",
					"message":      message,
				})
			}
		} else {
			message := "请勿插入重复数据或非法数据"
			c.JSON(201, gin.H{
				"insertStatus": "fail",
				"message":      message,
			})
		}
		//c.String(200,"您提交的数据为：%v,%v",assetsName,assetsAddress)

	}
}

//测试访问可不可达
func status(address string) bool {
	fmt.Println(address)
	resp, err := http.Get(address)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if address[0:4] == "http" {
			fmt.Println("检查通过")
			return true

		} else {
			fmt.Println("非http起始位")
			return false
		}
	} else {
		fmt.Println("状态码不正确")
		return false
	}
}
