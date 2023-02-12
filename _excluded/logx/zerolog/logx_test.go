package zerolog_test

import (
	"testing"

	"github.com/hedzr/log"
	"github.com/hedzr/logex/_excluded/logx/zerolog"

	log3 "github.com/rs/zerolog/log"
)

func TestBasic(t *testing.T) {

	// zerolog3.TimeFieldFormat = zerolog3.TimeFormatUnix

	log3.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log3.Debug().
		Str("Name", "Tom").
		Send()
}

func TestNormal(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "zerolog", "trace")
	logger := zerolog.NewWithConfig(config)
	logger.Printf("hello")
	logger.Infof("hello info")
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")
}

func TestShortTS(t *testing.T) {
	config := log.NewLoggerConfigWith(true, "zerolog", "trace", log.WithTimestamp(true))
	logger := zerolog.NewWithConfig(config)

	logger.Printf("hello, level = %v", logger.GetLevel())
	logger.Infof("hello info, level = %v", logger.GetLevel())
	logger.Warnf("hello warn")
	logger.Errorf("hello error")
	logger.Debugf("hello debug")
	logger.Tracef("hello trace")

	// With

	logger.With("A", "aaa").Infof("ddd fields")
}
