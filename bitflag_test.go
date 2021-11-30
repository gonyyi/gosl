// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/30/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestNewBitflag(t *testing.T) {
	const (
		B1 gosl.Bitflag = 1 << iota
		B2
		B3
		B4
	)
	buf := make(gosl.Buf, 0, 1024)
	f := func(b gosl.Bitflag) string {
		return string(b.Output(nil))
	}

	gosl.Test(t,
		f(gosl.Bitflag(1).Nth(2, 4, 6, 8, 10)),
		"11010101010000000000000000000000")

	bf := gosl.NewBitflag()
	gosl.Test(t, 0, int(bf))

	bf = bf.All()
	gosl.Test(t, true, int(bf) != int(bf.None()))         // All != None
	gosl.Test(t, true, int(bf)-int(bf.None()) == int(bf)) // All - None == All
	gosl.Test(t, true, int(bf.Reverse()) == int(bf))      // All().Reverse() == All() as this is reversed order 111.reverse() == 111

	bf = bf.None().Nth(1)
	gosl.Test(t, true, int(bf) != int(bf.Reverse()))
	gosl.Test(t, true, int(bf) == int(bf.Reverse().Reverse()))

	buf = buf.Reset()
	buf = bf.Output(buf)
	gosl.Test(t, "10000000000000000000000000000000", buf.String())
	buf = bf.Reverse().Output(buf.Reset())
	gosl.Test(t, "00000000000000000000000000000001", buf.String())

	bf2 := gosl.NewBitflag()
	gosl.Test(t, 0, int(bf2.Sub(bf)))
	gosl.Test(t, 1, int(bf.Sub(bf2)))
	gosl.Test(t, 0, int(bf.Sub(bf)))

	bf = bf.All()
	buf = bf.Output(buf.Reset())
	gosl.Test(t, "11111111111111111111111111111111", buf.String())

	bf = B1.Add(B2).Add(B3).Add(B4)
	buf = bf.Output(buf.Reset())
	gosl.Test(t, "11110000000000000000000000000000", buf.String())

	bf = bf.Toggle(B2)
	buf = bf.Output(buf.Reset())
	gosl.Test(t, "10110000000000000000000000000000", buf.String())

	bf = bf.Toggle(B2)
	buf = bf.Output(buf.Reset())
	gosl.Test(t, "11110000000000000000000000000000", buf.String())

	bf = bf.Toggle(B2 | B3)
	buf = bf.Output(buf.Reset())
	gosl.Test(t, "10010000000000000000000000000000", buf.String())

	gosl.Test(t, false, bf.Any(B2))
	gosl.Test(t, false, bf.Any(B3))
	gosl.Test(t, false, bf.Any(B2|B3))
	gosl.Test(t, true, bf.Any(B2|B3|B4))

	gosl.Test(t, false, bf.Has(B2|B3|B4))
	gosl.Test(t, true, bf.Has(B1|B4))

	gosl.Test(t, int(B1|B2), int(gosl.BitsAdd(uint32(B1), uint32(B2))))
	gosl.Test(t, int(B1|B3|B4), int(gosl.BitsSub(uint32(B1|B2|B3|B4), uint32(B2))))
	gosl.Test(t, int(B1), int(gosl.BitsToggle(uint32(B1|B2|B3|B4), uint32(B2|B3|B4))))
	gosl.Test(t, int(B1), int(gosl.BitsAnd(uint32(B1), uint32(B1|B4))))

	gosl.Test(t, true, gosl.BitsAny(uint32(B1|B2|B3), uint32(B2|B4)))
	gosl.Test(t, false, gosl.BitsAny(uint32(B1|B2), uint32(B4)))
	gosl.Test(t, false, gosl.BitsAny(uint32(B1|B2), uint32(B3|B4)))
	gosl.Test(t, true, gosl.BitsAny(uint32(B1|B2), uint32(B2|B3|B4)))

	gosl.Test(t, false, gosl.BitsHas(uint32(B1|B2), uint32(B2|B3|B4)))
	gosl.Test(t, false, gosl.BitsHas(uint32(B1|B2), uint32(B1|B2|B3|B4)))
	gosl.Test(t, true, gosl.BitsHas(uint32(B1|B2), uint32(B1|B2)))
	gosl.Test(t, true, gosl.BitsHas(uint32(B1|B2|B3|B4), uint32(B1|B2)))

	gosl.Test(t, 0, int(B1.And(B2)))
	gosl.Test(t, 1, int((B1|B2).And(B1)))
	gosl.Test(t, 2, int((B1|B2).And(B2)))
}
