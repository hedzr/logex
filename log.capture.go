/*
 * Copyright © 2019 Hedzr Yeh.
 */

package logex

import (
	"io"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/hedzr/log"
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
	testing.TB
	origOut io.Writer
}

func (tl logCapturer) Write(p []byte) (n int, err error) {
	tl.Logf((string)(p))
	return len(p), nil
}

func (tl logCapturer) Release() {
	log.SetOutput(tl.origOut)
}

// CaptureLogOld redirects logrus output to testing.Log
func CaptureLogOld(tb testing.TB) LogCapturer {
	lc := logCapturer{TB: tb, origOut: logrus.StandardLogger().Out}
	if !testing.Verbose() {
		log.SetOutput(lc)
	}
	return &lc
}

// CaptureLog redirects logrus output to testing.Log
func CaptureLog(tb testing.TB) LogCapturer {
	lc := logCapturer{TB: tb, origOut: log.GetOutput()}
	log.SetOutput(lc)
	return &lc
}

// CaptureLog redirects logrus output to testing.Log
func CaptureLogV(tb testing.TB) LogCapturer {
	lc := logCapturer{TB: tb, origOut: log.GetOutput()}
	if !testing.Verbose() {
		log.SetOutput(lc)
	}
	return &lc
}
