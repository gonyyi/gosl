package time64_test

import (
	"github.com/gonyyi/gosl"
	"github.com/gonyyi/gosl/time64"
	"testing"
)

func TestIntTime(t *testing.T) {
	t1 := time64.IntTime(20060102150405)
	gosl.Test(t, "2006/01/02 15:04:05", t1.String())

	t2 := t1.
		SetDate(2020, 10, 31).
		SetTime(5, 6, 0)
	gosl.Test(t, "2020/10/31 05:06:00", t2.String())

	gosl.Test(t, false, t1 > t2)
	gosl.Test(t, true, t1 < t2)

	t2 = t2.
		SetDate(2006, 1, 2).
		SetTime(15, 4, 5)
	gosl.Test(t, true, t1 == t2)

	t2, ok := t2.Parse("1981/10/02 09:10:11")
	gosl.Test(t, true, ok)
	gosl.Test(t, true, int64(t2) == 19811002091011)

	gosl.Test(t, true, t1 > t2)
	gosl.Test(t, false, t1 < t2)
	gosl.Test(t, false, t1 == t2)
}
