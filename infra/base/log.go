package base

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	//定义日志格式
	formatter := &prefixed.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:04:05.000000"
	formatter.ForceFormatting = true //强制格式化
	//设置颜色
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle: "green",
		WarnLevelStyle: "yellow",
		TimestampStyle: "38",
	})
	log.SetFormatter(formatter)
	//日志级别,配置了环境变量log.debug的话就用Debug级别的日志输出，否则使用默认的info级别的输出
	//level := os.Getenv("log.debug")
	//if level == "true" {
	log.SetLevel(log.DebugLevel)
	//}
	//控制台高亮显示
	formatter.ForceColors = true
	formatter.DisableColors = false
	//log.Info("测试info")
	//log.Debug("测试debug")
	//日志文件和滚动配置 (本身没有，通过第三方实现)
	//github.com/lestrrat/go-file-rotatelogs
	//
}
