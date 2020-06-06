package models

import "fmt"

func InsertPocInfo(name,region,desc,harm,filePath string) bool{
	fmt.Println(name,region,desc,harm,filePath)
	result,err := Db.Exec("INSERT INTO pocinfo(pocname,pocclass,poccontent,pocvulnreport,pocfilePath)VALUES (?,?,?,?,?)",name,region,desc,harm,filePath)
	if err != nil{
		fmt.Println("insert pocinfo failed,",err)
		return false
	}
	userId,err:= result.LastInsertId()
	fmt.Println(userId)
	return true
}