package echo

import (
	"slices"
	"testing"

	"gotest.tools/v3/assert"
)

func TestBitSet(t *testing.T) {
	assert.Check(t, !bitSet(16, 1)) // Test overflow
	assert.Check(t, !bitSet(0, 0))  // Test zeroth bit
	for i := 1; i < 9; i++ {        // Test all other bits
		var v uint8 = 1 << (8 - i)
		var pos uint8 = uint8(i)
		expectBits(t, v, pos)
	}
}

func expectBits(t *testing.T, v uint8, positions ...uint8) {
	for i := 1; i < 9; i++ {
		pos := uint8(i)
		if slices.Contains(positions, pos) {
			assert.Check(t, bitSet(pos, v))
		} else {
			assert.Check(t, !bitSet(pos, v))
		}
	}
}
