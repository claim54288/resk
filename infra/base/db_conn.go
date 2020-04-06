package base

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var database sqlbuilder.Database

func UpperDatabase() sqlbuilder.Database {
	if database == nil {
		InitUpperDatabase()
	}

	return database
}

func InitUpperDatabase() {
	//数据库链接配置
	settings := mysql.ConnectionURL{
		User:     "root",
		Password: "rppt",
		Database: "po",
		Host:     "127.0.0.1",
	}
	//打开数据库链接
	db, err := mysql.Open(settings)
	if err != nil {
		panic(err)
	}
	database = db
}
