// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/3/2021

package gosl

// NewBitflag returns 0 of bitflag
func NewBitflag() uint32 {
	return 0
}

// Bitflag is uint32 for bigflags
type Bitflag uint32

// All will turn all bitflag on
func (f Bitflag) All() Bitflag {
	return ^f.None()
}

// None will turn all bitflag off
func (f Bitflag) None() Bitflag {
	return 0
}

// Nth takes location of bits (0-31 since uint32) and add those.
// Bitflag(0)                 // 00000000000000000000000000000000
// Bitflag(0).Nth(1,3,5,7,9)  // 10101010100000000000000000000000
// Bitflag(0).Nth(2,4,6,8,10) // 01010101010000000000000000000000
func (f Bitflag) Nth(Nth ...uint8) Bitflag {
	for _, v := range Nth {
		if v > 0 {
			f = f.Add(1 << (v - 1))
		}
	}
	return f
}

// Reverse wil flip left to right
// Original: 01000000000000000000000000000000
// Reversed: 00000000000000000000000000000010
func (f Bitflag) Reverse() Bitflag {
	var ret = uint32(0)
	var power = uint32(31)
	for f != 0 {
		ret += (uint32(f) & 1) << power
		f = f >> 1
		power -= 1
	}
	return Bitflag(ret)
}

// Add will add `b` to the bitflag
func (f Bitflag) Add(b Bitflag) Bitflag {
	return f | b
}

// Sub will remove `b` from the bitflag
func (f Bitflag) Sub(b Bitflag) Bitflag {
	return f &^ b
}

func (f Bitflag) Toggle(b Bitflag) Bitflag {
	return f ^ b
}

func (f Bitflag) And(b Bitflag) Bitflag {
	return f & b
}

func (f Bitflag) Any(b Bitflag) bool {
	return f&b != 0
}

func (f Bitflag) Has(b Bitflag) bool {
	return f&b == b
}

// Output will take a byte slice and output representation of bitflag
// eg. 1 --> 10000000000000000000000000000000
//     2 --> 01000000000000000000000000000000
//     3 --> 11000000000000000000000000000000
//     4 --> 00100000000000000000000000000000
func (f Bitflag) Output(dst []byte) []byte {
	buf := make(Buf, 0, 32)
	for i := 0; i < 32; i++ {
		d, r := f/2, f%2
		buf = buf.WriteInt(int(r))
		f = d // update bit with newer value
	}
	return append(dst, buf...)
}

func BitsAdd(b1, b2 uint32) uint32 {
	return b1 | b2
}

func BitsSub(bFrom, bTo uint32) uint32 {
	return bFrom &^ bTo
}

func BitsToggle(bFrom, bToggle uint32) uint32 {
	return bFrom ^ bToggle
}

func BitsAnd(b1, b2 uint32) uint32 {
	return b1 & b2
}

func BitsAny(bFrom, bTo uint32) bool {
	return bFrom&bTo != 0
}

func BitsHas(b1, b2 uint32) bool {
	return b1&b2 == 0
}
