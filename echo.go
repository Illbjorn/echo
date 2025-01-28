package echo

func Debug(m string, vs ...any) error {
	return logf(LevelDebug, m, vs...)
}

func Warn(m string, vs ...any) error {
	return logf(LevelWarn, m, vs...)
}

func Info(m string, vs ...any) error {
	return logf(LevelInfo, m, vs...)
}

func Error(m string, vs ...any) error {
	return logf(LevelInfo, m, vs...)
}

func Fatal(m string, vs ...any) error {
	return logf(LevelFatal, m, vs...)
}
