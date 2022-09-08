package log

import (
	"os"

	"github.com/lilihx/chatRoom/common/config"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logInit()
}

func logInit() *logrus.Logger {
	var logger = logrus.New()
	logger.SetLevel(logrus.Level(config.Config.Log.Level))
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func Info(args ...interface{}) {
	logger.Infoln(args)
}

func Error(args ...interface{}) {
	logger.Errorln(args)
}

func Debug(args ...interface{}) {
	logger.Debugln(args)
}

//Deprecated
func Warning(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}
