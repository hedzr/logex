package logrus

import (
	"github.com/hedzr/logex"
	"github.com/hedzr/logex/exec"
	"github.com/hedzr/logex/formatter"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

// New create a sugared logrus logger
//
// level can be: "disable", "panic", "fatal", "error", "warn", "info", "debug", "trace"
//
func New(level string, traceMode, debugMode bool, opts ...Opt) logex.Logger {
	logex.SetTraceMode(traceMode)
	logex.SetDebugMode(debugMode)
	// ll := cmdr.GetStringR("logger.level", "info")
	lvl, _ := logex.ParseLevel(level)
	if logex.GetDebugMode() {
		if lvl < logex.DebugLevel {
			lvl = logex.DebugLevel
			level = "debug"
		}
	}
	if logex.GetTraceMode() {
		if lvl < logex.TraceLevel {
			lvl = logex.TraceLevel
			level = "trace"
		}
	}

	log := initLogger(logex.NewLoggerConfig())

	for _, opt := range opts {
		opt(log)
	}

	logger := &dzl{log}
	logger.Setup()
	return logger
}

func NewWithConfig(config *logex.LoggerConfig, opts ...Opt) logex.Logger {
	logex.SetTraceMode(config.TraceMode)
	logex.SetDebugMode(config.DebugMode)
	// ll := cmdr.GetStringR("logger.level", "info")
	lvl, _ := logex.ParseLevel(config.Level)
	if logex.GetDebugMode() {
		if lvl < logex.DebugLevel {
			lvl = logex.DebugLevel
			config.Level = "debug"
		}
	}
	if logex.GetTraceMode() {
		if lvl < logex.TraceLevel {
			lvl = logex.TraceLevel
			config.Level = "trace"
		}
	}

	log := initLogger(config)

	for _, opt := range opts {
		opt(log)
	}

	logger := &dzl{log}
	logger.Setup()
	return logger
}

type Opt func(logger *logrus.Logger)

func WithLoggingFormat(format string) Opt {
	return func(logger *logrus.Logger) {
		logex.SetupLoggingFormat(format, extraSkip)
	}
}

func initLogger(config *logex.LoggerConfig) *logrus.Logger {
	var ll logex.Level
	ll, _ = logex.ParseLevel(config.Level)
	if ll == logex.OffLevel {
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.SetOutput(ioutil.Discard)
		return logrus.New()
	}

	var err error
	if config.Target == "file" {
		logrus.SetLevel(logrus.Level(ll))

		logrus.SetFormatter(&formatter.TextFormatter{ForceColors: true})
		logrus.SetReportCaller(true)

		var file *os.File
		fPath := path.Join(os.ExpandEnv(config.Directory), "output.log")
		fDir := path.Dir(fPath)
		if err = exec.EnsureDirEnh(fDir); err == nil {
			file, err = os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		}
		if err == nil {
			logrus.SetOutput(file)
			return logrus.StandardLogger()
		} else {
			logrus.Warnf("Failed to log to file %q, using default stderr", fPath)
		}
	}

	// setupLoggingFormat(format, 0)
	logex.EnableWith(ll)

	format := "text" // cmdr.GetStringR("logger.format", "text")
	logex.SetupLoggingFormat(format, extraSkip)

	logger := logrus.StandardLogger()
	//logger.Infof("hello, logLevel = %q", logLevel)
	//logrus.Infof("hello, logLevel = %q", logLevel)
	return logger
}

const extraSkip = 1

//func setupLoggingFormat(format string, logexSkipFrames int) {
//	switch format {
//	case "json":
//		logrus.SetFormatter(&logrus.JSONFormatter{
//			TimestampFormat:  "2006-01-02 15:04:05.000",
//			DisableTimestamp: false,
//			PrettyPrint:      false,
//		})
//	default:
//		e := false
//		if logexSkipFrames > 0 {
//			e = true
//		}
//		logrus.SetFormatter(&formatter.TextFormatter{
//			ForceColors:               true,
//			DisableColors:             false,
//			FullTimestamp:             true,
//			TimestampFormat:           "2006-01-02 15:04:05.000",
//			Skip:                      logexSkipFrames,
//			EnableSkip:                e,
//			EnvironmentOverrideColors: true,
//		})
//	}
//}
