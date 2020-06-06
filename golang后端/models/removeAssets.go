package models

import "fmt"

func RemoveAssets(assetsName string) bool{
	result,err := Db.Exec("DELETE FROM assets WHERE assetsName = (?)",assetsName)
	if err != nil{
		fmt.Println("remove assets ERROR!! err: ,",err)
		return false
	}
	userId,err:= result.LastInsertId()
	fmt.Println(userId)
	return true
}