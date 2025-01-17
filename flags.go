package echo

var (
	flags Flag
)

func SetFlags(f Flag) {
	flags = f
}

type Flag uint8

const (
	WITH_LEVEL Flag = 1 << iota
	WITH_CALLER
	WITH_TIER
	WITH_COLOR
)
