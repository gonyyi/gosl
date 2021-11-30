// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/30/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestIntsJoin(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var a = []int{1, 3, 5, 8, 100, 0, -200, 23}
		buf := make(gosl.Buf, 0, 4<<10)
		buf = buf.WriteBytes('[')
		buf = gosl.IntsJoin(buf, a, ',')
		buf = buf.WriteBytes(']')
		gosl.Test(t, "[1,3,5,8,100,0,-200,23]", buf.String())
	})

	t.Run("no delim", func(t *testing.T) {
		var a = []int{1, 3, 5, 8, 100, 0, -200, 23}
		buf := make(gosl.Buf, 0, 4<<10)
		buf = buf.WriteBytes('[')
		buf = gosl.IntsJoin(buf, a)
		buf = buf.WriteBytes(']')
		gosl.Test(t, "[13581000-20023]", buf.String())
	})

	t.Run("null", func(t *testing.T) {
		var a = []int{}
		buf := make(gosl.Buf, 0, 4<<10)
		buf = buf.WriteBytes('[')
		buf = gosl.IntsJoin(buf, a, ',')
		buf = buf.WriteBytes(']')
		gosl.Test(t, "[]", buf.String())
	})

}

func BenchmarkIntsJoin(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 4<<10)
		test := []int{1, 3, 5, 7}
		for i := 0; i < b.N; i++ {
			test[3] = i
			buf = gosl.IntsJoin(buf.Reset(), test, ',', ' ')
		}
		// buf.Println()
	})
}
