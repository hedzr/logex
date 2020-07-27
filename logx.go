// logx provide the standard interface of logging for what any go libraries want strip off the direct dependency to a known logging library.

// Copyright © 2020 Hedzr Yeh.

package logex

type (
	Logger interface {
		Tracef(msg string, args ...interface{})
		Debugf(msg string, args ...interface{})
		Infof(msg string, args ...interface{})
		Warnf(msg string, args ...interface{})
		Errorf(msg string, args ...interface{})
		Fatalf(msg string, args ...interface{})
		Printf(msg string, args ...interface{})

		SetLevel(lvl Level)
		GetLevel() Level

		// Setup will be invoked once an instance created
		Setup()

		// AsFieldLogger() FieldLogger
	}

	LoggerConfig struct {
		Enabled   bool
		Backend   string // zap, sugar, logrus
		Level     string
		Format    string // text, json, ...
		Target    string // console, file, console+file
		Directory string
		DebugMode bool `json:"-" yaml:"-"`
		TraceMode bool `json:"-" yaml:"-"`

		// MaxSize is the maximum size in megabytes of the log file before it gets
		// rotated. It defaults to 100 megabytes.
		MaxSize int `json:"maxsize" yaml:"maxsize"`

		// MaxAge is the maximum number of days to retain old log files based on the
		// timestamp encoded in their filename.  Note that a day is defined as 24
		// hours and may not exactly correspond to calendar days due to daylight
		// savings, leap seconds, etc. The default is not to remove old log files
		// based on age.
		MaxAge int `json:"maxage" yaml:"maxage"`

		// MaxBackups is the maximum number of old log files to retain.  The default
		// is to retain all old log files (though MaxAge may still cause them to get
		// deleted.)
		MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

		// LocalTime determines if the time used for formatting the timestamps in
		// backup files is the computer's local time.  The default is to use UTC
		// time.
		LocalTime bool `json:"localtime" yaml:"localtime"`

		// Compress determines if the rotated log files should be compressed
		// using gzip. The default is not to perform compression.
		Compress bool `json:"compress" yaml:"compress"`
	}
)

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Enabled:   true,
		Backend:   "sugar",
		Level:     "info",
		Format:    "text",
		Target:    "console",
		Directory: "/var/log",
		DebugMode: GetDebugMode(),
		TraceMode: GetTraceMode(),

		MaxSize:    1024, // megabytes
		MaxBackups: 3,    // 3 backups kept at most
		MaxAge:     7,    // 7 days kept at most
		Compress:   true, // disabled by default
	}
}
