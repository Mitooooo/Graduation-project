package models

import (
	"database/sql"
	"fmt"
)

type UrlAddressInfocol struct {
	Urladdress string `json:"urladdress" form:"urladdress"`
}

func GetInfocolUrladdress(target string) *sql.Rows{
	rows,err := Db.Query("SELECT `urladdress` FROM infocol WHERE ipdomain = (?)",target)
	if err != nil {
		fmt.Println("查询urladdress数据查询错误",err)
	}
	return rows
}