package logrus

import (
	"github.com/hedzr/log"
	"github.com/sirupsen/logrus"
)

type dzl struct {
	*logrus.Logger
}

func (s *dzl) Tracef(msg string, args ...interface{}) {
	if log.GetTraceMode() {
		s.Logger.Tracef(msg, args...)
	}
}

func (s *dzl) Debugf(msg string, args ...interface{}) {
	s.Logger.Debugf(msg, args...)
}

func (s *dzl) Infof(msg string, args ...interface{}) {
	s.Logger.Infof(msg, args...)
}

func (s *dzl) Warnf(msg string, args ...interface{}) {
	s.Logger.Warnf(msg, args...)
}

func (s *dzl) Errorf(msg string, args ...interface{}) {
	s.Logger.Errorf(msg, args...)
}

func (s *dzl) Fatalf(msg string, args ...interface{}) {
	s.Logger.Fatalf(msg, args...)
}

func (s *dzl) Printf(msg string, args ...interface{}) {
	s.Logger.Infof(msg, args...)
}

//
//

func (s *dzl) Trace(args ...interface{}) {
	if log.GetTraceMode() {
		//s.Logger.Debugw(msg, fields...)
	}
}

func (s *dzl) Debug(args ...interface{}) {
	//s.Logger.Debugw(msg, fields...)
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

func (s *dzl) Print(args ...interface{}) {
	//s.Logger.Infow(msg, fields...)
}

//
//

func (s *dzl) SetLevel(lvl log.Level) {
	s.Logger.SetLevel(logrus.Level(lvl))
}

func (s *dzl) GetLevel() log.Level {
	return log.Level(s.Logger.Level)
}

func (s *dzl) Setup() {
	// initLogger("", "")
}

//func (s *dzl) AsFieldLogger() logx.FieldLogger {
//	return s
//}
