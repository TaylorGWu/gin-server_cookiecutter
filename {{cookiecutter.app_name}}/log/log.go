package log

import (
	"github.com/Sirupsen/logrus"
	"{{cookiecutter.app_name}}/config"
	"os"
)

type logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

var defaultLogger *logrus.Logger

func init() {
	defaultLogger = logrus.New()
	cf := config.Config()
	logConfig := cf.GetStringMap("log")
	file, err := os.Open(logConfig["path"], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 666)
	if err != nil {
		defaultLogger.Out = file
		Info("init logger successfully")
	}
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debug(format, args)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args)
}
