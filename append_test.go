// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Append_Path(t *testing.T) {
	var tmp []byte
	tmp = gosl.AppendPath(tmp, "/aaa", "bbb")
	gosl.Test(t, "/aaa/bbb", string(tmp))
	tmp = gosl.AppendPath(tmp, "", "d", "e")
	gosl.Test(t, "/aaa/bbb/d/e", string(tmp))
	tmp = gosl.AppendPath(tmp[:0], "/aaa/", "/bbb/")
	gosl.Test(t, "/aaa/bbb/", string(tmp))
}

func Benchmark_Append_Path(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		b.ReportAllocs()
		var tmp []byte
		for i := 0; i < b.N; i++ {
			// tmp = tmp[:0]
			tmp = gosl.AppendPath(tmp[:0], "/aaa", "bbb", "ccc", "/ddd")
		}
		// println(string(tmp))
	})
}

func Test_Append_AppendFill(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendFill(a, nil, 10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("negative-n", func(t *testing.T) {
		a := []byte("gon")
		b := []byte("123")
		a = gosl.AppendFill(a, b, -10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("basic", func(t *testing.T) {
		a := []byte("gon")
		b := []byte("123")
		a = gosl.AppendFill(a, b, 10)
		gosl.Test(t, "gon1231231231", string(a))
	})
}

func Test_Append_AppendRepeat(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendRepeat(a, nil, 10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("negative-n", func(t *testing.T) {
		a := []byte("gon")
		b := []byte("123")
		a = gosl.AppendRepeat(a, b, -10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("basic", func(t *testing.T) {
		a := []byte("gon")
		b := []byte("123")
		a = gosl.AppendRepeat(a, b, 2)
		gosl.Test(t, "gon123123", string(a))
	})

	t.Run("AppendRepeatByte-basic", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendRepeatByte(a, ':', 2)
		gosl.Test(t, "gon::", string(a))
	})

	t.Run("AppendRepeatByte-nil", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendRepeatByte(a, ':', -1)
		gosl.Test(t, "gon", string(a))
	})

}


func Benchmark_AppendRepeat(b *testing.B) {
	b.Run("AppendRepeat", func(b *testing.B) {
		b.ReportAllocs()
		var buf = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = gosl.AppendRepeat(buf, []byte("abc"), 10)
		}
		// println(buf.String())
	})
	b.Run("AppendRepeatByte", func(b *testing.B) {
		b.ReportAllocs()
		var buf = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = gosl.AppendRepeatByte(buf, '-', 10)
		}
		// println(buf.String())
	})
}
