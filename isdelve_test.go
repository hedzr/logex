// Copyright © 2020 Hedzr Yeh.

package logex

import (
	"github.com/hedzr/log"
	"testing"
)

func TestEnable(t *testing.T) {
	defer CaptureLog(t).Release()

	Enable()
	if GetLevel() != log.InfoLevel {
		t.Fatal("wrong level")
	}

	EnableWith(log.DebugLevel)
	if GetLevel() != log.DebugLevel {
		t.Fatal("wrong level")
	}

	SetupLoggingFormat("json", 1)
	SetupLoggingFormat("text", 1)
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
	SetupLoggingFormat("any", 1)

	t.Logf("%v, %v", GetDebugMode(), GetTraceMode())
	SetDebugMode(true)
	SetTraceMode(true)
	t.Logf("%v, %v, %v", GetDebugMode(), GetTraceMode(), InDebugging())
}
