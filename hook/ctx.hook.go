/*
 * Copyright © 2019 Hedzr Yeh.
 */

package hook

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// ContextHook ...
type ContextHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

var (
	DefaultContextHook = &ContextHook{"src", 5, []logrus.Level{}}
)

// Levels ...
func (hook *ContextHook) Levels() []logrus.Level {
	if len(hook.levels) == 0 {
		return logrus.AllLevels
	}
	return hook.levels
}

// Fire ...
func (hook *ContextHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = hook.findCaller(entry, hook.Skip)
	// entry.Data["file"] = fmt.Sprintf(" %v", entry.Caller.File)
	return nil
}

func (hook *ContextHook) findCaller(entry *logrus.Entry, skip int) string {
	file := ""
	line := 0
	funcName := ""
	for i := 0; i < 10; i++ {
		file, line, funcName = getCaller(entry, skip+i)
		if !strings.HasPrefix(funcName, "logrus.") {
			break
		}
	}
	return fmt.Sprintf("%s:%d:%s", file, line, funcName)
}

func getCaller(entry *logrus.Entry, skip int) (file string, line int, funcName string) {
	var ok bool
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip); ok {
		funcName = runtime.FuncForPC(pc).Name()
		pos := strings.Index(funcName, "/")
		part1 := funcName[0:pos]
		pos = strings.Index(file, part1)
		file = file[pos:]
		funcName = path.Base(funcName)
		// n := 0
		// for i := len(file) - 1; i > 0; i-- {
		// 	if file[i] == '/' {
		// 		n++
		// 		if n >= 2 {
		// 			file = file[i+1:]
		// 			break
		// 		}
		// 	}
		// }
	}
	return
}
