package logrus

import (
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"

	"github.com/hedzr/log"
	"github.com/hedzr/log/dir"
	"github.com/hedzr/log/states"
	"github.com/hedzr/logex/formatter"
)

// New create a sugared logrus logger
//
// level can be: "disable", "panic", "fatal", "error", "warn", "info", "debug", "trace"
func New(level string, traceMode, debugMode bool, opts ...Opt) log.Logger {
	states.Env().SetTraceMode(traceMode)
	states.Env().SetDebugMode(debugMode)
	// // ll := cmdr.GetStringR("logger.level", "info")
	// lvl, _ := log.ParseLevel(level)
	// if log.GetDebugMode() {
	//	if lvl < log.DebugLevel {
	//		lvl = log.DebugLevel
	//		level = "debug"
	//	}
	// }
	// if log.GetTraceMode() {
	//	if lvl < log.TraceLevel {
	//		lvl = log.TraceLevel
	//		level = "trace"
	//	}
	// }

	config := log.NewLoggerConfig()
	logger := &dzl{Config: config}

	logger.initLogger()

	for _, opt := range opts {
		opt(logger.Logger)
	}

	logger.Setup()
	log.SetLogger(logger) // .AddSkip(extraSkip))
	return logger
}

// const extraSkip = 1

func NewWithConfigSimple(config *log.LoggerConfig) log.Logger { return NewWithConfig(config) }

func NewWithConfig(config *log.LoggerConfig, opts ...Opt) log.Logger {
	states.Env().SetTraceMode(config.TraceMode)
	states.Env().SetDebugMode(config.DebugMode)
	// ll := cmdr.GetStringR("logger.level", "info")
	lvl, _ := log.ParseLevel(config.Level)
	if states.Env().GetDebugMode() {
		if lvl < log.DebugLevel {
			lvl = log.DebugLevel
			config.Level = "debug"
		}
	}
	if states.Env().GetTraceMode() {
		if lvl < log.TraceLevel {
			lvl = log.TraceLevel
			config.Level = "trace"
		}
	}

	logger := &dzl{Config: config, split: true}

	logger.initLogger()

	for _, opt := range opts {
		opt(logger.Logger)
	}

	logger.Setup()
	log.SetLogger(logger) // .AddSkip(extraSkip))
	return logger
}

type Opt func(logger *logrus.Logger)

// func WithLoggingFormat(format string) Opt {
//	return func(logger *logrus.Logger) {
//		logex.SetupLoggingFormat(format, extraSkip, false, "")
//	}
// }

func (s *dzl) initLogger() *logrus.Logger {
	var ll log.Level
	ll, _ = log.ParseLevel(s.Config.Level)
	if ll == log.OffLevel {
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.SetOutput(dir.Discard)
		return logrus.New()
	}

	var err error
	if s.Config.Target == "file" {
		logrus.SetLevel(logrus.Level(ll))

		logrus.SetFormatter(&formatter.TextFormatter{ForceColors: true, RelativePath: true})
		logrus.SetReportCaller(true)

		var file *os.File
		fPath := path.Join(os.ExpandEnv(s.Config.Directory), "output.log")
		fDir := path.Dir(fPath)
		err = dir.EnsureDir(fDir)
		if err != nil {
			fmt.Printf(`

You're been prompt with a "sudo" requesting because this folder was been creating but need more privileges:

- %v

We must have created the logging output file in it.

`, fDir)
			err = dir.EnsureDirEnh(fDir)
		}

		if err == nil {
			if file, err = os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660); err == nil {
				logrus.SetOutput(file)
				return logrus.StandardLogger()
			}
		}

		logrus.Warnf("Failed to log to file %q, using default stderr", fPath)
	} else if s.split {
		logrus.SetOutput(dir.Discard) // Send all logs to nowhere by default

		logrus.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
			Writer: os.Stderr,
			LogLevels: []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.WarnLevel,
			},
		})
		logrus.AddHook(&writer.Hook{ // Send info and debug logs to stdout
			Writer: os.Stdout,
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
				logrus.DebugLevel,
				logrus.TraceLevel,
			},
		})
	}

	// // setupLoggingFormat(format, 0)
	// logex.EnableWith(ll)
	//
	// format := "text" // cmdr.GetStringR("logger.format", "text")
	// logex.SetupLoggingFormat(format, extraSkip, config.ShortTimestamp, config.TimestampFormat)

	format := "text" // cmdr.GetStringR("logger.format", "text")
	s.setupLoggingFormat(format, 0, s.Config.ShortTimestamp, s.Config.TimestampFormat)

	// s.working = logger.WithField("SKIP", extraSkip)
	// logger.Infof("hello, logLevel = %q", logLevel)
	// logrus.Infof("hello, logLevel = %q", logLevel)
	return s.Logger
}

const (
	defaultTimestampFormat      = "2006-01-02 15:04:05.000000"
	defaultShortTimestampFormat = "01-02 15:04:05.000000"
	// defaultShortestTimestampFormat = "15:04:05.000"
	extraSkip = 0
)

// const extraSkip = 1

func (s *dzl) setupLoggingFormat(format string, logexSkipFrames int, shortTimestamp bool, tsFormat string) {
	if tsFormat == "" {
		tsFormat = defaultTimestampFormat
		if shortTimestamp {
			tsFormat = defaultShortTimestampFormat
		}
	}

	s.format = format
	s.tsFormat = tsFormat
	s.extraFrames = extraSkip
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:  tsFormat,
			DisableTimestamp: false,
			PrettyPrint:      false,
		})
	default:
		e := false
		if logexSkipFrames >= 0 {
			e = true
		}
		logrus.SetFormatter(&formatter.TextFormatter{
			ForceColors:               true,
			DisableColors:             false,
			FullTimestamp:             true,
			TimestampFormat:           tsFormat,
			Skip:                      logexSkipFrames,
			EnableSkip:                e,
			EnvironmentOverrideColors: true,
			QuoteEmptyFields:          true,
			RelativePath:              true,
		})
	}

	logrus.SetReportCaller(true)

	if log.GetLevel() == log.OffLevel {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	s.Logger = logrus.StandardLogger()
	s.AddSkip(0)
}
