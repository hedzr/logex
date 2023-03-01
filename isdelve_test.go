// Copyright © 2020 Hedzr Yeh.

package logex

import (
	"testing"

	"github.com/hedzr/log"
	"github.com/hedzr/log/dir"
	"github.com/hedzr/log/states"
)

func TestEnable(t *testing.T) {
	defer CaptureLog(t).Release()

	// Enable()
	if GetLevel() != log.InfoLevel {
		t.Fatal("wrong level")
	}

	EnableWith(log.DebugLevel)
	if GetLevel() != log.DebugLevel {
		t.Fatal("wrong level")
	}

	t.Logf("cwd: %v", dir.GetCurrentDir())

	// SetupLoggingFormat("json", 1, true, "")
	// SetupLoggingFormat("text", 1, false, "")
}

func TestSetupLoggingFormat(t *testing.T) {
	c := CaptureLog(t)
	defer c.Release()

	EnableWith(log.OffLevel, func() {
		_, _ = c.(interface {
			Write(p []byte) (n int, err error)
		}).Write([]byte("hello"))
	})
	if GetLevel() != log.OffLevel {
		t.Fatal("wrong level")
	}

	log.SetLevel(log.OffLevel)
	// SetupLoggingFormat("any", 1, false, "")

	t.Logf("%v, %v", states.Env().GetDebugMode(), states.Env().GetTraceMode())
	states.Env().SetDebugMode(true)
	states.Env().SetTraceMode(true)
	t.Logf("%v, %v, %v", states.Env().GetDebugMode(), states.Env().GetTraceMode(), states.Env().InDebugging())
}
