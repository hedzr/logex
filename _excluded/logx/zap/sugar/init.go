package sugar

import (
	"fmt"
	log2 "log"
	"os"
	"path"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/hedzr/log"
	"github.com/hedzr/log/dir"
	"github.com/hedzr/log/exec"
)

// New create a sugared zap sugared logger
//
// level can be: "disable", "panic", "fatal", "error", "warn", "info", "debug", "trace"
func New(level string, traceMode, debugMode bool, opts ...Opt) log.Logger {
	log.SetTraceMode(traceMode)
	log.SetDebugMode(debugMode)
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
	//		level = "debug"
	//	}
	// }

	zl := initLogger(log.NewLoggerConfig())

	for _, opt := range opts {
		opt(zl)
	}

	logger := &dzl{zl, nil}
	logger.Setup()
	log.SetLogger(logger)
	return logger
}

// NewWithConfigSimple create a sugared zap sugared logger
func NewWithConfigSimple(config *log.LoggerConfig) log.Logger { return NewWithConfig(config) }

// NewWithConfig create a sugared zap sugared logger
//
// level can be: "disable", "panic", "fatal", "error", "warn", "info", "debug", "trace"
func NewWithConfig(config *log.LoggerConfig, opts ...Opt) log.Logger {
	log.SetTraceMode(config.TraceMode)
	log.SetDebugMode(config.DebugMode)
	// // ll := cmdr.GetStringR("logger.level", "info")
	// lvl, _ := log.ParseLevel(config.Level)
	// if log.GetDebugMode() {
	//	if lvl < log.DebugLevel {
	//		lvl = log.DebugLevel
	//		config.Level = "debug"
	//	}
	// }
	// if log.GetTraceMode() {
	//	if lvl < log.TraceLevel {
	//		lvl = log.TraceLevel
	//		config.Level = "debug" // zap hasn't `trace` level
	//	}
	// }

	zl := initLogger(config)

	for _, opt := range opts {
		opt(zl)
	}

	logger := &dzl{zl, nil}
	logger.Setup()
	log.SetLogger(logger)
	return logger
}

type Opt func(logger *zap.SugaredLogger)

func initLogger(config *log.LoggerConfig) *zap.SugaredLogger {
	var level zapcore.Level
	_ = level.Set(config.Level)

	var skip = extraSkip
	if config.ExtraSkip > 0 {
		skip += config.ExtraSkip
	}

	if config.Target == "file" {
		var w zapcore.WriteSyncer

		fPath := path.Join(os.ExpandEnv(config.Directory), "output.log")
		fDir := path.Dir(fPath)
		err := exec.EnsureDir(fDir)
		if err != nil {
			fmt.Printf(`

You're been prompt with a "sudo" requesting because this folder was been creating but need more privileges:

- %v

We must have created the logging output file in it.

`, fDir)
			err = exec.EnsureDirEnh(fDir)
		}

		if err != nil {
			log2.Printf("cannot create logging dir %q, error: %v", fDir, err)
			return nil
		}

		hook := lumberjack.Logger{
			Filename:   fPath,             // the logging file path
			MaxSize:    config.MaxSize,    // megabytes
			MaxBackups: config.MaxBackups, // 3 backups kept at most
			MaxAge:     config.MaxAge,     // 7 days kept at most
			Compress:   config.Compress,   // disabled by default
		}
		w = zapcore.AddSync(&hook)

		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			w,
			level,
		)
		logger := zap.New(core)
		return logger.WithOptions(zap.AddCallerSkip(skip)).Sugar()

	} else {
		logCfg := zap.NewDevelopmentConfig()
		logCfg.Level = zap.NewAtomicLevelAt(level)
		logCfg.EncoderConfig.EncodeCaller = relCallerEncoder // zapcore.FullCallerEncoder
		logCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		if config.ShortTimestamp {
			// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
			logCfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.999999999")
		}
		logCfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		logger, _ := logCfg.Build()
		return logger.WithOptions(zap.AddCallerSkip(skip)).Sugar()
	}

}

func relCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	c := caller.String()
	if fil, err := filepath.Rel(curDir, c); err == nil {
		enc.AppendString(fil)
	} else {
		enc.AppendString(c)
	}
}

var curDir = dir.GetCurrentDir()

const extraSkip = 2
