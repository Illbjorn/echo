package echo

func bitSet(bitPosition uint8, value uint8) bool {
	shift := 8 - bitPosition
	if shift > 8 {
		return false
	}
	return value>>shift&1 == 1
}
