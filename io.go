package echo

import (
	"io"
	"os"
	"unsafe"
)

var writer io.Writer = os.Stderr

func SetWriter(w io.Writer) {
	writer = w
}

// Some frequently used byte slices
//
// Let's just allocate 'em up front
var (
	encloseLeft  = []byte("['")
	encloseRight = []byte("']")
	brackL       = []byte("[")
	brackR       = []byte("]")
	colon        = []byte(":")
	space        = []byte(" ")
	dash         = []byte("-")
	newline      = []byte("\n")
	zero         = []byte("0")
	fatArrow     = []byte("=>")
	_true        = []byte("true")
	_false       = []byte("false")
)

func writeDouble(w io.Writer, v string) (n int, err error) {
	acc := writeAccumulate(&n, &err)

	if len(v) == 1 {
		if acc(w.Write(zero)) {
			return
		}
	}

	if acc(writeString(w, v)) {
		return
	}

	return
}

func writeString(w io.Writer, v string) (int, error) {
	sdata := unsafe.StringData(v)
	return w.Write(unsafe.Slice(sdata, len(v)))
}

func writeAccumulate(n *int, err *error) func(nn int, e error) bool {
	return func(nn int, e error) bool {
		*n += nn
		if e != nil {
			*err = e
			return true
		}
		return false
	}
}
