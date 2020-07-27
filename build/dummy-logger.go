// Copyright © 2020 Hedzr Yeh.

package build

import "github.com/hedzr/logex"

type dummyLogger struct{}

func (d *dummyLogger) Tracef(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Debugf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Infof(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Warnf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Errorf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Fatalf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) Printf(msg string, args ...interface{}) {
	//panic("implement me")
}

func (d *dummyLogger) SetLevel(lvl logex.Level) {
	//panic("implement me")
}

func (d *dummyLogger) GetLevel() logex.Level {
	//panic("implement me")
	return logex.GetLevel()
}

func (d *dummyLogger) Setup() {
	//panic("implement me")
}
