package echo

import (
	"fmt"
	"os"
)

func write(v []byte) error {
  var _, err = wr.Write(v)
  return err
}

func log(l Level, m string) error {
	if l < level {
		return nil
	}

	var out = format([]byte(m), l, flags)
	out = append(out, bNewline)
	var err error
  if err = write(out); err != nil {
    return err
  }

	if l == LevelFatal {
		os.Exit(1)
	}

  return nil
}

func logf(l Level, m string, vs ...any) error {
	if l < level {
		return nil
	}

	return log(l, fmt.Sprintf(m, vs...))
}
