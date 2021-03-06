package conf

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	system_log "log"
	"registeruser/conf/global"
	"registeruser/conf/log"
)

// 加载log配置
func initLog() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	system_log.SetOutput(logger.Writer())
	log.Log.Logger = logger
	if global.CONFIG.Log.File.Status {
		writeFile()
	}
}

// log写入文件
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
