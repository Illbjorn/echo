package echo

import (
	"os"
	"syscall"
)

var wr syscall.Handle = syscall.Handle(os.Stdout.Fd())
