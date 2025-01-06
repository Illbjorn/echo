package echo

import (
	"testing"
)

func TestLog(t *testing.T) {
	// SetFlags(FlagWithLevel | FlagWithCaller)
	// SetFlags(FlagWithLevel)
	// SetLevel(LevelDebug)
	Info("Hello, %s!", "World")
	Debug("Hello, %s!", "World")
	Warn("Hello, %s!", "World")
	// SetLevel(LevelWarn)
	Debug("Hello, %s!", "World")
	// SetLevel(LevelInfo)
	Warn("Hello, %s!", "World")
}
