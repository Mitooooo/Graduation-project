package models

import (
	"database/sql"
	"fmt"
)

type BugInfo struct {
	Ipdomain string `json:"ipdomain" fomr:"assetsName"`
	Urladdress   string `json:"urladdress" form:"urladdress"`
	Bugname string `json:"bugname" form:"bugname"`
	Bugpoc string `json:"bugpoc" form:"bugpoc"`
}

func GetBugInfo(urladdress string) *sql.Rows {
	rows,err := Db.Query("SELECT ipdomain,urladdress,bugname,bugpoc FROM `buginfo` where ipdomain = (?)",urladdress)
	if err != nil{
		fmt.Println(err)
	}
	return rows
}