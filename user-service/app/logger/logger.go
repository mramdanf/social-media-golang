package logger

import "github.com/jfeng45/glogger"

var Log glogger.Logger

func SetLogger(newLogger glogger.Logger) {
	Log = newLogger
}