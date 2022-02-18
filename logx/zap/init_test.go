//go:build go1.16
// +build go1.16

package zap_test

import (
	"github.com/hedzr/log"
	"github.com/hedzr/logex/logx/zap"
	"testing"
)

func TestNormal(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "zap", "trace")
	logger := zap.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
}

func TestShortTS(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "zap", "trace", log.WithTimestamp(true))
	logger := zap.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
}
