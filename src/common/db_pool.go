package common

import (
	"fmt"
	"time"
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbPool struct {

    DriverName string
    DataSourceName string
    IsPrimary   bool
    MaxPoolSize int
    Conns       chan *sql.DB
}

//建立数据库连接池，传入参数为：连接池最大连接数，服务名的连接数，是主库连接池还是备库连接池

func CreateDbPool(maxPoolsize int, strDriverName string, strDataSourceName string,isPrimary bool) *DbPool {

    	dbPool := &DbPool{MaxPoolSize: maxPoolsize, DriverName: strDriverName,
	DataSourceName:strDataSourceName,IsPrimary: isPrimary}
    	flag := make(chan bool, dbPool.MaxPoolSize/2)
    	go func() {
        for i := 0; i < dbPool.MaxPoolSize/2; i++ {
            conn, err := sql.Open(strDriverName, strDataSourceName)
            if err != nil {
                fmt.Println(err)
            }
            dbPool.PutConn(conn)
            flag <- true
        }
    	}()

    	for i := 0; i < dbPool.MaxPoolSize/2; i++ {
        <-flag
    	}

    	return dbPool
}

//从连接池中获取连接
func (this *DbPool) GetConn() (*sql.DB, error) {
    if len(this.Conns) == 0 {
        go func() {
            for i := 0; i < this.MaxPoolSize/2; i++ {
                conn, err := sql.Open(this.DriverName, this.DataSourceName)
                if err != nil {
				fmt.Println("连接数据库失败")
				fmt.Println(err)
                }
                this.PutConn(conn)
            }
        }()
    }

    //判断是否能在3秒内获取连接，如果不能就报错
    select {
    //读取通道里的数据库连接，如果读不到就返回报错
    case connChan, ok := <-this.Conns:
        {
            if ok {
                return connChan, nil
            } else {
                return nil, errors.New("数据库连接获取异常，可能已经被关闭！")
            }
        }
    //如果被阻塞三秒仍没有获取到连接，则就返回错误
    case <-time.After(time.Second * 3):
        return nil, errors.New("获取数据库连接超时！")
    }
}

//把连接放入连接池中
func (this *DbPool) PutConn(conn *sql.DB) {
    if this.Conns == nil {
        this.Conns = make(chan *sql.DB, this.MaxPoolSize)
    }
    if len(this.Conns) >= this.MaxPoolSize {
        conn.Close()
        return
    }
    this.Conns <- conn
}
