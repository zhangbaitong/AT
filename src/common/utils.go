package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangbaitong/go-uuid/uuid"
)

var Conn *sql.DB

func GetDB() (db *sql.DB) {
	if Conn == nil {
		db, err := sql.Open("mysql", "root:111111@tcp(117.78.19.76:3306)/at_db")
		if err != nil {
			fmt.Println("连接数据库失败")
			fmt.Println(err)
			return nil
		}
		Conn = db
	}
	return Conn
}

func GetUID() string {
	return uuid.NewUUID().String()
}
