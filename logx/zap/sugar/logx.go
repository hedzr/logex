package sugar

import (
	"github.com/hedzr/log"
	"go.uber.org/zap"
	"io"
)

type dzl struct {
	Logger *zap.SugaredLogger
	fields []zap.Field
}

func (s *dzl) With(key string, val interface{}) log.Logger {
	s.fields = append(s.fields, zap.Any(key, val))
	return s
}

func (s *dzl) WithFields(fields map[string]interface{}) log.Logger {
	for key, val := range fields {
		s.fields = append(s.fields, zap.Any(key, val))
	}
	return s
}

func (s *dzl) AddSkip(skip int) log.Logger {
	return s
}

func (s *dzl) Tracef(msg string, args ...interface{}) {
	if log.GetTraceMode() {
		s.Logger.With(fcvt(s.fields)...).Debugf(msg, args...)
	}
}

func (s *dzl) Debugf(msg string, args ...interface{}) {
	if log.GetDebugMode() {
		s.Logger.With(fcvt(s.fields)...).Debugf(msg, args...)
	}
}

func (s *dzl) Infof(msg string, args ...interface{}) {
	s.Logger.With(fcvt(s.fields)...).Infof(msg, args...)
}

func (s *dzl) Warnf(msg string, args ...interface{}) {
	s.Logger.With(fcvt(s.fields)...).Warnf(msg, args...)
}

func (s *dzl) Errorf(msg string, args ...interface{}) {
	s.Logger.With(fcvt(s.fields)...).Errorf(msg, args...)
}

func (s *dzl) Fatalf(msg string, args ...interface{}) {
	s.Logger.With(fcvt(s.fields)...).Fatalf(msg, args...)
}

func (s *dzl) Panicf(msg string, args ...interface{}) {
	s.Logger.With(fcvt(s.fields)...).Panicf(msg, args...)
}

func (s *dzl) Printf(msg string, args ...interface{}) {
	s.Logger.With(fcvt(s.fields)...).Infof(msg, args...)
}

func fcvt(fields []zap.Field) (ret []interface{}) {
	for _, i := range fields {
		ret = append(ret, i)
	}
	return
}

//
//

func (s *dzl) Trace(args ...interface{}) {
	if log.GetTraceMode() {
		// s.Logger.Debugw(msg, fields...)
		s.Logger.With(fcvt(s.fields)...).Info(args...)
	}
}

func (s *dzl) Debug(args ...interface{}) {
	if log.GetDebugMode() {
		// s.Logger.Debugw(msg, fields...)
		s.Logger.With(fcvt(s.fields)...).Debug(args...)
	}
}

func (s *dzl) Info(args ...interface{}) {
	// s.Logger.Infow(msg, fields...)
	s.Logger.With(fcvt(s.fields)...).Info(args...)
}

func (s *dzl) Warn(args ...interface{}) {
	// s.Logger.Warnw(msg, fields...)
	s.Logger.With(fcvt(s.fields)...).Warn(args...)
}

func (s *dzl) Error(args ...interface{}) {
	// s.Logger.Errorw(msg, fields...)
	s.Logger.With(fcvt(s.fields)...).Error(args...)
}

func (s *dzl) Fatal(args ...interface{}) {
	// s.Logger.Fatalw(msg, fields...)
	s.Logger.With(fcvt(s.fields)...).Fatal(args...)
}

func (s *dzl) Panic(args ...interface{}) {
	// s.Logger.Fatalw(msg, fields...)
	s.Logger.With(fcvt(s.fields)...).Panic(args...)
}

func (s *dzl) Print(args ...interface{}) {
	// s.Logger.Infow(msg, fields...)
	s.Logger.With(fcvt(s.fields)...).Info(args...)
}

func (s *dzl) Println(args ...interface{}) {
	// s.Logger.Infow(msg, fields...)
	s.Logger.With(fcvt(s.fields)...).Info(args...)
}

//
//

func (s *dzl) SetLevel(lvl log.Level) {
	// panic("implement me")
}

func (s *dzl) GetLevel() log.Level {
	// panic("implement me")
	return log.DebugLevel
}

func (s *dzl) SetOutput(out io.Writer) {
	// s.Logger.Out = out
}

func (s *dzl) GetOutput() (out io.Writer) {
	// out = log.GetOutput() // NOTE zap cannot return a writer
	out = s
	return
}

func (s *dzl) Write(p []byte) (n int, err error) {
	s.Logger.Info(string(p))
	n = len(string(p))
	return
}

func (s *dzl) Setup() {
	// initLogger("", "")
}

// func (s *dzl) AsFieldLogger() logx.FieldLogger {
//	return s
// }

func AsFieldLogger(s log.Logger) FieldLogger {
	if l, ok := s.(FieldLogger); ok {
		return l
	}
	return nil
}

type FieldLogger interface {
	log.Logger
	Trace(args ...zap.Field)
	Debug(args ...zap.Field)
	Info(args ...zap.Field)
	Warn(args ...zap.Field)
	Error(args ...zap.Field)
	Fatal(args ...zap.Field)
	Print(args ...zap.Field)
}
