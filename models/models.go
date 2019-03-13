package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

const (
	USERNAME = "root"
	PASSWORD = "li19980105li"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "ontheroad"
)

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@/%s", USERNAME, PASSWORD, DATABASE)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Open mysql failed,err:%v\n", err)
	}
	return
}
