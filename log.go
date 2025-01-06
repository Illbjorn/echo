package echo

import (
	"fmt"
	"os"
	"syscall"
)

func write(v []byte) {
	_, _ = syscall.Write(
		wr,
		v,
	)
}

func log(l Level, m string) {
	if l < level {
		return
	}

	if flags > 0 {
		flags.Write(l)
	}

	write([]byte(m))
	write(newline)

	if l == LevelFatal {
		os.Exit(1)
	}
}

func logf(l Level, m string, vs ...any) {
	if l < level {
		return
	}

	log(l, fmt.Sprintf(m, vs...))
}
