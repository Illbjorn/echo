package echo

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type logFn = func(v string) (int, error)

func TestFlags(t *testing.T) {
	b := setup(t)

	nowDate := time.Now().Format("01-02-06")
	nowTime := time.Now().Format("03:04:05")
	expectTime := fmt.Sprintf("[%s] \n", nowTime)
	expectFlags(t, b, "", expectTime, Info, WITH_TIME)
	expectDate := fmt.Sprintf("[%s] \n", nowDate)
	expectFlags(t, b, "", expectDate, Info, WITH_DATE)
	expectDateTime := fmt.Sprintf("[%s][%s] \n", nowDate, nowTime)
	expectFlags(t, b, "", expectDateTime, Info, WITH_TIME, WITH_DATE)
}

type testWriter interface {
	io.Writer
	fmt.Stringer
}

func expectFlags(t *testing.T, w testWriter, v, expectv string, fn logFn, fs ...Flags) {
	if b, ok := w.(*bytes.Buffer); ok {
		defer b.Reset()
	}

	expectn := len(expectv)
	SetFlags(fs...)
	n, err := fn(v)
	assert.NoError(t, err)
	assert.Equal(t, expectn, n)
	assert.Equal(t, expectv, w.String())
}
