package app

import (
	"github.com/sirupsen/logrus"
	system_log "log"
	"registeruser/log"
)

func InitLog() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	//logFile := global.CONFIG.Log.Path
	//var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println("Could Not Open Log File : " + err.Error())
	//}
	//logger.SetOutput(io.MultiWriter(file, os.Stdout))
	system_log.SetOutput(logger.Writer())
	log.Log.Logger = logger
}

//func writeFile()  {
//	log.Log.Logger.SetOutput()
//}
//
//func writeKafka()  {
//
//}
//
//func writeES()  {
//
//}
