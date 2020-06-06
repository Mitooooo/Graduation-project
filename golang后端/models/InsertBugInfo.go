package models

import "fmt"

func InsertBugInfo(ipdomain,urladdress,bugname,bugpoc string) bool{
	fmt.Println("InsertBugInfo:::::::",ipdomain,urladdress,bugname,bugpoc)
	result,err := Db.Exec("INSERT INTO buginfo(ipdomain,urladdress,bugname,bugpoc) VALUES (?,?,?,?)",ipdomain,urladdress,bugname,bugpoc)
	if err != nil{
		fmt.Println("insert buginfo failed,",err)
		return false
	}
	userId,err:= result.LastInsertId()
	fmt.Println("InsertBugId: ",userId)
	return true
}