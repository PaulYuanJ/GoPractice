package useful_demos

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
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

var myTime = map[string]interface{}{
	"timestamp": time.Now().Unix(),
}

func getMyTime() map[string]interface{} {
	return myTime
}

func EP_log() {
	fmt.Println(time.Now().Unix())
}
