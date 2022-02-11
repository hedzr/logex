package logrus

import (
	"github.com/hedzr/log"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

type entry struct {
	*logrus.Entry
}

func (s *entry) SetLevel(lvl log.Level)     { s.Logger.SetLevel(logrus.Level(lvl)) }
func (s *entry) GetLevel() log.Level        { return log.Level(s.Logger.Level) }
func (s *entry) SetOutput(out io.Writer)    { s.Logger.Out = out }
func (s *entry) GetOutput() (out io.Writer) { return s.Logger.Out }
func (s *entry) Setup()                     {}

func (s *entry) AddSkip(skip int) log.Logger {
	return &entry{
		s.Entry.WithField("SKIP", skip),
	}
}

//
//

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
	//e := s.Entry // .WithContext(context.TODO())
	//sav := e.Logger.Out
	//e.Logger.Out = os.Stderr
	//e.Warnf(msg, args...)
	//e.Logger.Out = sav
	s.Entry.Warnf(msg, args...)
}

func (s *entry) Errorf(msg string, args ...interface{}) {
	//e := s.Entry // .WithContext(context.TODO())
	//sav := e.Logger.Out
	//e.Logger.Out = os.Stderr
	//e.Errorf(msg, args...)
	//e.Logger.Out = sav
	s.Entry.Errorf(msg, args...)
}

func (s *entry) Fatalf(msg string, args ...interface{}) {
	//e := s.Entry // .WithContext(context.TODO())
	//sav := e.Logger.Out
	//e.Logger.Out = os.Stderr
	//e.Fatalf(msg, args...)
	//e.Logger.Out = sav
	s.Entry.Fatalf(msg, args...)
}

func (s *entry) Panicf(msg string, args ...interface{}) {
	//e := s.Entry // .WithContext(context.TODO())
	//sav := e.Logger.Out
	//e.Logger.Out = os.Stderr
	//e.Panicf(msg, args...)
	//e.Logger.Out = sav
	s.Entry.Panicf(msg, args...)
}

func (s *entry) Printf(msg string, args ...interface{}) {
	s.Entry.Infof(msg, args...)
}

//
//

type dzl struct {
	*logrus.Logger
	Config *log.LoggerConfig
}

func (s *dzl) AddSkip(skip int) log.Logger {
	return &entry{
		s.Logger.WithField("SKIP", skip),
	}
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
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.Logger.Warnf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Errorf(msg string, args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.Logger.Errorf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Fatalf(msg string, args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.Logger.Fatalf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Panicf(msg string, args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.Logger.Panicf(msg, args...)
	//s.Logger.Out = sav
}

func (s *dzl) Printf(msg string, args ...interface{}) {
	s.Logger.Infof(msg, args...)
}

//
//

func (s *dzl) Trace(args ...interface{}) {
	if log.GetTraceMode() {
		s.Logger.Trace(args...)
	}
}

func (s *dzl) Debug(args ...interface{}) {
	s.Logger.Debug(args...)
}

func (s *dzl) Info(args ...interface{}) {
	s.Logger.Info(args...)
}

func (s *dzl) Warn(args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.Logger.Warn(args...)
	//s.Logger.Out = sav
}

func (s *dzl) Error(args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.Logger.Error(args...)
	//s.Logger.Out = sav
}

func (s *dzl) Fatal(args ...interface{}) {
	//sav := s.Logger.Out
	//s.Logger.Out = os.Stderr
	s.Logger.Fatal(args...)
	//s.Logger.Out = sav
}

func (s *dzl) Print(args ...interface{}) {
	s.Logger.Print(args...)
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
