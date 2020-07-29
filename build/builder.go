// Copyright © 2020 Hedzr Yeh.

package build

import (
	"github.com/hedzr/log"
	"github.com/hedzr/logex"
	"github.com/hedzr/logex/logx/logrus"
	"github.com/hedzr/logex/logx/zap"
	"github.com/hedzr/logex/logx/zap/sugar"
)

func New(config *log.LoggerConfig) log.Logger {
	if logex.GetLevel() == log.OffLevel {
		return log.NewDummyLogger()
	}

	var logger log.Logger
	switch config.Backend {
	case "logrus":
		logger = logrus.NewWithConfig(config)
	case "sugar":
		logger = sugar.NewWithConfig(config)
	default:
		logger = zap.NewWithConfig(config)
		//logger = zap.New(config.Level, config.TraceMode, config.DebugMode)
	}
	return logger
}

func NewLoggerConfig() *log.LoggerConfig {
	return log.NewLoggerConfig()
}
