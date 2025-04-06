package echo

import (
	"fmt"
	"io"
	"strconv"
	"unsafe"
)

type Logger struct {
	w     io.Writer
	pairs [][2][]byte
}

func (l *Logger) With(pairs ...any) *Logger {
	l = l.clone()

	to := len(pairs)
	if to%2 != 0 {
		for i := len(pairs); i > 0; i-- {
			if i%2 == 0 {
				to = i
				break
			}
		}
	}

	for i := 0; i < to; i += 2 {
		k := pairs[i]
		v := pairs[i+1]
		l.pairs = append(l.pairs, [2][]byte{conv(k), conv(v)})
	}

	return l
}

func (l *Logger) clone() *Logger {
	clone := *l
	return &clone
}

func conv(v any) []byte {
	switch v := v.(type) {
	case string:
		return stob(v)
	case fmt.Stringer:
		return stob(v.String())
	case error:
		return stob(v.Error())
	case bool:
		if v {
			return _true
		}
		return _false
	case int:
		str := strconv.Itoa(v)
		b := stob(str)
		return b
	default:
		Fatalf("Found unexpected ['conv'] type ['%T'].", v)
		panic("impossible")
	}
}

func stob(v string) []byte {
	data := unsafe.StringData(v)
	slice := unsafe.Slice(data, len(v))
	return slice
}
