package zap

import (
	"github.com/hedzr/logex"
	"go.uber.org/zap"
)

type dzl struct {
	*zap.Logger
	sugar *zap.SugaredLogger
}

func (s *dzl) Tracef(msg string, args ...interface{}) {
	if logex.GetTraceMode() {
		s.sugar.Debugf(msg, args...)
	}
}

func (s *dzl) Debugf(msg string, args ...interface{}) {
	if logex.GetDebugMode() {
		s.sugar.Debugf(msg, args...)
	}
}

func (s *dzl) Infof(msg string, args ...interface{}) {
	s.sugar.Infof(msg, args...)
}

func (s *dzl) Warnf(msg string, args ...interface{}) {
	s.sugar.Warnf(msg, args...)
}

func (s *dzl) Errorf(msg string, args ...interface{}) {
	s.sugar.Errorf(msg, args...)
}

func (s *dzl) Fatalf(msg string, args ...interface{}) {
	s.sugar.Fatalf(msg, args...)
}

func (s *dzl) Printf(msg string, args ...interface{}) {
	s.sugar.Infof(msg, args...)
}

//
//

func (s *dzl) Trace(args ...interface{}) {
	if logex.GetTraceMode() {
		//s.Logger.Debug("", args...)
	}
}

func (s *dzl) Debug(args ...interface{}) {
	if logex.GetDebugMode() {
		//s.Logger.Debugw(msg, fields...)
	}
}

func (s *dzl) Info(args ...interface{}) {
	//s.Logger.Infow(msg, fields...)
}

func (s *dzl) Warn(args ...interface{}) {
	//s.Logger.Warnw(msg, fields...)
}

func (s *dzl) Error(args ...interface{}) {
	//s.Logger.Errorw(msg, fields...)
}

func (s *dzl) Fatal(args ...interface{}) {
	//s.Logger.Fatalw(msg, fields...)
}

func (s *dzl) Print(args ...zap.Field) {
	//s.Logger.Infow(msg, fields...)
}

//
//

func (s *dzl) SetLevel(lvl logex.Level) {
	// s.Logger.s
}

func (s *dzl) GetLevel() logex.Level {
	// panic("implement me")
	return logex.DebugLevel
}

func (s *dzl) Setup() {
	//initLogger("", "")
	//s.sugar = s.Logger.Sugar()
}

//func (s *dzl) AsFieldLogger() FieldLogger {
//	return s
//}

func AsFieldLogger(s logex.Logger) FieldLogger {
	if l, ok := s.(FieldLogger); ok {
		return l
	}
	return nil
}

type FieldLogger interface {
	logex.Logger
	Trace(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Print(msg string, fields ...zap.Field)
}
