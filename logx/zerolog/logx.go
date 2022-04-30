package zerolog

import (
	"fmt"
	"github.com/hedzr/log"
	"github.com/hedzr/log/exec"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
)

func newZerologWrapper(config *log.LoggerConfig, opts ...Opt) *dzl {
	s := &dzl{cfg: config, skip: config.ExtraSkip}

	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	lvl, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		log.Errorf("cannot parse %q to zerolog.Level", err)
	}
	// println("level ", lvl, config.Level)
	zerolog.SetGlobalLevel(lvl)
	log.SetLevel(llvl(lvl)) // sync zerolog level to hedzr/log

	var out zerolog.ConsoleWriter = zerolog.NewConsoleWriter()
	out.Out = os.Stderr
	if config.Format == "json" {
		out.NoColor = true
	}
	if config.TimestampFormat != "" {
		out.TimeFormat = config.TimestampFormat
	} else if !config.ShortTimestamp {
		// time.RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		out.TimeFormat = "2006-01-02 15:04:05.000000"
	} else {
		out.TimeFormat = "01-02 15:04:05.000000"
	}
	zerolog.TimeFieldFormat = out.TimeFormat

	if config.Target == "file" {
		out.NoColor = true

		var file *os.File
		fPath := path.Join(os.ExpandEnv(s.cfg.Directory), "output.log")
		fDir := path.Dir(fPath)
		err = exec.EnsureDir(fDir)
		if err != nil {
			fmt.Printf(`

You're been prompt with a "sudo" requesting because this folder was been creating but need more privileges:

- %v

We must have created the logging output file in it.

`, fDir)
			err = exec.EnsureDirEnh(fDir)
		}

		if err == nil {
			if file, err = os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660); err != nil {
				log.Errorf("cannot create logging file %q", fPath)
			} else {
				out.Out = file
			}
		}
	}

	s.lout = zerolog.MultiLevelWriter(out)
	s.with = zerolog.New(s.lout).With()

	// s.Logger = s.with.Logger()
	s.working = s.with
	return s
}

type dzl struct {
	cfg     *log.LoggerConfig
	with    zerolog.Context
	working zerolog.Context
	sl      *zerolog.Logger
	skip    int

	lout zerolog.LevelWriter
}

func (s *dzl) get() *zerolog.Logger {
	// if s.sl == nil {
	slc := s.working.CallerWithSkipFrameCount(extraSkip + s.skip)
	var sl zerolog.Logger
	sl = slc.Timestamp().Logger()
	s.sl = &sl
	// }
	return s.sl
}

func (s *dzl) With(key string, val interface{}) log.Logger {
	s.working = s.working.Interface(key, val)
	return s
}

func (s *dzl) WithFields(fields map[string]interface{}) log.Logger {
	for key, val := range fields {
		s.working = s.working.Interface(key, val)
	}
	return s
}

func (s *dzl) AddSkip(skip int) log.Logger {
	s.skip += skip
	return s
}

func (s *dzl) Tracef(msg string, args ...interface{}) {
	if log.GetTraceMode() {
		s.get().Trace().Msgf(msg, args...)
	}
}

func (s *dzl) Debugf(msg string, args ...interface{}) {
	if log.GetDebugMode() {
		s.get().Debug().Msgf(msg, args...)
	}
}
func (s *dzl) Infof(msg string, args ...interface{})  { s.get().Info().Msgf(msg, args...) }
func (s *dzl) Warnf(msg string, args ...interface{})  { s.get().Warn().Msgf(msg, args...) }
func (s *dzl) Errorf(msg string, args ...interface{}) { s.get().Error().Msgf(msg, args...) }
func (s *dzl) Fatalf(msg string, args ...interface{}) { s.get().Fatal().Msgf(msg, args...) }
func (s *dzl) Panicf(msg string, args ...interface{}) { s.get().Panic().Msgf(msg, args...) }
func (s *dzl) Printf(msg string, args ...interface{}) { s.get().Printf(msg, args...) }

//
//

func (s *dzl) Trace(args ...interface{}) {
	if log.GetTraceMode() {
		s.get().Trace().Msg(fmt.Sprint(args...))
	}
}
func (s *dzl) Debug(args ...interface{}) {
	if log.GetDebugMode() {
		s.get().Debug().Msg(fmt.Sprint(args...))
	}
}
func (s *dzl) Info(args ...interface{})    { s.get().Info().Msg(fmt.Sprint(args...)) }
func (s *dzl) Warn(args ...interface{})    { s.get().Warn().Msg(fmt.Sprint(args...)) }
func (s *dzl) Error(args ...interface{})   { s.get().Error().Msg(fmt.Sprint(args...)) }
func (s *dzl) Fatal(args ...interface{})   { s.get().Fatal().Msg(fmt.Sprint(args...)) }
func (s *dzl) Panic(args ...interface{})   { s.get().Panic().Msg(fmt.Sprint(args...)) }
func (s *dzl) Print(args ...interface{})   { s.get().Print(args...) }
func (s *dzl) Println(args ...interface{}) { s.get().Print(args...) }

//
//

func (s *dzl) SetLevel(lvl log.Level) {
	l := zlvl(lvl)
	// log.Infof("put lvl %v", l)
	zerolog.SetGlobalLevel(l)
}

func zlvl(lvl log.Level) zerolog.Level {
	switch lvl {
	case log.PanicLevel:
		return zerolog.PanicLevel
	case log.FatalLevel:
		return zerolog.FatalLevel
	case log.ErrorLevel:
		return zerolog.ErrorLevel
	case log.WarnLevel:
		return zerolog.WarnLevel
	case log.InfoLevel:
		return zerolog.InfoLevel
	case log.DebugLevel:
		return zerolog.DebugLevel
	case log.TraceLevel:
		return zerolog.TraceLevel
	case log.OffLevel:
		return zerolog.Disabled
	}
	return zerolog.NoLevel
}

func llvl(lvl zerolog.Level) log.Level {
	switch lvl {
	case zerolog.DebugLevel:
		return log.DebugLevel
	case zerolog.InfoLevel:
		return log.InfoLevel
	case zerolog.WarnLevel:
		return log.WarnLevel
	case zerolog.ErrorLevel:
		return log.ErrorLevel
	case zerolog.FatalLevel:
		return log.FatalLevel
	case zerolog.PanicLevel:
		return log.PanicLevel
	case zerolog.NoLevel:
		return log.OffLevel
	case zerolog.Disabled:
		return log.OffLevel
	case zerolog.TraceLevel:
		return log.TraceLevel
	}
	return log.InfoLevel
}

func (s *dzl) GetLevel() log.Level {
	lvl := zerolog.GlobalLevel()
	// log.Infof("got lvl = %v", lvl)
	return llvl(lvl)
}

func (s *dzl) SetOutput(out io.Writer) {
	s.lout = zerolog.MultiLevelWriter(out)
	s.with = zerolog.New(s.lout).With()
	s.working = s.with
}

func (s *dzl) GetOutput() (out io.Writer) {
	return s.lout
}

func (s *dzl) Setup() {
	// initLogger("", "")
}

// func (s *dzl) AsFieldLogger() logx.FieldLogger {
//	return s
// }
