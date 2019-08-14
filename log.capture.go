/*
 * Copyright © 2019 Hedzr Yeh.
 */

package logex

import (
	"github.com/sirupsen/logrus"
	"io"
	"testing"
)

// see also: https://github.com/sirupsen/logrus/issues/834
//
// Usage:
//
//   func TestFoo(t *testing.T) {
//     defer logex.CaptureLog(t).Release()
//     …
//   }
//


// LogCapturer reroutes testing.T log output
type LogCapturer interface {
	Release()
}

type logCapturer struct {
	*testing.T
	origOut io.Writer
}

func (tl logCapturer) Write(p []byte) (n int, err error) {
	tl.Logf((string)(p))
	return len(p), nil
}

func (tl logCapturer) Release() {
	logrus.SetOutput(tl.origOut)
}

// CaptureLog redirects logrus output to testing.Log
func CaptureLog(t *testing.T) LogCapturer {
	lc := logCapturer{T: t, origOut: logrus.StandardLogger().Out}
	if !testing.Verbose() {
		logrus.SetOutput(lc)
	}
	return &lc
}
