package infoCollectionModel

import (
	"fmt"
	"gin/models"
)

func InsertAllInfoModel(ip,assetsAddress,port,urladdress,dir string) (bool) {
	result,err := models.Db.Exec("INSERT INTO infocol(ipdomain,assetsAddress,port,urladdress,dir) VALUES (?,?,?,?,?)",ip,assetsAddress,port,urladdress,dir)
	if err != nil{
		fmt.Println("insert infocol All failed,",err)
		return false
	}
	userId,err:= result.LastInsertId()
	fmt.Println(userId)
	return true
}