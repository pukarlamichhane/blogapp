package database

import (
	"database/sql"
	"fmt"
)

func Getconnect() *sql.DB {
	db, err := sql.Open("mysql", "root:123456789@tcp(localhost:3306)/mydb")
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	return db

}
