package echo

import (
	"io"
	"strconv"
	"time"
)

var flags Flags

type Flags = uint8

func SetFlags(fs ...Flags) {
	flags = 0
	for _, f := range fs {
		flags |= 1 << (8 - f)
	}
}

const (
	//                                  Bit
	//                                  ---
	WithCallerFunc Flags = 1 + iota //  1
	WithCallerFile                  //  2
	WithCallerLine                  //  3
	WithCallStack                   //  4
	WithLevel                       //  5
	WithTime                        //  6
	WithDate                        //  7
	WithColor                       //  8
)

func withCallerFunc(f Flags) bool { return bitSet(WithCallerFunc, f) }
func withCallerFile(f Flags) bool { return bitSet(WithCallerFile, f) }
func withCallerLine(f Flags) bool { return bitSet(WithCallerLine, f) }
func withCallStack(f Flags) bool  { return bitSet(WithCallStack, f) } // TODO
func withLevel(f Flags) bool      { return bitSet(WithLevel, f) }
func withTime(f Flags) bool       { return bitSet(WithTime, f) }
func withDate(f Flags) bool       { return bitSet(WithDate, f) }
func withColor(f Flags) bool      { return bitSet(WithColor, f) }

func writeFlagOpts(w io.Writer, f Flags, l Level) (n int, err error) {
	acc := writeAccumulator(&n, &err)

	//////////////////////////////////////////////////////////////////////////////
	// Callers

	if acc(writeCallers(w, f)) {
		return
	}

	//////////////////////////////////////////////////////////////////////////////
	// Level

	if acc(writeLevel(w, f, l)) {
		return
	}

	//////////////////////////////////////////////////////////////////////////////
	// Date

	if acc(writeTimestamp(w, f, l)) {
		return
	}

	//////////////////////////////////////////////////////////////////////////////
	// Call Stack
	// TODO

	return
}

func writeTimestamp(w io.Writer, f Flags, _ Level) (n int, err error) {
	withTime := withTime(f)
	withDate := withDate(f)
	if !(withTime || withDate) {
		return
	}

	acc := writeAccumulator(&n, &err)

	now := time.Now()
	if withDate {
		// [
		if acc(w.Write(brackL)) {
			return
		}

		// 03
		month := strconv.Itoa(int(now.Month()))
		if acc(writeDouble(w, month)) {
			return
		}

		// -
		if acc(w.Write(dash)) {
			return
		}

		// 19
		day := strconv.Itoa(now.Day())
		if acc(writeDouble(w, day)) {
			return
		}

		// -
		if acc(w.Write(dash)) {
			return
		}

		// 25
		year := strconv.Itoa(now.Year())
		if acc(writeString(w, year[2:])) {
			return
		}

		// ]
		if acc(w.Write(brackR)) {
			return
		}
	}

	if withTime {
		// [
		if acc(w.Write(brackL)) {
			return
		}

		// 00
		hours := now.Hour()
		if hours > 12 {
			hours -= 12
		}
		hourStr := strconv.Itoa(hours)
		if acc(writeDouble(w, hourStr)) {
			return
		}

		// :
		if acc(w.Write(colon)) {
			return
		}

		// 00
		min := strconv.Itoa(now.Minute())
		if acc(writeDouble(w, min)) {
			return
		}

		// :
		if acc(w.Write(colon)) {
			return
		}

		// 00
		sec := strconv.Itoa(now.Second())
		if acc(writeDouble(w, sec)) {
			return
		}

		// ]
		if acc(w.Write(brackR)) {
			return
		}
	}

	if acc(w.Write(space)) {
		return
	}

	return
}
