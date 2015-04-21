package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangbaitong/go-uuid/uuid"
)

//the global var of db connection,it will be create only once.
var Conn *sql.DB

//Get db connection from mysql
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

//get uuid lik a227cedf-e806-11e4-8666-3c075419d855
func GetUID() string {
	return uuid.NewUUID().String()
}
