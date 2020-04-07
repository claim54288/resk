package base

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisrecover "github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"resk/infra"
	"time"
)

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisServerStarter struct {
	infra.BaseStarter
}

func (i *IrisServerStarter) Init(ctx infra.StarterContext) {
	//创建iris application实例
	irisApplication = initIris()
	//日志组件配置和扩展
	logger2 := irisApplication.Logger()
	logger2.Install(logrus.StandardLogger())
}

func (i *IrisServerStarter) Start(ctx infra.StarterContext) {
	//把路由信息打印到控制台
	routers := Iris().GetRoutes()
	for _, r := range routers {
		logrus.Infof(r.Trace())
	}
	//启动iris
	port := ctx.Props().GetDefault("app.server.port", "18080")
	Iris().Run(iris.Addr(":" + port)) //冒号前面可以加网卡，不加就是监听所有的网卡，一般不加
}

func (i *IrisServerStarter) StartBlocking() bool {
	return true
}

func initIris() *iris.Application {
	//主要中间件的配置,revocer,日志输出中间件的自定义
	app := iris.New()
	app.Use(irisrecover.New())
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string,
			message interface{}, headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s |",
				now.Format("2006-01-02.15:04:05.000000"), latency.String(), status, ip, method, path, headerMessage)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
