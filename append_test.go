// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/30/2021

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

func TestAppendBool(t *testing.T) {
	buf := make(gosl.Buf, 0, 128)
	buf = gosl.AppendBool(buf, true)
	buf = gosl.AppendBool(buf, false)
	gosl.Test(t, "truefalse", buf.String())
}

func TestAppendFills(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendFills(a, nil, 10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("negative-n", func(t *testing.T) {
		a := []byte("gon")
		b := []byte("123")
		a = gosl.AppendFills(a, b, -10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("basic", func(t *testing.T) {
		a := []byte("gon")
		b := []byte{'-', '_'}
		a = gosl.AppendFills(a, b, 9)
		gosl.Test(t, "gon-_-_-_-_-", string(a))
	})
}

func TestAppendRepeat(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendRepeats(a, nil, 10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("negative-n", func(t *testing.T) {
		a := []byte("gon")
		b := []byte("123")
		a = gosl.AppendRepeats(a, b, -10)
		gosl.Test(t, "gon", string(a))
	})

	t.Run("basic", func(t *testing.T) {
		a := []byte("gon")
		b := []byte("123")
		a = gosl.AppendRepeats(a, b, 2)
		gosl.Test(t, "gon123123", string(a))
	})

	t.Run("AppendRepeat-basic", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendRepeat(a, ':', 2)
		gosl.Test(t, "gon::", string(a))
	})

	t.Run("AppendRepeat-nil", func(t *testing.T) {
		a := []byte("gon")
		a = gosl.AppendRepeat(a, ':', -1)
		gosl.Test(t, "gon", string(a))
	})

}

func BenchmarkAppendRepeat(b *testing.B) {
	b.Run("AppendRepeats", func(b *testing.B) {
		b.ReportAllocs()
		var buf = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = gosl.AppendRepeats(buf, []byte("abc"), 10)
		}
		// println(Buffer.String())
	})
	b.Run("AppendRepeat", func(b *testing.B) {
		b.ReportAllocs()
		var buf = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = gosl.AppendRepeat(buf, '-', 10)
		}
		// println(Buffer.String())
	})
}

func TestAppendString(t *testing.T) {
	var out gosl.Buf
	t.Run("standard-no-trim", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendString(out, "   tes    ", false)
		gosl.Test(t, "   tes    ", out.String())
	})
	t.Run("standard-trim", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendString(out, "   tes   ", true)
		gosl.Test(t, "tes", out.String())
	})
	t.Run("standard-trim-2", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendString(out, "   test    words  ", true)
		gosl.Test(t, "test words", out.String())
	})
}

func TestAppendStringFit(t *testing.T) {
	out := make(gosl.Buf, 0, 1024)

	t.Run("Fit", func(t *testing.T) {
		gosl.Test(t, "hello th..", string(gosl.AppendStringFit(out, "hello there", 10, '-', true)))
		gosl.Test(t, "hello ther", string(gosl.AppendStringFit(out, "hello there", 10, '-', false)))
		gosl.Test(t, "", string(gosl.AppendStringFit(out, "hello there", -1, '-', false)))
		gosl.Test(t, "hello-----", string(gosl.AppendStringFit(out, "hello", 10, '-', false)))
	})

	t.Run("FitCenter", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringFitCenter(out, "tes", 10, ' ', false)
		gosl.Test(t, "   tes    ", out.String())
	})
	t.Run("FitCenter-0", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringFitCenter(out, "tes", 0, ' ', false)
		gosl.Test(t, "", out.String())
	})
	t.Run("FitCenter-1", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringFitCenter(out, "tes", 1, ' ', false)
		gosl.Test(t, "t", out.String())
	})
	t.Run("FitRight", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringFitRight(out, "tes", 10, ' ', false)
		gosl.Test(t, "       tes", out.String())
	})
	t.Run("FitRight-0", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringFitRight(out, "tes", 0, ' ', false)
		gosl.Test(t, "", out.String())
	})
	t.Run("FitRight-1", func(t *testing.T) {
		out = out.Reset()
		out = gosl.AppendStringFitRight(out, "tes", 1, ' ', false)
		gosl.Test(t, "t", out.String())
	})
}

func BenchmarkAppendString(b *testing.B) {
	b.Run("standard", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendString(out, "  test  ", false)
		}
		// println(out.String())
	})

	b.Run("standard-trim", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendString(out, "  test  ", true)
		}
		//println(out.String())
	})
}

func BenchmarkAppendStringFit(b *testing.B) {
	b.Run("Fit", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendStringFit(out, "test", 10, ' ', false)
		}
		// println(out.String())
	})
	b.Run("FitCenter", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendStringFitCenter(out, "test", 10, ' ', false)
		}
		// println(out.String())
	})
	b.Run("FitRight", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		// t := "test "
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = gosl.AppendStringFitRight(out, "test", 10, ' ', false)
		}
		// println(out.String())
	})
}
