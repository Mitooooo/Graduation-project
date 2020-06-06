package models

import (
	"database/sql"
	"fmt"
)

type SelectSingkeAll struct {
	Pluginname string `json:"pluginname" fomr:"pluginname"`
	Pluginpath string `json:"pluginpath" form:"pluginpath"`
	Urladdress string `json:"urladdress" form: "urladdress"`
	Returninfo string `json:"returninfo" form: "returninfo"`
}

func InsertSingleBugInfo(pluginname, pluginpath, urladdress, returninfo string) bool {
	result, err := Db.Exec("INSERT INTO singleonbuginfo(pluginname, pluginpath, urladdress,returninfo) VALUES (?,?,?,?)", pluginname, pluginpath, urladdress, returninfo)
	if err != nil {
		fmt.Println("InsertSingleBugInfo failed,", err)
		return false
	}
	userId, err := result.LastInsertId()
	fmt.Println("InsertSingleBugInfo: ", userId)
	return true
}
func SelectSingkeAllInfo() *sql.Rows {
	rows, err := Db.Query("SELECT pluginname, pluginpath, urladdress, returninfo FROM `singleonbuginfo`")
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func SelectSingkeInfo(pluginname, pluginpath, urladdress string) *sql.Rows {
	rows, err := Db.Query("SELECT pluginname, pluginpath, urladdress, returninfo FROM `singleonbuginfo` where pluginname=(?) and pluginpath=(?) and urladdress=(?)", pluginname, pluginpath, urladdress)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}
