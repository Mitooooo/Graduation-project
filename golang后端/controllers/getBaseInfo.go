package controllers

import (
	"database/sql"
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetBaseInfo(c *gin.Context) {
	var rows *sql.Rows

	assetsAddress := c.PostForm("address")
	rows = models.GetBaseInfoAll(assetsAddress)
	fmt.Println(rows)
	PrimaryInfoColData := make([]models.PrimaryInfoCol, 0)
	for rows.Next() {
		var infocol models.PrimaryInfoCol
		rows.Scan(&infocol.AssetsName, &infocol.AssetsAddress, &infocol.Ipdomain, &infocol.Port, &infocol.Urladdress, &infocol.Dir)
		PrimaryInfoColData = append(PrimaryInfoColData, infocol)
	}
	for k, _ := range PrimaryInfoColData {
		PrimaryInfoColData[k].Urladdress = ChangDataUrladdress(PrimaryInfoColData[k].Urladdress)
		PrimaryInfoColData[k].Port = ChangeDatePort(PrimaryInfoColData[k].Port)
		PrimaryInfoColData[k].Dir = ChangDataUrladdress(PrimaryInfoColData[k].Dir)
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data":   PrimaryInfoColData,
	})
}

//对数据进行处理，更改成前端容易接纳的形式
func ChangDataUrladdress(data string) string {
	new := strings.ReplaceAll(data, " ", "<br><p>")
	ne := strings.ReplaceAll(new, "[", "")
	n := strings.ReplaceAll(ne, "]", "")
	return n
}
func ChangeDatePort(data string) string {
	ne := strings.ReplaceAll(data, "[", "")
	n := strings.ReplaceAll(strings.ReplaceAll(ne, "]", ""), ",", " ")
	return n
}
