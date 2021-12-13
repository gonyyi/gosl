// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/13/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)

func BenchmarkRollingBuffer(b *testing.B) {
	w := gosl.NewRollingBuffer(3)
	ba := [][]byte{
		[]byte("abc-1"),
		[]byte("abc-2"),
		[]byte("abc-3"),
		[]byte("abc-4"),
		[]byte("abc-5"),
	}

	b.ReportAllocs()
	max := 0
	for i := 0; i < b.N; i++ {
		w.Write(ba[i%5])
		max = i
	}
	_ = max

	buf := make(gosl.Buf, 0, 1024)
	buf = w.Last(buf, 5)
	// println(len(buf), ":", max+1, "\n"+buf.String())
}

func TestRollingBuffer(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		w := gosl.NewRollingBuffer(3)
		w.NewLine = true
		buf := make(gosl.Buf, 0, 2048)
		buf = w.Last(buf.Reset(), 10)
		gosl.Test(t, "", buf.String())

		w.Write([]byte("OK-1"))
		w.Write([]byte("OK-2"))
		buf = w.Last(buf.Reset(), 10)
		gosl.Test(t, "OK-1\nOK-2\n", buf.String())

		w.Write([]byte("OK-3"))
		w.Write([]byte("OK-4"))
		buf = w.Last(buf.Reset(), 10)
		gosl.Test(t, "OK-2\nOK-3\nOK-4\n", buf.String())

		w.Write([]byte("OK-5"))
		buf = w.Last(buf.Reset(), 10)
		gosl.Test(t, "OK-3\nOK-4\nOK-5\n", buf.String())

		buf = w.Last(buf.Reset(), 0)
		gosl.Test(t, "OK-3\nOK-4\nOK-5\n", buf.String())
	})

	t.Run("t2", func(t *testing.T) {
		w := gosl.NewRollingBuffer(3)

		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "[%02d] OK-%d\n", i, i)
		}

		buf := make(gosl.Buf, 0, 2048)
		buf = w.Last(buf.Reset(), 2)
		gosl.Test(t, "[07] OK-7\n[08] OK-8\n", buf.String())
		buf = w.Last(buf.Reset(), 3)
		gosl.Test(t, "[07] OK-7\n[08] OK-8\n[09] OK-9\n", buf.String())

		buf = w.Last(buf.Reset(), 0)
		gosl.Test(t, "[07] OK-7\n[08] OK-8\n[09] OK-9\n", buf.String())

		w.Reset(2)
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "[%02d] OK-%d\n", i, i)
		}
		buf = w.Last(buf.Reset(), 10)
		gosl.Test(t, "[08] OK-8\n[09] OK-9\n", buf.String())

		for i := 10; i < 20; i++ {
			fmt.Fprintf(w, "[%02d] OK-%d\n", i, i)
		}
		buf = w.Last(buf.Reset(), 0)
		gosl.Test(t, "[18] OK-18\n[19] OK-19\n", buf.String())
	})
}

func TestNewRollingIndex(t *testing.T) {
	ri := gosl.NewRollingIndex(3)
	ri = ri.Next()
	gosl.Test(t, 0, ri.Curr())
	gosl.Test(t, "[0]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 1, ri.Curr())
	gosl.Test(t, "[0 1]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 2, ri.Curr())
	gosl.Test(t, "[0 1 2]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 0, ri.Curr())
	gosl.Test(t, "[1 2 0]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 1, ri.Curr())
	gosl.Test(t, "[2 0 1]", fmt.Sprint(ri.List()))

	ri = ri.Next()
	gosl.Test(t, 2, ri.Curr())
	gosl.Test(t, "[0 1 2]", fmt.Sprint(ri.List()))
}
