package models

import (
	"database/sql"
	"fmt"
)

//对 pocinfo 中 pocfilepath 的获取
// 主要用做controllers 中的execTask 做任务的执行


type PocInfo struct {
	PocName string `json:"pocname" form:"pocname"`
	PocFilePath string `json:"pocfilePath" form:"pocfilePath"`
}

func GetPocFilePath() *sql.Rows {
	rows,err := Db.Query("SELECT pocname,pocfilePath FROM `pocinfo`")
	if err != nil{
		fmt.Println(err)
	}
	return rows
}
