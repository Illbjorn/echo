package echo

func format(m []byte, l Level, f Flag) []byte {
	var ret []byte

	//////////////////////////////////////////////////////////////////////////////
	// FLAGS
	//
	// Colorize output
	if f&WITH_COLOR == WITH_COLOR {
		ret = append(ret, []byte(l.Color())...)
	}

	// Indent output based on callstack depth
	if f&WITH_TIER == WITH_TIER {
		for range stackDepth() {
			ret = append(ret, bSpace)
		}
	}

	// Log with log level
	if f&WITH_LEVEL == WITH_LEVEL {
		// Produce level
		ret = append(ret, l.Bytes()...)
		ret = append(ret, bSpace)
	}

	// Log with caller (file:line)
	if f&WITH_CALLER == WITH_CALLER {
		// Produce caller
		ret = append(ret, caller(5)...)
		ret = append(ret, bSpace)
	}

	// Reset color if colorized output
	if f&WITH_COLOR == WITH_COLOR {
		ret = append(ret, []byte(RESET)...)
	}

	//////////////////////////////////////////////////////////////////////////////
	// Append message
	ret = append(ret, m...)

	return ret
}
