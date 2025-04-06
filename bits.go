package echo

func bitSet(i uint8, v uint8) bool {
	shift := 8 - i
	if shift < 0 {
		return false
	}

	return v>>shift&1 == 1
}
