// Copyright © 2020 Hedzr Yeh.

package logex

import (
	"github.com/hedzr/logex/isdelve"
	"github.com/hedzr/logex/trace"
)

type CmdrMinimal interface {
	InDebugging() bool
	GetDebugMode() bool
	GetTraceMode() bool
	SetDebugMode(b bool)
	SetTraceMode(b bool)
}

func InDebugging() bool   { return env.InDebugging() }
func GetDebugMode() bool  { return env.GetDebugMode() }
func GetTraceMode() bool  { return env.GetTraceMode() }
func SetDebugMode(b bool) { env.SetDebugMode(b) }
func SetTraceMode(b bool) { env.SetTraceMode(b) }

type Env struct {
	debugMode bool
	traceMode bool
}

func (e *Env) InDebugging() bool   { return isdelve.Enabled }
func (e *Env) GetDebugMode() bool  { return e.debugMode || isdelve.Enabled }
func (e *Env) GetTraceMode() bool  { return e.traceMode || trace.IsEnabled() }
func (e *Env) SetDebugMode(b bool) { e.debugMode = b }
func (e *Env) SetTraceMode(b bool) { e.traceMode = b }

var env = &Env{}
