// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/19/2022

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestIntTime(t *testing.T) {
	t1 := gosl.Timestamp(20060102150405000)
	gosl.Test(t, "2006/01/02 15:04:05.000", t1.String())

	t2 := t1.
		SetDate(2020, 10, 31).
		SetTime(5, 6, 0)
	gosl.Test(t, "2020/10/31 05:06:00.000", t2.String())

	gosl.Test(t, false, t1 > t2)
	gosl.Test(t, true, t1 < t2)

	t2 = t2.
		SetDate(2006, 1, 2).
		SetTime(15, 4, 5)
	gosl.Test(t, true, t1 == t2)

	t2, ok := t2.Parse("1981/10/02 09:10:11")
	gosl.Test(t, true, ok)
	gosl.Test(t, true, int64(t2) == 19811002091011000)
	gosl.Test(t, true, t1 > t2)

	gosl.Test(t, false, t1 < t2)
	gosl.Test(t, false, t1 == t2)
}

func BenchmarkTimestamp_Parse(b *testing.B) {
	b.Run("b1", func(b *testing.B) {
		b.ReportAllocs()
		ss := []string{"2006/01/02 15:04:05", "2006/01/02 15:04:05.000", "2006/01/02 19:04:05", "2006/01/02 15:04:54"}
		var ts gosl.Timestamp
		var ok bool
		_ = ok
		for i := 0; i < b.N; i++ {
			ts, ok = ts.Parse(ss[i%4])
		}
	})
}
