package main

import (
	"github.com/tietang/props/ini"
	_ "resk"
	"resk/infra"
	"resk/infra/base"
)

func main() {
	//获取程序运行文件所在地路径
	//file := kvs.GetCurrentFilePath("config.ini", 1) //这一行会报错，记录了编译时候的绝对路径，具体原因未知，包的原因吧
	file := "./config.ini"
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()

}
