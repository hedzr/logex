package sugar_test

import (
	"testing"

	"github.com/hedzr/log"
	"github.com/hedzr/logex/_excluded/logx/zap/sugar"
)

func TestNormal(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "sugar", "trace")
	logger := sugar.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
}

func TestShortTS(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "sugar", "trace", log.WithTimestamp(true))
	logger := sugar.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
	// With

	logger.With("A", "aaa").Infof("hello info")
}
