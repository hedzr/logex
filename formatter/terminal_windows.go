//go:build !appengine && !js && windows
// +build !appengine,!js,windows

/*
 * Copyright © 2019 Hedzr Yeh.
 */

package formatter

import (
	"io"
	"os"
	"syscall"

	sequences "github.com/konsorten/go-windows-terminal-sequences"
)

func initTerminal(w io.Writer) {
	switch v := w.(type) {
	case *os.File:
		sequences.EnableVirtualTerminalProcessing(syscall.Handle(v.Fd()), true)
	}
}
