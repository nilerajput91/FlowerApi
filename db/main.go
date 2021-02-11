package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Init for run before main func execution
func Init() *sql.DB {
	db, err := sql.Open("mysql", "nilesh:nilesh@123@tcp(127.0.0.1:3306)/flowerdb")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	fmt.Printf("DB Connected Successfully...")
	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
