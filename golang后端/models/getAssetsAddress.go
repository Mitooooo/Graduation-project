package models

import (
	"database/sql"
	"fmt"
)

type AssetsAddress struct {
	//Id         int    `json:"id" form:"id"`
	AssetsName    string `json:"assetsName" fomr:"assetsName"`
	AssetsAddress string `json:"assetsAddress" form:"assetsAddress"`
	StartTime     string `json:"startTime" form:"startTime"`
	EndTime       string `json:"endTime" form:"endTime"`
}

func GetAssetsAddress() *sql.Rows {
	rows, err := Db.Query("SELECT assetsName,assetsAddress,StartTime,EndTime FROM `primaryinfocol`")
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
