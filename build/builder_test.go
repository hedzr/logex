package build

import (
	"github.com/hedzr/log"
	"testing"
)

func TestNew(t *testing.T) {
	for k := range builders {
		test(t, k)
	}
}

func test(t *testing.T, backend string) {
	t.Logf("For backend %q", backend)
	config := log.NewLoggerConfigWith(true, backend, "trace")
	logger := New(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")

	config = log.NewLoggerConfigWith(true, backend, "trace", log.WithTimestamp(true))
	logger = New(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")

	// With

	logger.With("A", "aaa").Infof("hello info")
}
