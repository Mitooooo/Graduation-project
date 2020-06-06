package models

import (
	"fmt"
)

func UpdateExecStatus(assetsAddress string,execStatus int) (bool) {
	fmt.Println("UpdateExecStatus ...")
	result,err := Db.Exec("UPDATE assets SET execStatus = (?) WHERE assetsAddress = (?)",execStatus,assetsAddress)
		//UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
	if err != nil{
		fmt.Println("UpdateExecStatus failed,",err)
		return false
	}
	userId,err:= result.LastInsertId()
	fmt.Println(userId)
	return true
}