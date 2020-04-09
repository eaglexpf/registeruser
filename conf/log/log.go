package log

import "github.com/sirupsen/logrus"

var (
	Log log
)

type log struct {
	Logger *logrus.Logger
}

func Trace(args ...interface{}) {
	Log.Logger.Trace(args)
}

func Tracef(format string, args ...interface{}) {
	Log.Logger.Tracef(format, args)
}

func Traceln(args ...interface{}) {
	Log.Logger.Traceln(args)
}

func Debug(args ...interface{}) {
	Log.Logger.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	Log.Logger.Debugf(format, args)
}

func Debugln(args ...interface{}) {
	Log.Logger.Debugln(args)
}

func Info(args ...interface{}) {
	Log.Logger.Info(args)
}

func Infof(format string, args ...interface{}) {
	Log.Logger.Infof(format, args)
}

func Infoln(args ...interface{}) {
	Log.Logger.Infoln(args)
}

func Warn(args ...interface{}) {
	Log.Logger.Warn(args)
}

func Warnf(format string, args ...interface{}) {
	Log.Logger.Warnf(format, args)
}

func Warnln(args ...interface{}) {
	Log.Logger.Warnln(args)
}

func Error(args ...interface{}) {
	Log.Logger.Error(args)
}

func Errorf(format string, args ...interface{}) {
	Log.Logger.Errorf(format, args)
}

func Errorln(args ...interface{}) {
	Log.Logger.Errorln(args)
}

func Fatal(args ...interface{}) {
	Log.Logger.Fatal(args)
}

func Fatalf(format string, args ...interface{}) {
	Log.Logger.Fatalf(format, args)
}

func Fatalln(args ...interface{}) {
	Log.Logger.Fatalln(args)
}

func Panic(args ...interface{}) {
	Log.Logger.Panic(args)
}

func Panicf(format string, args ...interface{}) {
	Log.Logger.Panicf(format, args)
}

func Panicln(args ...interface{}) {
	Log.Logger.Panicln(args)
}
