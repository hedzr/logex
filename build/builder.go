// Copyright © 2020 Hedzr Yeh.

package build

import (
	"github.com/hedzr/logex"
	"github.com/hedzr/logex/logx/logrus"
	"github.com/hedzr/logex/logx/zap"
	"github.com/hedzr/logex/logx/zap/sugar"
)

func New(config *logex.LoggerConfig) logex.Logger {
	if logex.GetLevel() == logex.OffLevel {
		return &dummyLogger{}
	}

	var logger logex.Logger
	switch config.Backend {
	case "logrus":
		logger = logrus.NewWithConfig(config)
	case "sugar":
		logger = sugar.NewWithConfig(config)
	default:
		logger = zap.New(config.Level, config.TraceMode, config.DebugMode)
	}
	return logger
}

func NewLoggerConfig() *logex.LoggerConfig {
	return logex.NewLoggerConfig()
}
