package echo

import (
	"io"
	"runtime"
	"strconv"
	"strings"
)

// For some reason the call stack skip count required is different between
// actual code and unit tests - this var exists to swap out the value in tests
var callersSkip = 4

func writeCallers(w io.Writer, f Flags) (n int, err error) {
	withFunc, withFile, withLine := withCallerFunc(f), withCallerFile(f), withCallerLine(f)
	if !(withFunc || withFile || withLine) {
		return
	}

	frames := callers(callersSkip, 1)
	if len(frames) < 1 {
		return
	}

	acc := writeAccumulator(&n, &err)

	// [
	if acc(w.Write(brackL)) {
		return
	}

	// FWithFile
	if withFile {
		file := frames[0].File
		i := strings.LastIndexByte(file, '/')
		if i >= 0 && i < len(file)-1 {
			file = file[i+1:]
		}
		if acc(writeString(w, file)) {
			return
		}
	}

	// FWithLine
	if withLine {
		line := frames[0].Line
		if withFile {
			if acc(w.Write(colon)) {
				return
			}
		}
		if acc(writeString(w, strconv.Itoa(line))) {
			return
		}
	}

	// FWithFunc
	if withFunc {
		if withFile {
			if acc(w.Write(fatArrow)) {
				return
			}
		}

		fn := frames[0].Function
		i := strings.LastIndexByte(fn, '.')
		if i >= 0 && i < len(fn)-1 {
			fn = fn[i+1:]
		}
		if acc(writeString(w, fn)) {
			return
		}
	}

	// ]
	if acc(w.Write(brackR)) {
		return
	}

	if acc(w.Write(space)) {
		return
	}

	return
}

func callers(skip int, count int) []runtime.Frame {
	callerPCs := make([]uintptr, count)
	n := runtime.Callers(skip+2, callerPCs)
	callerFrames := runtime.CallersFrames(callerPCs)

	frames := make([]runtime.Frame, n)
	for i := range n {
		f, _ := callerFrames.Next()
		frames[i] = f
	}

	return frames
}
