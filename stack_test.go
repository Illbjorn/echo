package echo

import (
	"bytes"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCallers(t *testing.T) {
	b := setup(t)

	const thisFunc = "[TestCallers] "
	testCallers(t, b, thisFunc, WithCallerFunc)
	const thisFile = "[stack_test.go] "
	testCallers(t, b, thisFile, WithCallerFile)
	const thisFuncThisLine = "[stack_test.go=>TestCallers] "
	testCallers(t, b, thisFuncThisLine, WithCallerFunc, WithCallerFile)
}

func testCallers(t *testing.T, b *bytes.Buffer, expectv string, fs ...Flags) {
	t.Helper()
	defer b.Reset()

	SetFlags(fs...)
	_, err := writeCallers(b, flags)
	assert.NilError(t, err)
	assert.Equal(t, expectv, b.String())
}
