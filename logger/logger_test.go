package logger

import (
	"testing"

	"go-micro.dev/v4/logger"
)

func TestLog(t *testing.T) {
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
