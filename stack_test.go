package echo

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallers(t *testing.T) {
	b := setup(t)

	const thisFunc = "[TestCallers] "
	testCallers(t, b, thisFunc, WITH_CALLER_FUNC)
	const thisFile = "[stack_test.go] "
	testCallers(t, b, thisFile, WITH_CALLER_FILE)
	const thisFuncThisLine = "[stack_test.go=>TestCallers] "
	testCallers(t, b, thisFuncThisLine, WITH_CALLER_FUNC, WITH_CALLER_FILE)
}

func testCallers(t *testing.T, b *bytes.Buffer, expectv string, fs ...Flags) {
	t.Helper()
	defer b.Reset()

	SetFlags(fs...)
	_, err := writeCallers(b, flags)
	assert.NoError(t, err)
	assert.Equal(t, expectv, b.String())
}
