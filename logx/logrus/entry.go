package logrus

import (
	"github.com/hedzr/log"
	"github.com/sirupsen/logrus"
	"io"
)

type entry struct {
	*logrus.Entry
	owner *dzl
}

func (s *entry) SetLevel(lvl log.Level)     { s.Logger.SetLevel(logrus.Level(lvl)) }
func (s *entry) GetLevel() log.Level        { return log.Level(s.Logger.Level) }
func (s *entry) SetOutput(out io.Writer)    { s.Logger.Out = out }
func (s *entry) GetOutput() (out io.Writer) { return s.Logger.Out }
func (s *entry) Setup()                     {}

func (s *entry) AddSkip(skip int) log.Logger {
	return s.owner.AddSkip(skip)
}

//
//

func (s *entry) With(key string, val interface{}) log.Logger {
	e := &entry{s.Entry.WithField(key, val), s.owner}
	return e
}

func (s *entry) WithFields(fields map[string]interface{}) log.Logger {
	e := &entry{s.Entry.WithFields(fields), s.owner}
	return e
}

func (s *entry) Tracef(msg string, args ...interface{}) {
	if log.GetTraceMode() {
		s.Entry.Tracef(msg, args...)
	}
}

func (s *entry) Debugf(msg string, args ...interface{}) {
	s.Entry.Debugf(msg, args...)
}

func (s *entry) Infof(msg string, args ...interface{}) {
	s.Entry.Infof(msg, args...)
}

func (s *entry) Warnf(msg string, args ...interface{}) {
	// e := s.Entry // .WithContext(context.TODO())
	// sav := e.Logger.Out
	// e.Logger.Out = os.Stderr
	// e.Warnf(msg, args...)
	// e.Logger.Out = sav
	s.Entry.Warnf(msg, args...)
}

func (s *entry) Errorf(msg string, args ...interface{}) {
	// e := s.Entry // .WithContext(context.TODO())
	// sav := e.Logger.Out
	// e.Logger.Out = os.Stderr
	// e.Errorf(msg, args...)
	// e.Logger.Out = sav
	s.Entry.Errorf(msg, args...)
}

func (s *entry) Fatalf(msg string, args ...interface{}) {
	// e := s.Entry // .WithContext(context.TODO())
	// sav := e.Logger.Out
	// e.Logger.Out = os.Stderr
	// e.Fatalf(msg, args...)
	// e.Logger.Out = sav
	s.Entry.Fatalf(msg, args...)
}

func (s *entry) Panicf(msg string, args ...interface{}) {
	// e := s.Entry // .WithContext(context.TODO())
	// sav := e.Logger.Out
	// e.Logger.Out = os.Stderr
	// e.Panicf(msg, args...)
	// e.Logger.Out = sav
	s.Entry.Panicf(msg, args...)
}

func (s *entry) Printf(msg string, args ...interface{}) {
	s.Entry.Infof(msg, args...)
}
