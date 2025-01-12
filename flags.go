package echo

import (
	"path/filepath"
	"runtime"
	"strconv"
)

var flags Flag

func SetFlags(f Flag) {
	flags = f
}

type Flag uint8

const (
	FlagWithLevel Flag = 1 << iota
	FlagWithCaller
)

func (f Flag) Write(l Level) {
	if flags&FlagWithLevel == FlagWithLevel {
		// Produce level
		write(l.Bytes())
		write(space)
	}

	if flags&FlagWithCaller == FlagWithCaller {
		// Produce caller
		caller(3)
		write(space)
	}
}

func caller(i int) {
	var (
		f, l string
		ln   int
	)

	_, f, ln, _ = runtime.Caller(i)
	f = filepath.Base(f)
	l = strconv.FormatInt(int64(ln), 10)

	write(brLeft)
	write([]byte(f))
	write(colon)
	write([]byte(l))
	write(brRight)
}
