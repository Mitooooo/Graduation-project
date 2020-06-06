package models

import (
	"database/sql"
	"fmt"
)

func EnterAssetsWithPrimaryInfoCol(assetsName, assetsAddress, startTime, endTime string) bool {
	result, err := Db.Exec("INSERT INTO primaryinfocol(assetsName,assetsAddress,startTime,endTime) VALUES (?,?,?,?)", assetsName, assetsAddress, startTime, endTime)
	if err != nil {
		fmt.Println("insert failed,", err)
		return false
	}
	userId, err := result.LastInsertId()
	fmt.Println(userId)
	return true
}

func EnterAssetsWithSecondaryInfoCol(assetsName, assetsAddress string) bool {
	result, err := Db.Exec("INSERT INTO  secondaryinfocol(assetsName,assetsAddress) VALUES (?,?)", assetsName, assetsAddress)
	if err != nil {
		fmt.Println("insert failed,", err)
		return false
	}
	userId, err := result.LastInsertId()
	fmt.Println(userId)
	return true
}

func common_func(result sql.Result, err error, update_type string) bool {
	if err != nil {
		fmt.Println("insert|update failed,", update_type, err)
		return false
	}
	userId, err := result.LastInsertId()
	fmt.Println(userId)
	return true
}
func UpdateIpdomainWithPrimaryInfoCol(ipdomain, assetsAddress string) bool {
	result, err := Db.Exec("update primaryinfocol set ipdomain=(?) where assetsAddress=(?) ", ipdomain, assetsAddress)
	//UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
	return common_func(result, err, "UpdateIpdomainWithSecondaryInfoCol")
}

//port
func UpdatePortWithPrimaryInfoCol(port, assetsAddress string) bool {
	result, err := Db.Exec("update primaryinfocol set port=(?) where  assetsAddress=(?) ", port, assetsAddress)
	//UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
	return common_func(result, err, "UpdatePortWithSecondaryInfoCol")
}

//urladdress
func UpdateUrlAddressWithPrimaryInfoCol(urladdress string, assetsAddress string) bool {
	result, err := Db.Exec("update primaryinfocol set urladdress=(?) where assetsAddress=(?) ", urladdress, assetsAddress)
	//UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
	return common_func(result, err, "UpdateUrlAddressWithSecondaryInfoCol")
}

//dir
func UpdateDirWithPrimaryInfoCol(dir string, assetsAddress string) bool {
	result, err := Db.Exec("update primaryinfocol set dir=(?) where assetsAddress=(?) ", dir, assetsAddress)
	//UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
	return common_func(result, err, "UpdateDirWithSecondaryInfoCol")
}
func InsertStartTimeInfoCol(startTime string, assetsAddress string) bool {
	result, err := Db.Exec("update primaryinfocol set startTime=(?) where assetsAddress=(?) ", startTime, assetsAddress)
	return common_func(result, err, "InsertStartTimeInfoCol")
}

//fmt.Sprintf("%v",time.Now().Format("2006-01-02 15:04:05"))
func InsertEndTimeInfoCol(startTime string, assetsAddress string) bool {
	result, err := Db.Exec("update primaryinfocol set endTime=(?) where assetsAddress=(?) ", startTime, assetsAddress)
	return common_func(result, err, "InsertEndTimeInfoCol")
}
