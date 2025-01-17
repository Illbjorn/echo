package echo

import (
	"path/filepath"
	"runtime"
	"strconv"
)

var pc = make([]uintptr, 100)

func stackDepth() int {
	var res = runtime.Callers(9, pc)
	pc = pc[0:]
	return res
}

func caller(i int) []byte {
	var (
		ret  []byte
		f, l string
		ln   int
	)

	_, f, ln, _ = runtime.Caller(i)
	f = filepath.Base(f)
	l = strconv.FormatInt(int64(ln), 10)

	ret = append(ret, bBrLeft)
	ret = append(ret, []byte(f)...)
	ret = append(ret, bColon)
	ret = append(ret, []byte(l)...)
	ret = append(ret, bBrRight)

	return ret
}
