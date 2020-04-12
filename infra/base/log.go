package base

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	ut "github.com/tietang/go-utils"

	"github.com/rifflock/lfshook"
	"github.com/tietang/props/kvs"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var formatter *prefixed.TextFormatter
var lfh *ut.LineNumLogrusHook

func init() {
	//定义日志格式
	formatter = &prefixed.TextFormatter{}
	//设置高亮显示的色彩样式
	formatter.ForceColors = true
	formatter.DisableColors = false
	formatter.ForceFormatting = true //强制格式化
	//设置颜色
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		WarnLevelStyle:  "yellow",
		ErrorLevelStyle: "red",
		FatalLevelStyle: "41",
		PanicLevelStyle: "41",
		DebugLevelStyle: "blue",
		PrefixStyle:     "cyan",
		TimestampStyle:  "37",
	})
	//开启完整时间戳输出和时间戳格式
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02.15:05:05.000"
	//设置日志formatter
	log.SetFormatter(formatter)
	log.SetOutput(colorable.NewColorableStdout())
	//日志级别,配置了环境变量log.debug的话就用Debug级别的日志输出，否则使用默认的info级别的输出
	level := os.Getenv("log.debug")
	if level == "true" {
		log.SetLevel(log.DebugLevel)
	}

	//开启调用函数、文件、代码行信息的输出
	log.SetReportCaller(true)

	//设置函数、文件、代码行信息的输出的hook
	SetLineNumLogrusHook()
}

func SetLineNumLogrusHook() {
	lfh = ut.NewLineNumLogrusHook()
	lfh.EnableFileNameLog = true
	lfh.EnableFuncNameLog = true
	log.AddHook(lfh)
}

//将滚动日志writer共享给iris glog output
var log_writer io.Writer

//初始化log配置，配置logrus日志文件滚动生成
func InitLog(conf kvs.ConfigSource) {
	//设置日志输出级别
	level, err := log.ParseLevel(conf.GetDefault("log.level", "info"))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	if conf.GetBoolDefault("log.enableLineLog", true) {
		lfh.EnableFuncNameLog = true
		lfh.EnableFileNameLog = true
	} else {
		lfh.EnableFuncNameLog = false
		lfh.EnableFileNameLog = false
	}

	//配置日志输出目录
	logDir := conf.GetDefault("log.dir", "./logs")
	logTestDir, err := conf.Get("log.test.dir")
	if err != nil {
		logDir = logTestDir
	}
	logPath := logDir
	logFilePath, _ := filepath.Abs(logPath)
	log.Infof("log dir:%s", logFilePath)
	logFileName := conf.GetDefault("log.file.name", "red_envelop")
	maxAge := conf.GetDurationDefault("log.max.age", time.Hour*24)
	retationTime := conf.GetDurationDefault("log.retation.time", time.Hour*1)
	os.Mkdir(logPath, os.ModePerm)

	baseLoPath := path.Join(logPath, logFileName)
	//设置滚动日志输出writer
	writer, err := rotatelogs.New(
		strings.TrimSuffix(baseLoPath, ".log")+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(baseLoPath),       //生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             //文件最大保存时间
		rotatelogs.WithRotationTime(retationTime), //日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", err)
	}

	//设置日志文件输出的日志格式
	formatter := &log.TextFormatter{}
	formatter.CallerPrettyfier = func(frame *runtime.Frame) (function string, file string) {
		function = frame.Function
		dir, filename := path.Split(frame.File)
		f := path.Base(dir)
		return function, fmt.Sprintf("%s/%s:%d", f, filename, frame.Line)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, //为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, formatter)
	log.AddHook(lfHook)
	log_writer = writer
}
