package models

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@onetheroad")

	if err == nil {
		DB = db
	}
}

func UserLogin(username, password string) string {
	row, _ := DB.Query("SELECT nickname FROM usr WHERE username=? AND password=?", username, password)
	var nickName string

	row.Scan(&nickName)
	log.Println(nickName)
	return nickName

}
