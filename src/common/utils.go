package common

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhangbaitong/go-uuid/uuid"
	"log"
	"os"
	"time"
)

//the global var of db connection,logger,redis pool.it will be create only once.
var (
	Conn   *sql.DB
	Logger *log.Logger
	pool   *redis.Pool
	dbpool *DbPool
)

//Get db connection from mysql
func GetDB() (db *sql.DB) {
	if dbpool == nil {
		dbpool = CreateDbPool(20, "mysql", "tomzhao:111111@tcp(127.0.0.1:3306)/at_db",true)
	}

	conn, err := dbpool.GetConn()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return conn	
}

func FreeDB(db *sql.DB) {
	dbpool.PutConn(db)
}

//get uuid lik a227cedf-e806-11e4-8666-3c075419d855
func GetUID() string {
	return uuid.NewUUID().String()
}

//get app logger
func Log() *log.Logger {
	if Logger == nil {
		Logger = log.New(os.Stdout, "AT-Resource : ", log.Ldate|log.Ltime|log.Lshortfile)
		Logger.Print("logger init success ...")
	}
	return Logger
}

//get redis connection pool
func GetRedisPool() *redis.Pool {
	if pool == nil {
		pool = &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", "10.122.75.218:6379")
				if err != nil {
					return nil, err
				}
				// if _, err := c.Do("AUTH", password); err != nil {
				//  c.Close()
				//  return nil, err
				// }
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return pool
}
