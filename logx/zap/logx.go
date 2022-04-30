package zap

import (
	"fmt"
	"github.com/hedzr/log"
	"go.uber.org/zap"
	"io"
)

type dzl struct {
	*zap.Logger
	sugar  *zap.SugaredLogger
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
		s.sugar.With(fcvt(s.fields)...).Debugf(msg, args...)
	}
}

func (s *dzl) Debugf(msg string, args ...interface{}) {
	if log.GetDebugMode() {
		s.sugar.With(fcvt(s.fields)...).Debugf(msg, args...)
	}
}

func (s *dzl) Infof(msg string, args ...interface{}) {
	s.sugar.With(fcvt(s.fields)...).Infof(msg, args...)
}

func (s *dzl) Warnf(msg string, args ...interface{}) {
	s.sugar.With(fcvt(s.fields)...).Warnf(msg, args...)
}

func (s *dzl) Errorf(msg string, args ...interface{}) {
	s.sugar.With(fcvt(s.fields)...).Errorf(msg, args...)
}

func (s *dzl) Fatalf(msg string, args ...interface{}) {
	s.sugar.With(fcvt(s.fields)...).Fatalf(msg, args...)
}

func (s *dzl) Panicf(msg string, args ...interface{}) {
	s.sugar.With(fcvt(s.fields)...).Panicf(msg, args...)
}

func (s *dzl) Printf(msg string, args ...interface{}) {
	s.sugar.With(fcvt(s.fields)...).Infof(msg, args...)
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
		msg := fmt.Sprint(args...)
		// s.Logger.Debugw(msg, fields...)
		s.Logger.Info(msg, s.fields...)
	}
}

func (s *dzl) Debug(args ...interface{}) {
	if log.GetDebugMode() {
		// s.Logger.Debugw(msg, fields...)
		msg := fmt.Sprint(args...)
		s.Logger.Debug(msg, s.fields...)
	}
}

func (s *dzl) Info(args ...interface{}) {
	// s.Logger.Infow(msg, fields...)
	msg := fmt.Sprint(args...)
	s.Logger.Info(msg, s.fields...)
}

func (s *dzl) Warn(args ...interface{}) {
	// s.Logger.Warnw(msg, fields...)
	msg := fmt.Sprint(args...)
	s.Logger.Warn(msg, s.fields...)
}

func (s *dzl) Error(args ...interface{}) {
	// s.Logger.Errorw(msg, fields...)
	msg := fmt.Sprint(args...)
	s.Logger.Error(msg, s.fields...)
}

func (s *dzl) Fatal(args ...interface{}) {
	// s.Logger.Fatalw(msg, fields...)
	msg := fmt.Sprint(args...)
	s.Logger.Fatal(msg, s.fields...)
}

func (s *dzl) Panic(args ...interface{}) {
	// s.Logger.Fatalw(msg, fields...)
	msg := fmt.Sprint(args...)
	s.Logger.Panic(msg, s.fields...)
}

func (s *dzl) Print(args ...interface{}) {
	// s.Logger.Infow(msg, fields...)
	msg := fmt.Sprint(args...)
	s.Logger.Info(msg, s.fields...)
}

func (s *dzl) Println(args ...interface{}) {
	// s.Logger.Infow(msg, fields...)
	msg := fmt.Sprint(args...)
	s.Logger.Info(msg, s.fields...)
}

//
//

func (s *dzl) SetLevel(lvl log.Level) {
	// s.Logger.s
}

func (s *dzl) GetLevel() log.Level {
	// panic("implement me")
	return log.DebugLevel
}

func (s *dzl) SetOutput(out io.Writer) {
	// s.Logger.Out = out
}

func (s *dzl) GetOutput() (out io.Writer) {
	out = s // NOTE zap cannot return a writer
	return
}

func (s *dzl) Write(b []byte) (int, error) {
	s.Logger.Info(string(b), s.fields...)
	return 0, nil
}

func (s *dzl) Setup() {
	// initLogger("", "")
	// s.sugar = s.Logger.Sugar()
}

// func (s *dzl) AsFieldLogger() FieldLogger {
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
	Trace(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Print(msg string, fields ...zap.Field)
}
