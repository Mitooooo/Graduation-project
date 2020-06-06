package models

import (
	"fmt"
)

func EnterAssets(assetsName,assetsAddress string) (bool) {
	fmt.Println(assetsName,assetsAddress)
	if SelectAssets(assetsName,assetsAddress) == 0 {
		result,err := Db.Exec("INSERT INTO assets(assetsName,assetsAddress)VALUES (?,?)",assetsName,assetsAddress)
		if err != nil{
			fmt.Println("insert failed,",err)
			return false
		}
		userId,err:= result.LastInsertId()
		fmt.Println(userId)
		return true
	}else {
		fmt.Println("该资产链接已存在")
		return false
	}
}
func SelectAssets(assetsName,assetsAddress string) int {
	var count int64
	err := Db.QueryRow("select count(*) from assets where assetsName = (?) or assetsAddress = (?)",assetsName,assetsAddress).Scan(&count)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("count: ",count)
	return int(count)
}