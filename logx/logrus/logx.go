package logrus

import (
	"github.com/hedzr/log"
	"github.com/hedzr/logex/formatter"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

type dzl struct {
	*logrus.Logger

	working     *entry // *logrus.Entry
	extraFrames int    // tell txtfmtr how many extra-frames should be ignored
	skip        int    // and further skips
	format      string
	tsFormat    string

	Config *log.LoggerConfig
}

func (s *dzl) AddSkip(increments int) log.Logger {
	s.skip += increments
	s.working = &entry{s.Logger.WithField(formatter.SKIP, s.skip), s}
	return s.working
}

func (s *dzl) Tracef(msg string, args ...interface{}) {
	if log.GetTraceMode() {
		s.working.Tracef(msg, args...)
	}
}

func (s *dzl) Debugf(msg string, args ...interface{}) {
	s.working.Debugf(msg, args...)
}

func (s *dzl) Infof(msg string, args ...interface{}) {
	s.working.Infof(msg, args...)
}

func (s *dzl) Warnf(msg string, args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.working.Warnf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Errorf(msg string, args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.working.Errorf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Fatalf(msg string, args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.working.Fatalf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Panicf(msg string, args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.working.Panicf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Printf(msg string, args ...interface{}) {
	s.working.Infof(msg, args...)
}

//
//

func (s *dzl) Trace(args ...interface{}) {
	if log.GetTraceMode() {
		s.working.Trace(args...)
	}
}

func (s *dzl) Debug(args ...interface{}) {
	s.working.Debug(args...)
}

func (s *dzl) Info(args ...interface{}) {
	s.working.Info(args...)
}

func (s *dzl) Warn(args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.working.Warn(args...)
	//s.Logger.Out = sav
}

func (s *dzl) Error(args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.working.Error(args...)
	//s.Logger.Out = sav
}

func (s *dzl) Fatal(args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.working.Fatal(args...)
	//s.Logger.Out = sav
}

func (s *dzl) Print(args ...interface{}) {
	s.working.Print(args...)
}

//
//

func (s *dzl) SetLevel(lvl log.Level) { s.Logger.SetLevel(logrus.Level(lvl)) }
func (s *dzl) GetLevel() log.Level    { return log.Level(s.Logger.Level) }
func (s *dzl) SetOutput(out io.Writer) {
	s.Logger.Out = out
}
func (s *dzl) GetOutput() (out io.Writer) { return s.Logger.Out }

func (s *dzl) Setup() {
	//logrus.SetFormatter(&formatter.TextFormatter{ForceColors: true})
	//logrus.SetReportCaller(true)

	lvl := s.GetLevel()
	if lvl == log.OffLevel {
		logrus.SetLevel(logrus.ErrorLevel)
		logrus.SetOutput(ioutil.Discard)
	} else {
		logrus.SetLevel(logrus.Level(lvl))
		//logrus.SetOutput(os.Stdout)
	}
	log.SetLevel(lvl)

	//logrus.SetFormatter(&formatter.TextFormatter{ForceColors: true})
	//logrus.SetReportCaller(true)
}

// func (s *dzl) AsFieldLogger() logx.FieldLogger {
//	return s
// }
