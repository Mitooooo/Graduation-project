package controllers

import (
	"database/sql"
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
)

func GetAllAssetsInfo(c *gin.Context){
	var rows *sql.Rows
	rows = models.GetAllAssetsInfo()
	AllAssetsInfoData := make([]models.AllAssetsInfo,0)
	var Assets  models.AllAssetsInfo

	for rows.Next(){
		rows.Scan(&Assets.AssetsName,&Assets.AssetsAddress,&Assets.ExecStatus)
		AllAssetsInfoData = append(AllAssetsInfoData,Assets)
		//fmt.Println(AssetsAddress)
	}
	fmt.Println(AllAssetsInfoData)
	c.JSON(200,gin.H{
		"status": 200,
		"data": AllAssetsInfoData,
	})
}