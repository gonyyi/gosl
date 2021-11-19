// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestAppendPath(t *testing.T) {
	var tmp []byte
	tmp = gosl.AppendPath(tmp, "/aaa", "bbb")
	gosl.Test(t, "/aaa/bbb", string(tmp))
	tmp = gosl.AppendPath(tmp, "", "d", "e")
	gosl.Test(t, "/aaa/bbb/d/e", string(tmp))
	tmp = gosl.AppendPath(tmp[:0], "/aaa/", "/bbb/")
	gosl.Test(t, "/aaa/bbb/", string(tmp))
}

func BenchmarkAppendPath(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		b.ReportAllocs()
		var tmp []byte
		for i := 0; i < b.N; i++ {
			// tmp = tmp[:0]
			tmp = gosl.AppendPath(tmp[:0], "/aaa", "bbb", "ccc", "/ddd")
		}
		//println(string(tmp))
	})
}

func TestAppendFill(t *testing.T) {
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

func TestAppendRepeat(t *testing.T) {
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

func BenchmarkAppendRepeat(b *testing.B) {
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

func TestAppendFit(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)
	gosl.Test(t, "hello th..", string(gosl.AppendFit(buf, "hello there", 10, '-', true)))
	gosl.Test(t, "hello ther", string(gosl.AppendFit(buf, "hello there", 10, '-', false)))
	gosl.Test(t, "", string(gosl.AppendFit(buf, "hello there", -1, '-', false)))
	gosl.Test(t, "hello-----", string(gosl.AppendFit(buf, "hello", 10, '-', false)))
}

func BenchmarkAppendFit(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			// buf = gosl.AppendFit(buf, msg, 20, ' ', false)
			buf = gosl.AppendFit(buf, "hello how are you?", 20, '-', true)
		}
		// println("["+buf.String()+"]")
	})
}

// TODO: Add AppendString()
func TestAppendString(t *testing.T) {
	var out gosl.Buf
	t.Run("standard-no-trim", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendString(out, "   tes    ", false)
		gosl.Test(t, "   tes    ", out.String())
	})
	t.Run("standard-trim", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendString(out, "tes", false)
		gosl.Test(t, "tes", out.String())
	})
	t.Run("middle", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringMiddle(out, "tes", 10, false)
		gosl.Test(t, "   tes    ", out.String())
	})
	t.Run("middle-0", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringMiddle(out, "tes", 0, false)
		gosl.Test(t, "", out.String())
	})
	t.Run("right", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringRight(out, "tes", 10, false)
		gosl.Test(t, "       tes", out.String())
	})
	t.Run("right-0", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringRight(out, "tes", 0, false)
		gosl.Test(t, "", out.String())
	})
}

// TODO: Add AppendString()
func BenchmarkAppendString(b *testing.B) {
	b.Run("standard", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendString(out, "  test  ",  false)
		}
		// println(out.String())
	})
	b.Run("standard-trim", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendString(out, "  test  ",  true)
		}
		//println(out.String())
	})
	b.Run("middle", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendStringMiddle(out, "test", 10, false)
		}
		// println(out.String())
	})
	b.Run("right", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendStringRight(out, "test", 10, false)
		}
		// println(out.String())
	})
}
