package controllers

import (
	"database/sql"
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
)

func GetAssetsAddress(c *gin.Context) {
	var rows *sql.Rows
	rows = models.GetAssetsAddress()
	AssetsAddressData := make([]models.AssetsAddress, 0)
	var AssetsAddress models.AssetsAddress

	for rows.Next() {
		rows.Scan(&AssetsAddress.AssetsName, &AssetsAddress.AssetsAddress, &AssetsAddress.StartTime, &AssetsAddress.EndTime)
		AssetsAddressData = append(AssetsAddressData, AssetsAddress)
		fmt.Println(AssetsAddress)
	}
	fmt.Println(AssetsAddressData)
	c.JSON(200, gin.H{
		"status": 200,
		"data":   AssetsAddressData,
	})
}
