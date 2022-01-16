// 数据库初始化，建立连接
package model

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
)

type DbWorker struct {
	dsn string
	Db  *sql.DB
}

func NewDb() DbWorker { //"mysql", "root:skrlab7365@tcp(47.104.87.68:3306)/test1")
	//dbw := DbWorker{dsn: "root:123456@tcp(localhost:3306)/media?charset=utf8mb4"}
	dbw := DbWorker{dsn: "root:skrlab7365@tcp(47.104.87.68:3306)/test1?charset=utf8"}
	//支持下面几种DSN写法，具体看mysql服务端配置，常见为第2种
	//user@unix(/path/to/socket)/dbname?charset=utf8
	//user:Password@tcp(localhost:5555)/dbname?charset=utf8
	//user:Password@/dbname
	//user:Password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	dbtemp, err := sql.Open("mysql", dbw.dsn)
	if err != nil {
		fmt.Println("@@@Error when connecting DB:", err)
	}
	dbw.Db = dbtemp
	return dbw
}

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no Password set
		DB:       0,  // use default DB
	})
	return client
}
