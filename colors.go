package echo

type Color []byte

var (
	RESET   Color = []byte("\033[0m")
	RED     Color = []byte("\033[31m")
	GREEN   Color = []byte("\033[32m")
	YELLOW  Color = []byte("\033[33m")
	BLUE    Color = []byte("\033[34m")
	MAGENTA Color = []byte("\033[35m")
	CYAN    Color = []byte("\033[36m")
	GRAY    Color = []byte("\033[37m")
	WHITE   Color = []byte("\033[97m")
)
