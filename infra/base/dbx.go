package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"resk/infra"
)

//dbx 数据库实例
var database *dbx.Database

func DbxDataBase() *dbx.Database {
	return database
}

//dbx数据库starter,并且设为全局
type DbxDatabaseStarter struct {
	infra.BaseStarter
}

func (s *DbxDatabaseStarter) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	//数据库配置
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	dbx2, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(dbx2.Ping())
	database = dbx2
}
