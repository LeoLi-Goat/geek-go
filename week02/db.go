package week02

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8mb4")
	if err != nil {
		log.Fatal("sql.Open err:", err)
	}
}

func GetDB() *sql.DB {
	return db
}
