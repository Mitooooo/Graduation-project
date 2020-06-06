package controllers

import (
	"database/sql"
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
)

func GetBugInfo(c *gin.Context){
	address := c.PostForm("address")
	fmt.Println("获取的address ",address)
	var rows *sql.Rows
	rows = models.GetBugInfo(address)
	BugInfoData := make([]models.BugInfo,0)
	fmt.Println(BugInfoData)
	for rows.Next(){
		var BugInfo  models.BugInfo
		rows.Scan(&BugInfo.Ipdomain, &BugInfo.Urladdress, &BugInfo.Bugname, &BugInfo.Bugpoc)
		BugInfoData = append(BugInfoData,BugInfo)
	}
	c.JSON(200,gin.H{
		"status": 200,
		"data": BugInfoData,
	})
}