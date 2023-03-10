package Service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

const dsn = "root:123@tcp(127.0.0.1:6603)/Shedule?charset=utf8mb4&parseTime=True"

func InitDB() (err error) {
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	Db.SetMaxOpenConns(50)
	Db.SetMaxIdleConns(50)
	return
}
