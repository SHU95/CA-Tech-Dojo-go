package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbconf := "go_user:password@tcp(mysql:3306)/db"
	db, err := sql.Open("mysql", dbconf)
}
