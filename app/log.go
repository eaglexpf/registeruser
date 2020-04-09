package app

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	system_log "log"
	"registeruser/entity/global"
	"registeruser/log"
)

func InitLog() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	system_log.SetOutput(logger.Writer())
	log.Log.Logger = logger
	if global.CONFIG.Log.File.Status {
		writeFile()
	}
}

func writeFile() {
	log.Log.Logger.SetOutput(&lumberjack.Logger{
		Filename:   global.CONFIG.Log.File.Path,
		MaxSize:    global.CONFIG.Log.File.MaxSize, // 文件大小[单位mb]
		MaxBackups: global.CONFIG.Log.File.FileNum, // 保留文件个数
		MaxAge:     global.CONFIG.Log.File.DayNum,  // 保留天数
		Compress:   true,                           // 是否压缩日志
	})
}

//
//func writeKafka()  {
//
//}
//
//func writeES()  {
//
//}
