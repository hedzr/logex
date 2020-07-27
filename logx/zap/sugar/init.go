package sugar

import (
	"github.com/hedzr/logex"
	"github.com/hedzr/logex/exec"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path"
)

// New create a sugared zap sugared logger
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

// New create a sugared zap sugared logger
//
// level can be: "disable", "panic", "fatal", "error", "warn", "info", "debug", "trace"
//
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

type Opt func(logger *zap.SugaredLogger)

func initLogger(config *logex.LoggerConfig) *zap.SugaredLogger {
	var level zapcore.Level
	_ = level.Set(config.Level)

	if config.Target == "file" {
		var w zapcore.WriteSyncer

		fPath := path.Join(os.ExpandEnv(config.Directory), "output.log")
		fDir := path.Dir(fPath)
		if err := exec.EnsureDirEnh(fDir); err != nil {
			log.Printf("cannot create logging dir %q, error: %v", fDir, err)
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
		return logger.WithOptions(zap.AddCallerSkip(extraSkip)).Sugar()

	} else {
		logCfg := zap.NewDevelopmentConfig()
		logCfg.Level = zap.NewAtomicLevelAt(level)
		logCfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		logCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logCfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		logger, _ := logCfg.Build()
		return logger.WithOptions(zap.AddCallerSkip(extraSkip)).Sugar()
	}

}

const extraSkip = 1
