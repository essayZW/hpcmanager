package logger

import (
	"testing"

	"go-micro.dev/v4/logger"
)

func TestNew(t *testing.T) {
	log, err := New("test")
	if err != nil {
		t.Error(err)
		return
	}
	log.Log(logger.DebugLevel, "test", 123, 789)
	log.Log(logger.InfoLevel, "test", 123, 789)
	log.Log(logger.WarnLevel, "test", 123, 789)
	log.Log(logger.ErrorLevel, "test", 123, 789)
}

func TestLog(t *testing.T) {
	SetName("test")
	Debug("logtest", 1)
	Debugf("debug: %v", 2)
	Info("logtest", 3)
	Infof("info: %v", 4)
	Warn("logtest", 5)
	Warnf("warn: %v", 6)
	Error("logtest", 7)
	Errorf("error: %v", 8)
}
