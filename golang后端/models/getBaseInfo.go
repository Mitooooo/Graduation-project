package models

import (
	"database/sql"
	"fmt"
)

type PrimaryInfoCol struct {
	AssetsName    string `json:"assetsName" form:"assetsName"`
	AssetsAddress string `json:"assetsAddress" form:"assetsAddress"`
	Ipdomain      string `json:"ipdomain" form:"ipdomain"`
	Port          string `json:"port" form:"port"`
	Urladdress    string `json:"urladdress" form:"urladdress"`
	Dir           string `json:"dir" form:"dir"`
}

func GetBaseInfoAll(assetsAddress string) *sql.Rows {
	rows, err := Db.Query("SELECT assetsName,assetsAddress,ipdomain,port,urladdress,dir FROM `primaryinfocol` where assetsAddress = (?)", assetsAddress)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
