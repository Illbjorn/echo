package echo

var level Level

func SetLevel(l Level) {
	level = l
}

type Level uint8

const (
	LevelDebug Level = 1 + iota
	LevelWarn
	LevelInfo
	LevelError
	LevelFatal
)

var (
	prefixDebug = []byte("DBG")
	prefixWarn  = []byte("WRN")
	prefixInfo  = []byte("INF")
	prefixError = []byte("ERR")
	prefixFatal = []byte("FTL")
)

func (l Level) Bytes() []byte {
	switch l {
	case LevelDebug:
		return prefixDebug
	case LevelWarn:
		return prefixWarn
	case LevelInfo:
		return prefixInfo
	case LevelError:
		return prefixError
	case LevelFatal:
		return prefixFatal
	default:
		return nil
	}
}

func (l Level) Color() Color {
	switch l {
	case LevelDebug:
		return BLUE
	case LevelInfo:
		return GREEN
	case LevelWarn:
		return YELLOW
	case LevelError:
		return RED
	case LevelFatal:
		return RED
	default:
		return WHITE
	}
}
