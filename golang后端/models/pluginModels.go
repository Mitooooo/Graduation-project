package models

import (
	"database/sql"
	"fmt"
)

// 查询 插件名称 插件类别
//移除 删除插件
type Plugin struct {
	Pocname     string `json:"pocname" form:"pocname"`
	Pocclass    string `json:"pocclass" form:"pocclass"`
	PocfilePath string `json:"pocfilePath" form:"pocfilePath"`
}
type PluginNameAll struct {
	Pocname     string `json:"pocname" form:"pocname"`
	PocfilePath string `json:"pocfilePath" form:"pocfilePath"`
}

func GetPocInfo() *sql.Rows {
	rows, err := Db.Query("SELECT pocname,pocclass,pocfilePath FROM `pocinfo`")
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
func DelPocInfo(pocname string) bool {
	result, err := Db.Exec("DELETE FROM pocinfo WHERE pocname = (?)", pocname)
	if err != nil {
		fmt.Println("delete pocinfo failed,", err)
		return false
	}
	userId, err := result.LastInsertId()
	fmt.Println(userId)
	return true
}
func GetPocNameAll() *sql.Rows {
	rows, err := Db.Query("SELECT pocname,pocfilePath FROM `pocinfo`")
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
