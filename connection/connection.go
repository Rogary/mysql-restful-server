package connection

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db == nil {
		InitConnection("root:root@tcp(127.0.0.1:3306)/huginn")
	}
	return db
}

func InitConnection(dataSourceName string){
	fmt.Println(dataSourceName)
	db1, err := sql.Open("mysql", dataSourceName)
	db = db1
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	if err != nil {
		fmt.Print(err.Error())
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
}