package models

import (
	"database/sql"
	"fmt"
)

type AllAssetsInfo struct {
	//Id         int    `json:"id" form:"id"`
	AssetsName string `json:"assetsName" fomr:"assetsName"`
	AssetsAddress   string `json:"assetsAddress" form:"assetsAddress"`
	ExecStatus string `json:"execStatus" form: "execStatus"`
}

func GetAllAssetsInfo() *sql.Rows {
	rows,err := Db.Query("SELECT assetsName,assetsAddress,execStatus FROM `assets`")
	if err != nil{
		fmt.Println(err)
	}
	return rows
}
