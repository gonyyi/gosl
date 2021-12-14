// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/14/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestAppendSize(t *testing.T) {
	tmp := make(gosl.Buf, 0, 1024)
	tmp = gosl.AppendSize(tmp, 1221233, 2)
	gosl.Test(t, "1.16MB", tmp.String())

	tmp = gosl.AppendSize(tmp.Reset(), 123*gosl.KB, 2)
	gosl.Test(t, "123.00KB", tmp.String())

	tmp = gosl.AppendSize(tmp.Reset(), 123*gosl.MB, 2)
	gosl.Test(t, "123.00MB", tmp.String())

	tmp = gosl.AppendSize(tmp.Reset(), 123*gosl.GB, 2)
	gosl.Test(t, "123.00GB", tmp.String())

	tmp = gosl.AppendSize(tmp.Reset(), 123*gosl.TB+345*gosl.GB, 2)
	gosl.Test(t, "123.33TB", tmp.String())

	tmp = gosl.AppendSize(tmp.Reset(), 1*gosl.PB+330*gosl.TB, 2)
	gosl.Test(t, "1.32PB", tmp.String())
}

func TestAppendSizeIn(t *testing.T) {
	tmp := make(gosl.Buf, 0, 1024)

	tmp = gosl.AppendSizeIn(tmp.Reset(), 1*gosl.TB+345*gosl.GB, gosl.GB, 2, true)
	gosl.Test(t, "1,369.00GB", tmp.String())

	tmp = gosl.AppendSizeIn(tmp.Reset(), 1*gosl.TB+345*gosl.GB, gosl.GB, 2, false)
	gosl.Test(t, "1369.00GB", tmp.String())
}

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
		// println(string(tmp))
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
		// println(out.String())
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

func TestAppendStringMask(t *testing.T) {
	a := "gon12345"
	gosl.Test(t, "go***345", string(gosl.AppendStringMask(nil, a, 2, 3)))
	gosl.Test(t, "gon***45", string(gosl.AppendStringMask(nil, a, 3, 2)))
	gosl.Test(t, "gon*****", string(gosl.AppendStringMask(nil, a, 3, 0)))
	gosl.Test(t, "*****345", string(gosl.AppendStringMask(nil, a, 0, 3)))
	gosl.Test(t, "gon12345", string(gosl.AppendStringMask(nil, a, 2, 10)))
	gosl.Test(t, "********", string(gosl.AppendStringMask(nil, a, -10, -10)))
}

func BenchmarkAppendStringMask(b *testing.B) {
	a := "gon12345"
	buf := make(gosl.Buf, 0, 1024)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf = gosl.AppendStringMask(buf.Reset(), a, 2, i%5)
	}
	// buf.Println()
}

func TestAppendLinePrefix(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)
	src := []byte("\nthis is something\nhahaha\n")
	buf = gosl.AppendLinePrefix(buf.Reset(), src, "  > ")
	gosl.Test(t, "  > \n  > this is something\n  > hahaha\n  > ", buf.String())
	// buf.Println()

	src = []byte("abc\n123\ndef\n\nds")
	buf = gosl.AppendLinePrefix(buf.Reset(), src, "  > ")
	gosl.Test(t, "  > abc\n  > 123\n  > def\n  > \n  > ds", buf.String())
	// buf.Println()
}

func BenchmarkAppendLinePrefix(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		src := []byte("this is something\nhahaha")
		for i := 0; i < b.N; i++ {
			buf = gosl.AppendLinePrefix(buf.Reset(), src, " -> ")
		}
		// buf.Println()
	})
}

func TestAppendStringLinePrefix(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)
	buf = gosl.AppendStringLinePrefix(buf.Reset(), "\nthis is something\nhahaha\n", "  > ")
	gosl.Test(t, "  > \n  > this is something\n  > hahaha\n  > ", buf.String())
	// buf.Println()

	buf = gosl.AppendStringLinePrefix(buf.Reset(), "abc\n123\ndef\n\nds", "  > ")
	gosl.Test(t, "  > abc\n  > 123\n  > def\n  > \n  > ds", buf.String())
	// buf.Println()
}

func BenchmarkAppendStringLinePrefix(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf = gosl.AppendStringLinePrefix(buf.Reset(), "this is something\nhahaha", " -> ")
		}
		// buf.Println()
	})
}
