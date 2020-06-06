package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/info?parseTime=true")

	if err != nil {
		log.Fatal("Error")
	}

	//defer db.Close()
}
