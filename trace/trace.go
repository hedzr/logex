/*
 * Copyright © 2020 Hedzr Yeh.
 */

package trace

import (
	"sync"
	"sync/atomic"
)

var tracing struct {
	sync.Mutex
	enabled int32
}

func Start() (err error) {
	if atomic.CompareAndSwapInt32(&tracing.enabled, 0, 1) {
		tracing.Lock()
		defer tracing.Unlock()

		// trace.Start()
		// tracing.enabled

		for _, fn := range handlers {
			fn(false, true)
		}
	}

	return
}

func Stop() {
	if atomic.CompareAndSwapInt32(&tracing.enabled, 1, 0) {
		tracing.Lock()
		defer tracing.Unlock()

		// ...

		for _, fn := range handlers {
			fn(true, false)
		}
	}
}

func IsEnabled() bool {
	enabled := atomic.LoadInt32(&tracing.enabled)
	return enabled == 1
}

func RegisterOnTraceModeChanges(onTraceModeChanged Handler) {
	handlers = append(handlers, onTraceModeChanged)
}

type Handler func(old, new bool)

var handlers []Handler
