package echo

import "io"

var level Level = LevelWarn

func SetLevel(l Level) Level {
	// Not allowed to set above Fatal!
	if l >= LevelFatal {
		return level
	}

	current := level
	level = l
	return current
}

type Level = uint8

const (
	LevelDebug Level = 1 + iota
	LevelWarn
	LevelInfo
	LevelError
	LevelFatal
)

var (
	lDebug = []byte("DBG")
	lWarn  = []byte("WRN")
	lInfo  = []byte("INF")
	lError = []byte("ERR")
	lFatal = []byte("FTL")
)

func writeLevel(w io.Writer, f Flags, l Level) (n int, err error) {
	if !withLevel(f) {
		return 0, nil
	}

	fwrite := func(nn int, e error) bool {
		n += nn
		if e != nil {
			err = e
			return true
		}
		return false
	}

	var color []byte
	var level []byte
	switch l {
	case LevelDebug:
		color = colorGray
		level = lDebug
	case LevelWarn:
		color = colorYellow
		level = lWarn
	case LevelInfo:
		color = colorGreen
		level = lInfo
	case LevelError:
		color = colorRed
		level = lError
	case LevelFatal:
		color = colorRed
		level = lFatal
	default:
		panic("impossible")
	}

	if withColor(f) && fwrite(w.Write(color)) {
		return
	}

	if fwrite(w.Write(level)) {
		return
	}

	if withColor(f) && fwrite(w.Write(colorReset)) {
		return
	}

	if fwrite(w.Write(space)) {
		return
	}

	return
}
