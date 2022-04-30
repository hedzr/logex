package zerolog

import (
	"github.com/hedzr/log"
	"github.com/rs/zerolog"
)

// New create a sugared zap logger
//
// level can be: "disable", "panic", "fatal", "error", "warn", "info", "debug", "trace"
//
func New(level string, traceMode, debugMode bool, opts ...Opt) log.Logger {
	log.SetTraceMode(traceMode)
	log.SetDebugMode(debugMode)
	// ll := cmdr.GetStringR("logger.level", "info")
	lvl, _ := log.ParseLevel(level)
	if log.GetDebugMode() {
		if lvl < log.DebugLevel {
			lvl = log.DebugLevel
			level = "debug"
		}
	}
	if log.GetTraceMode() {
		if lvl < log.TraceLevel {
			lvl = log.TraceLevel
			level = "debug"
		}
	}

	logger := newZerologWrapper(log.NewLoggerConfig(), opts...)
	logger.Setup()
	log.SetLogger(logger)
	return logger
}

// NewWithConfigSimple create a sugared zap sugared logger
func NewWithConfigSimple(config *log.LoggerConfig) log.Logger { return NewWithConfig(config) }

// NewWithConfig create a sugared zap sugared logger
//
// level can be: "disable", "panic", "fatal", "error", "warn", "info", "debug", "trace"
//
func NewWithConfig(config *log.LoggerConfig, opts ...Opt) log.Logger {
	log.SetTraceMode(config.TraceMode)
	log.SetDebugMode(config.DebugMode)
	// ll := cmdr.GetStringR("logger.level", "info")
	lvl, _ := log.ParseLevel(config.Level)
	if log.GetDebugMode() {
		if lvl < log.DebugLevel {
			lvl = log.DebugLevel
			config.Level = "debug"
		}
	}
	if log.GetTraceMode() {
		if lvl < log.TraceLevel {
			lvl = log.TraceLevel
			config.Level = "trace" // zap hasn't `trace` level
		}
	}

	logger := newZerologWrapper(config, opts...)
	logger.Setup()
	log.SetLogger(logger)
	return logger
}

type Opt func(logger zerolog.Logger)

//var curDir = dir.GetCurrentDir()

const extraSkip = 5
