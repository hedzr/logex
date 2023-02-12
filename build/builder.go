// Copyright © 2020 Hedzr Yeh.

package build

import (
	"github.com/hedzr/log"
	"github.com/hedzr/logex/logx/logrus"
)

func init() {
	savedStdLogger = log.GetLogger()
}

var savedStdLogger log.Logger

// New creates and returns a Logger instance from LoggerConfig
func New(config *log.LoggerConfig) log.Logger {
	l, _ := log.ParseLevel(config.Level)
	if l == log.OffLevel {
		return log.NewDummyLogger()
	}

	if bf, ok := builders[config.Backend]; ok {
		return bf(config)
	}

	var logger log.Logger
	switch config.Backend {
	case "dummy", "none", "off":
		logger = log.NewDummyLogger()
	case "std", "standard":
		logger = savedStdLogger // log.NewStdLogger()
	// case "logrus":
	default:
		logger = logrus.NewWithConfig(config)
		// case "zero","zerolog:
		// 	logger = zerolog.NewWithConfig(config)
		// case "zap":
		// 	logger = zap.NewWithConfig(config)
		// case "sugar":
		// 	logger = sugar.NewWithConfig(config)
		// default:
		// 	logger = zap.NewWithConfig(config)
		// 	// logger = zap.New(config.Level, config.TraceMode, config.DebugMode)
	}
	return logger
}

var builders map[string]log.BuilderFunc

// RegisterBuilder register a builder for your logger.
func RegisterBuilder(backendName string, builderFunc log.BuilderFunc) {
	builders[backendName] = builderFunc
}

func init() {
	builders = make(map[string]log.BuilderFunc)

	builders["dummy"] = log.NewDummyLoggerWithConfig
	builders["none"] = log.NewDummyLoggerWithConfig
	builders["off"] = log.NewDummyLoggerWithConfig

	builders["std"] = NewStdLoggerWithConfig
	builders["standard"] = NewStdLoggerWithConfig
	builders["go"] = NewStdLoggerWithConfig

	builders["logrus"] = logrus.NewWithConfigSimple
	// builders["sugar"] = sugar.NewWithConfigSimple
	// builders["zap"] = zap.NewWithConfigSimple
	// builders["zerolog"] = zerolog.NewWithConfigSimple
	// builders["zero"] = zerolog.NewWithConfigSimple
}

// NewLoggerConfig returns a default LoggerConfig
func NewLoggerConfig() *log.LoggerConfig {
	c := log.NewLoggerConfig()
	// c.DebugMode = log.GetDebugMode()
	// c.TraceMode = log.GetTraceMode()
	return c
}

// NewStdLoggerWithConfig return a stdlib `log` logger
func NewStdLoggerWithConfig(config *log.LoggerConfig) log.Logger {
	l, _ := log.ParseLevel(config.Level)
	// return &stdLogger{Level: l, skip: 1}
	savedStdLogger.SetLevel(l)
	return savedStdLogger
}

// NewLoggerConfigWith returns a default LoggerConfig
func NewLoggerConfigWith(enabled bool, backend, level string, opts ...log.Opt) *log.LoggerConfig {
	c := log.NewLoggerConfigWith(enabled, backend, level, opts...)
	return c
}
