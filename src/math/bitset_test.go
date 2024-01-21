package math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinomial(t *testing.T) {
	b := binomialInternalHelper()

	str := ""
	for i := uint64(0); i < 64; i++ {
		str += "{\n"
		for j := uint64(0); j <= i; j++ {
			str += fmt.Sprintf("%d,\n", b[i][j])
		}
		str += "},\n"
	}

	assert.Equal(t, b, binomial, str)
}

func TestBinomialCount(t *testing.T) {
	assert.Equal(t, binomialCountInternalHelper(), binomialCount)
}

func TestBitsetCount(t *testing.T) {
	assert.Equal(t, uint64(1), CountBitsets(0))
	assert.Equal(t, uint64(64), CountBitsets(1))
	assert.Equal(t, uint64(2016), CountBitsets(2))
	assert.Equal(t, uint64(41664), CountBitsets(3))
	assert.Equal(t, uint64(635376), CountBitsets(4))
	assert.Equal(t, uint64(7624512), CountBitsets(5))
	assert.Equal(t, uint64(74974368), CountBitsets(6))
	assert.Equal(t, uint64(621216192), CountBitsets(7))
	assert.Equal(t, uint64(4426165368), CountBitsets(8))
	assert.Equal(t, uint64(27540584512), CountBitsets(9))
	assert.Equal(t, uint64(151473214816), CountBitsets(10))
	assert.Equal(t, uint64(743595781824), CountBitsets(11))
	assert.Equal(t, uint64(3284214703056), CountBitsets(12))
	assert.Equal(t, uint64(13136858812224), CountBitsets(13))
	assert.Equal(t, uint64(47855699958816), CountBitsets(14))
	assert.Equal(t, uint64(159518999862720), CountBitsets(15))
	assert.Equal(t, uint64(488526937079580), CountBitsets(16))
	assert.Equal(t, uint64(1379370175283520), CountBitsets(17))
	assert.Equal(t, uint64(3601688791018080), CountBitsets(18))
	assert.Equal(t, uint64(8719878125622720), CountBitsets(19))
	assert.Equal(t, uint64(19619725782651120), CountBitsets(20))
	assert.Equal(t, uint64(41107996877935680), CountBitsets(21))
	assert.Equal(t, uint64(80347448443237920), CountBitsets(22))
	assert.Equal(t, uint64(146721427591999680), CountBitsets(23))
	assert.Equal(t, uint64(250649105469666120), CountBitsets(24))
	assert.Equal(t, uint64(401038568751465792), CountBitsets(25))
	assert.Equal(t, uint64(601557853127198688), CountBitsets(26))
	assert.Equal(t, uint64(846636978475316672), CountBitsets(27))
	assert.Equal(t, uint64(1118770292985239888), CountBitsets(28))
	assert.Equal(t, uint64(1388818294740297792), CountBitsets(29))
	assert.Equal(t, uint64(1620288010530347424), CountBitsets(30))
	assert.Equal(t, uint64(1777090076065542336), CountBitsets(31))
	assert.Equal(t, uint64(1832624140942590534), CountBitsets(32))
	assert.Equal(t, uint64(1777090076065542336), CountBitsets(33))
	assert.Equal(t, uint64(1620288010530347424), CountBitsets(34))
	assert.Equal(t, uint64(1388818294740297792), CountBitsets(35))
	assert.Equal(t, uint64(1118770292985239888), CountBitsets(36))
	assert.Equal(t, uint64(846636978475316672), CountBitsets(37))
	assert.Equal(t, uint64(601557853127198688), CountBitsets(38))
	assert.Equal(t, uint64(401038568751465792), CountBitsets(39))
	assert.Equal(t, uint64(250649105469666120), CountBitsets(40))
	assert.Equal(t, uint64(146721427591999680), CountBitsets(41))
	assert.Equal(t, uint64(80347448443237920), CountBitsets(42))
	assert.Equal(t, uint64(41107996877935680), CountBitsets(43))
	assert.Equal(t, uint64(19619725782651120), CountBitsets(44))
	assert.Equal(t, uint64(8719878125622720), CountBitsets(45))
	assert.Equal(t, uint64(3601688791018080), CountBitsets(46))
	assert.Equal(t, uint64(1379370175283520), CountBitsets(47))
	assert.Equal(t, uint64(488526937079580), CountBitsets(48))
	assert.Equal(t, uint64(159518999862720), CountBitsets(49))
	assert.Equal(t, uint64(47855699958816), CountBitsets(50))
	assert.Equal(t, uint64(13136858812224), CountBitsets(51))
	assert.Equal(t, uint64(3284214703056), CountBitsets(52))
	assert.Equal(t, uint64(743595781824), CountBitsets(53))
	assert.Equal(t, uint64(151473214816), CountBitsets(54))
	assert.Equal(t, uint64(27540584512), CountBitsets(55))
	assert.Equal(t, uint64(4426165368), CountBitsets(56))
	assert.Equal(t, uint64(621216192), CountBitsets(57))
	assert.Equal(t, uint64(74974368), CountBitsets(58))
	assert.Equal(t, uint64(7624512), CountBitsets(59))
	assert.Equal(t, uint64(635376), CountBitsets(60))
	assert.Equal(t, uint64(41664), CountBitsets(61))
	assert.Equal(t, uint64(2016), CountBitsets(62))
	assert.Equal(t, uint64(64), CountBitsets(63))
	assert.Equal(t, uint64(1), CountBitsets(64))
}

