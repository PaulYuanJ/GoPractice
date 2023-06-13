package useful_demos

import (
	"github.com/sirupsen/logrus"
	"strings"
)

type LineNumberHook struct{}

func (h *LineNumberHook) Levels() []logrus.Level {
	// Choose specific log level to fire.
	return logrus.AllLevels
}

func (h *LineNumberHook) Fire(entry *logrus.Entry) error {
	//
	strList := strings.Split(entry.Caller.File, "/")

	fileName := strList[len(strList)-1]
	entry.Caller.File = fileName
	return nil
}

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// Add the LineNumberHook to the logger, Can add multi-hooks.
	log.AddHook(&LineNumberHook{})
	return log
}

func EP_log() {
	logger := NewLogger()
	logger.Errorln("Hello")
}