package log

import (
	"io"
	"os"

    logrus "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 日志切割
	logFile := &lumberjack.Logger{
		Filename:   "./backend.log",
		MaxSize:    100,  // 日志文件大小，单位是 MB
		MaxBackups: 3,    // 最大过期日志保留个数
		MaxAge:     28,   // 保留过期文件最大时间，单位 天
		Compress:   true, // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetLevel(logrus.InfoLevel)
}

var Logger = logrus.New()