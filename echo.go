package echo

func Debug(m string, vs ...any) {
	logf(LevelDebug, m, vs...)
}

func Warn(m string, vs ...any) {
	logf(LevelWarn, m, vs...)
}

func Info(m string, vs ...any) {
	logf(LevelInfo, m, vs...)
}

func Error(m string, vs ...any) {
	logf(LevelInfo, m, vs...)
}

func Fatal(m string, vs ...any) {
	logf(LevelFatal, m, vs...)
}
