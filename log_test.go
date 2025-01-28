package echo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	// SetFlags(FlagWithLevel | FlagWithCaller)
	// SetFlags(FlagWithLevel)
	// SetLevel(LevelDebug)
	assert.NoError(t, Info("Hello, %s!", "World"))
	assert.NoError(t, Debug("Hello, %s!", "World"))
	assert.NoError(t, Warn("Hello, %s!", "World"))
	// SetLevel(LevelWarn)
	assert.NoError(t, Debug("Hello, %s!", "World"))
	// SetLevel(LevelInfo)
	assert.NoError(t, Warn("Hello, %s!", "World"))
}
