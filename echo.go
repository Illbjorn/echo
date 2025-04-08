package echo

import (
	"fmt"
	"os"
)

var sprintf = fmt.Sprintf

/*------------------------------------------------------------------------------
 * Debug
 *----------------------------------------------------------------------------*/

func Debug(v string) (int, error) {
	return log(flags, LevelDebug, v)
}

func Debugf(v string, vs ...any) (int, error) {
	return log(flags, LevelDebug, sprintf(v, vs...))
}

/*------------------------------------------------------------------------------
 * Warn
 *----------------------------------------------------------------------------*/

func Warn(v string) (int, error) {
	return log(flags, LevelWarn, v)
}

func Warnf(v string, vs ...any) (int, error) {
	return log(flags, LevelWarn, sprintf(v, vs...))
}

/*------------------------------------------------------------------------------
 * Info
 *----------------------------------------------------------------------------*/

func Info(v string) (int, error) {
	return log(flags, LevelInfo, v)
}

func Infof(v string, vs ...any) (int, error) {
	return log(flags, LevelInfo, sprintf(v, vs...))
}

/*------------------------------------------------------------------------------
 * Error
 *----------------------------------------------------------------------------*/

func Error(v string) (int, error) {
	return log(flags, LevelError, v)
}

func Errorf(v string, vs ...any) (int, error) {
	return log(flags, LevelError, sprintf(v, vs...))
}

/*------------------------------------------------------------------------------
 * Fatal
 *----------------------------------------------------------------------------*/

func Fatal(v string) {
	log(flags, LevelFatal, v)
	os.Exit(1)
}

func Fatalf(v string, vs ...any) {
	log(flags, LevelFatal, sprintf(v, vs...))
	os.Exit(1)
}

/*------------------------------------------------------------------------------
 * Log
 *----------------------------------------------------------------------------*/

func log(f Flags, l Level, msg string) (n int, err error) {
	if l < level {
		return
	}

	acc := writeAccumulator(&n, &err)

	if acc(writeFlagOpts(writer, f, l)) {
		return
	}

	acc(writeString(writer, msg))

	if acc(writer.Write(newline)) {
		return
	}

	return
}
