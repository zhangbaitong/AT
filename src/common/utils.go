package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangbaitong/go-uuid/uuid"
)

func GetDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:111111@tcp(117.78.19.76:3306)/at_db")
	if err != nil {
		fmt.Println("连接数据库失败")
		fmt.Println(err)
		return nil
	}
	return db
}

func GetUID() string {
	return uuid.NewUUID().String()
}
