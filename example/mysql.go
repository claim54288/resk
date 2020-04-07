package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	dsName := "root:root@tcp(127.0.0.1:3306)/po?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsName)
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(2)                //设置最大空闲连接数
	db.SetMaxOpenConns(3)                //最大连接数
	db.SetConnMaxLifetime(7 * time.Hour) //最大存活时间，默认是8小时

	defer db.Close()
	//db.Query()
	fmt.Println(db.Ping())

}
