package logrus

import (
	"github.com/hedzr/log"
	"github.com/hedzr/log/dir"
	"github.com/hedzr/logex/formatter"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

type dzl struct {
	*logrus.Logger

	rw          sync.RWMutex
	working     *entry // *logrus.Entry
	extraFrames int    // tell txtfmtr how many extra-frames should be ignored
	skip        int    // and further skips
	format      string
	tsFormat    string
	split       bool

	Config *log.LoggerConfig
}

func (s *dzl) w() (w *entry) {
	s.rw.Lock()
	defer s.rw.Unlock()
	w = s.working
	return w
}

func (s *dzl) r() (w *entry) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	w = s.working
	return w
}

func (s *dzl) With(key string, val interface{}) log.Logger {
	l := s.r().With(key, val)
	return l
}

func (s *dzl) WithFields(fields map[string]interface{}) log.Logger {
	l := s.r().WithFields(fields)
	return l
}

func (s *dzl) AddSkip(increments int) log.Logger {
	s.rw.Lock()
	defer s.rw.Unlock()
	skip := s.skip + increments
	s.working = &entry{s.Logger.WithField(formatter.SKIP, skip), s}
	return s.working
}

func (s *dzl) Tracef(msg string, args ...interface{}) {
	if log.GetTraceMode() {
		s.r().Tracef(msg, args...)
	}
}

func (s *dzl) Debugf(msg string, args ...interface{}) {
	s.r().Debugf(msg, args...)
}

func (s *dzl) Infof(msg string, args ...interface{}) {
	s.r().Infof(msg, args...)
}

func (s *dzl) Warnf(msg string, args ...interface{}) {
	// sav := s.Logger.Out
	// s.Logger.Out = os.Stderr
	s.r().Warnf(msg, args...)
	// s.Logger.Out = sav
}

func (s *dzl) Errorf(msg string, args ...interface{}) {
	// sav := s.Logger.Out
	// s.Logger.Out = os.Stderr
	s.r().Errorf(msg, args...)
	// s.Logger.Out = sav
}

func (s *dzl) Fatalf(msg string, args ...interface{}) {
	// sav := s.Logger.Out
	// s.Logger.Out = os.Stderr
	s.r().Fatalf(msg, args...)
	// s.Logger.Out = sav
}

func (s *dzl) Panicf(msg string, args ...interface{}) {
	// sav := s.Logger.Out
	// s.Logger.Out = os.Stderr
	s.r().Panicf(msg, args...)
	// s.Logger.Out = sav
}

func (s *dzl) Printf(msg string, args ...interface{}) {
	s.r().Infof(msg, args...)
}

//
//

func (s *dzl) Trace(args ...interface{}) {
	if log.GetTraceMode() {
		s.r().Trace(args...)
	}
}

func (s *dzl) Debug(args ...interface{}) {
	s.r().Debug(args...)
}

func (s *dzl) Info(args ...interface{}) {
	s.r().Info(args...)
}

func (s *dzl) Warn(args ...interface{}) {
	// sav := s.Logger.Out
	// s.Logger.Out = os.Stderr
	s.r().Warn(args...)
	// s.Logger.Out = sav
}

func (s *dzl) Error(args ...interface{}) {
	// sav := s.Logger.Out
	// s.Logger.Out = os.Stderr
	s.r().Error(args...)
	// s.Logger.Out = sav
}

func (s *dzl) Fatal(args ...interface{}) {
	// sav := s.Logger.Out
	// s.Logger.Out = os.Stderr
	s.r().Fatal(args...)
	// s.Logger.Out = sav
}

func (s *dzl) Print(args ...interface{}) {
	s.r().Print(args...)
}

//
//

func (s *dzl) SetLevel(lvl log.Level)     { s.Logger.SetLevel(logrus.Level(lvl)) }
func (s *dzl) GetLevel() log.Level        { return log.Level(s.Logger.Level) }
func (s *dzl) SetOutput(out io.Writer)    { s.Logger.Out = out }
func (s *dzl) GetOutput() (out io.Writer) { return s }

func (s *dzl) Write(p []byte) (n int, err error) {
	s.r().AddSkip(0).Infof("%s", string(p))
	n = len(string(p))
	return
}

func (s *dzl) Setup() {
	// logrus.SetFormatter(&formatter.TextFormatter{ForceColors: true})
	// logrus.SetReportCaller(true)

	lvl := s.GetLevel()
	if lvl == log.OffLevel {
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.SetOutput(dir.Discard)
	} else {
		logrus.SetLevel(logrus.Level(lvl))
		// logrus.SetOutput(os.Stdout)
	}
	log.SetLevel(lvl)

	// logrus.SetFormatter(&formatter.TextFormatter{ForceColors: true})
	// logrus.SetReportCaller(true)
}

// func (s *dzl) AsFieldLogger() logx.FieldLogger {
//	return s
// }
