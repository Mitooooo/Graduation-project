package models

import (
	"database/sql"
	"fmt"
)

func Checkup(username,password string) (bool) {
	fmt.Println("models中的usenmae 和password",username,password)
	err := Db.QueryRow("SELECT username,password  FROM user where username = ? and password = ?",
		username,password).Scan()
	if err == sql.ErrNoRows {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return false
	}else {
		return true
	}
}