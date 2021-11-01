// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

const (
	BitsNone Bitflag = 0
	BitsAll          = ^BitsNone
)

// NewBits will take Nth number, starting with 1 (not 0).
func NewBits(Nth ...int) Bitflag {
	var out Bitflag
	for _, v := range Nth {
		if v > 0 {
			out = out.Add(1 << (v - 1))
		}
	}
	return out
}

func (f Bitflag) Reverse() Bitflag {
	var ret = uint64(0)
	var power = uint64(63)
	for f != 0 {
		ret += (uint64(f) & 1) << power
		f = f >> 1
		power -= 1
	}
	return Bitflag(ret)
}

type Bitflag uint64

func (f Bitflag) Add(b Bitflag) Bitflag {
	return f | b
}

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

func BitsAdd(b1, b2 uint64) uint64 {
	return b1 | b2
}

func BitsSub(bFrom, bTo uint64) uint64 {
	return bFrom &^ bTo
}

func BitsToggle(bFrom, bToggle uint64) uint64 {
	return bFrom ^ bToggle
}

func BitsAnd(b1, b2 uint64) uint64 {
	return b1 & b2
}

func BitsAny(bFrom, bTo uint64) bool {
	return bFrom&bTo != 0
}

func BitsHas(b1, b2 uint64) bool {
	return b1&b2 == 0
}

