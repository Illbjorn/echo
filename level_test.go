package echo

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestLevel(t *testing.T) {
	b := setup(t)

	SetFlags(0)

	// ////////////////////////////////////////////////////////////////////////////
	// Walk up the Levels

	l := SetLevel(LevelDebug)
	assert.Equal(t, LevelWarn, l)
	expectLog(t, LevelDebug, b, basicStdin, basicStdout, len(basicStdout))

	l = SetLevel(LevelWarn)
	assert.Equal(t, LevelDebug, l)
	expectLog(t, LevelDebug, b, basicStdin, "", 0)
	expectLog(t, LevelWarn, b, basicStdin, basicStdout, len(basicStdout))

	l = SetLevel(LevelInfo)
	assert.Equal(t, LevelWarn, l)
	expectLog(t, LevelDebug, b, basicStdin, "", 0)
	expectLog(t, LevelWarn, b, basicStdin, "", 0)
	expectLog(t, LevelInfo, b, basicStdin, basicStdout, len(basicStdout))

	l = SetLevel(LevelError)
	assert.Equal(t, LevelInfo, l)
	expectLog(t, LevelDebug, b, basicStdin, "", 0)
	expectLog(t, LevelWarn, b, basicStdin, "", 0)
	expectLog(t, LevelInfo, b, basicStdin, "", 0)
	expectLog(t, LevelError, b, basicStdin, basicStdout, len(basicStdout))

	l = SetLevel(LevelFatal)
	assert.Equal(t, LevelError, l)
	l = SetLevel(LevelFatal)       // No LevelFatal!
	assert.Equal(t, LevelError, l) // No LevelFatal!
	expectLog(t, LevelDebug, b, basicStdin, "", 0)
	expectLog(t, LevelWarn, b, basicStdin, "", 0)
	expectLog(t, LevelInfo, b, basicStdin, "", 0)
	expectLog(t, LevelError, b, basicStdin, basicStdout, len(basicStdout))
}
