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

	var out = format([]byte(m), l, flags)
	out = append(out, bNewline)
	write(out)

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
