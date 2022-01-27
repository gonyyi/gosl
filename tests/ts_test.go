// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/19/2022

package gosl_test

import (
	"encoding/binary"
	"github.com/gonyyi/gosl"
	"testing"
)

func TestTS(t *testing.T) {
	var tsz = gosl.TS(0)

	eqTS := func(id int, exp, act gosl.TS) {
		if exp != act {
			t.Errorf("[%04d] Failed: Exp=%d, Act=%d\n", id, exp, act)
		}
	}
	eqI64 := func(id int, exp, act int64) {
		eqTS(id, gosl.TS(exp), gosl.TS(act))
	}
	eqStr := func(id int, exp, act string) {
		if exp != act {
			t.Errorf("[%04d] Failed: Exp=%s, Act=%s\n", id, exp, act)
		}
	}
	eqBool := func(id int, exp, act bool) {
		if exp != act {
			t.Errorf("[%04d] Failed: Exp=%t, Act=%t\n", id, exp, act)
		}
	}
	t.Run("Date(),Time(),MS()", func(t *testing.T) {
		tmp := tsz.Parse("20060102150405123", 0)
		eqI64(1010, 20060102, tmp.Date())
		eqI64(1110, 150405, tmp.Time())
		eqI64(1210, 123, tmp.MS())
	})

	t.Run("Valid()", func(t *testing.T) {
		eqBool(1010, true, gosl.TS(20060102150405123).IsValid())
		eqBool(1010, false, gosl.TS(2006010215040512).IsValid())
	})

	t.Run("SetDate(),SetTime()", func(t *testing.T) {
		tmp := tsz.Parse("20060102150405123", 0)
		eqTS(1010, 19811002150405123, tmp.SetDate(1981, 10, 2))
		eqTS(1110, 20060102091011123, tmp.SetTime(9, 10, 11))
	})

	t.Run("Parse()", func(t *testing.T) {
		eqTS(1010, 20060102150405000, tsz.Parse("20060102150405", 0))
		eqTS(1020, 20060102150405123, tsz.Parse("20060102150405123", 0))
		eqTS(1030, 20060102150405000, tsz.Parse("2006/01/02 15:04:05", 0))
		eqTS(1040, 20060102150405123, tsz.Parse("2006/01/02 15:04:05.123", 0))
		eqTS(1041, 20060102150405000, tsz.Parse("2006-01-02-15-04-05", 0))
		eqTS(1042, 20060102150405123, tsz.Parse("2006-01-02-15-04-05-123", 0))
		eqTS(2050, 0, tsz.Parse("20062102150405", 0))
		eqTS(2060, 0, tsz.Parse("20060132150405123", 0))
		eqTS(2070, 0, tsz.Parse("2006/01/02 24:04:05", 0))
		eqTS(2080, 0, tsz.Parse("2006/01/02 15:60:05.123", 0))
		eqTS(2080, 0, tsz.Parse("2006/01/02 15:04:60.123", 0))
	})

	t.Run("Byte(),ParseByte()", func(t *testing.T) {
		tmp := tsz.Parse("20060102150405123", 0)
		if int(tmp) == 0 {
			t.Fail()
		}
		// byte() with timestamp, parsebyte with binary func
		{
			out := tmp.Byte()
			res := binary.LittleEndian.Uint64(out[:]) // parse [8]byte to int
			eqTS(100, tmp, gosl.TS(res))
		}
		// byte using binary, then parse it with TS
		{
			out := make([]byte, 8)
			binary.LittleEndian.PutUint64(out, uint64(tmp)) // create uint64 to byte
			var out2 [8]byte                                // TS takes array, but binary library returns slice
			copy(out2[:], out[:])
			eqTS(200, 20060102150405123, gosl.TS(0).ParseByte(out2))
		}
	})

	t.Run("String()", func(t *testing.T) {
		eqStr(1010, "20060102150405000", tsz.Parse("20060102150405", 0).String())
		eqStr(1020, "20060102150405123", tsz.Parse("20060102150405123", 0).String())
		eqStr(1030, "20060102150405000", tsz.Parse("2006/01/02 15:04:05", 0).String())
		eqStr(1040, "20060102150405123", tsz.Parse("2006/01/02 15:04:05.123", 0).String())
		eqStr(1050, "20060102150405000", tsz.Parse("2006-01-02 15-04-05", 0).String())
		eqStr(1060, "20060102150405123", tsz.Parse("2006-01-02-15-04-05-123", 0).String())
	})

	t.Run("Format()", func(t *testing.T) {
		eqStr(1010, "2006/01/02 15:04:05.000", tsz.Parse("20060102150405", 0).Format())
		eqStr(1020, "2006/01/02 15:04:05.123", tsz.Parse("20060102150405123", 0).Format())
		eqStr(1030, "2006/01/02 15:04:05.000", tsz.Parse("2006/01/02 15:04:05", 0).Format())
		eqStr(1040, "2006/01/02 15:04:05.123", tsz.Parse("2006/01/02 15:04:05.123", 0).Format())
		eqStr(1050, "2006/01/02 15:04:05.000", tsz.Parse("2006-01-02 15-04-05", 0).Format())
		eqStr(1060, "2006/01/02 15:04:05.123", tsz.Parse("2006-01-02-15-04-05-123", 0).Format())
	})
}