func TestIndexToBitset(t *testing.T) {
	// Basic Functionality Test
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000000000001), IndexToBitset(1, 0))
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000000000010), IndexToBitset(1, 1))
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000011000001), IndexToBitset(3, 300))

	// Boundary Conditions
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000000000000), IndexToBitset(0, 0)) // n=0
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000000000011), IndexToBitset(2, 0)) // Smallest n > 1
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000000000101), IndexToBitset(2, 1)) // n=2, index=1

	// Testing with larger index values
	assert.Equal(t, uint64(0b1000000000000000000000000000000000000000000000000000000000000000), IndexToBitset(1, 63)) // Largest index for n=1
	assert.Equal(t, ^uint64(0), IndexToBitset(64, 0))                                                                 // n=64 (all bits set)
}

func TestBitsetToIndex(t *testing.T) {
	// Basic Functionality Test
	assert.Equal(t, uint64(0), BitsetToIndex(1, 0b0000000000000000000000000000000000000000000000000000000000000001))
	assert.Equal(t, uint64(1), BitsetToIndex(1, 0b0000000000000000000000000000000000000000000000000000000000000010))
	assert.Equal(t, uint64(300), BitsetToIndex(3, 0b0000000000000000000000000000000000000000000000000000000011000001))

	// Boundary Conditions
	assert.Equal(t, uint64(0), BitsetToIndex(0, 0b0000000000000000000000000000000000000000000000000000000000000000)) // n=0
	assert.Equal(t, uint64(0), BitsetToIndex(2, 0b0000000000000000000000000000000000000000000000000000000000000011)) // Smallest n > 1
	assert.Equal(t, uint64(1), BitsetToIndex(2, 0b0000000000000000000000000000000000000000000000000000000000000101)) // n=2, bitset=101

	// Testing with larger bitset values
	assert.Equal(t, uint64(63), BitsetToIndex(1, 0b1000000000000000000000000000000000000000000000000000000000000000)) // Largest bitset for n=1
	assert.Equal(t, uint64(0), BitsetToIndex(64, ^uint64(0)))                                                         // n=64 (all bits set)
}

func TestNextValidIndex(t *testing.T) {
	nextIndex := NextValidIndex(3, 300, IndexToBitset(3, 300), 6)
	assert.Equal(t, uint64(356), nextIndex)
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000110000001), IndexToBitset(3, nextIndex+1))

	nextIndex = NextValidIndex(3, 300, IndexToBitset(3, 300), 7)
	assert.Equal(t, uint64(300), nextIndex)
	assert.Equal(t, uint64(0b0000000000000000000000000000000000000000000000000000000101000001), IndexToBitset(3, nextIndex+1))

	assert.Equal(t, uint64(0b1101000000000000000000000000000000000000000000000000000000000000), IndexToBitset(3, 41662))

	nextIndex = NextValidIndex(3, 62, IndexToBitset(3, 41662), 62)
	assert.Equal(t, uint64(41662), nextIndex)
}

// func TestPrint(t *testing.T) {
// 	for i := uint64(0); i < 1000; i++ {
// 		t.Logf("%d: %064b", i, IndexToBitset(4, i))
// 	}
// 	assert.Fail(t, "Test not implemented")
// }
