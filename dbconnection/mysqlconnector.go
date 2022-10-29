package dbconnection

import (
	"database/sql"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connector() *sql.DB {
	db, err := sql.Open("mysql", "root:mysqltoor@tcp(localhost:3306)/cryptoprice")
	if err != nil {
		f.Println("Error validating sql.Open argument")
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		f.Println("Error verifying connection with db.Ping")
		panic(err.Error())
	}

	f.Println("Successfully connected to Database!")
	return db
}
