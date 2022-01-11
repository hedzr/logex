package logrus_test

import (
	"github.com/hedzr/log"
	"github.com/hedzr/logex/logx/logrus"
	"testing"
)

func TestNormal(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "logrus", "trace")
	logger := logrus.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
}

func TestShortTS(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "logrus", "trace", log.WithTimestamp(true))
	logger := logrus.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
}
