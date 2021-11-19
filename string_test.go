// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/16/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_String_StringJoin(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		var in = []string{"gon", "is", "always", "awesome"}
		var out gosl.Buf
		out = gosl.Joins(out, in, ' ')
		gosl.Test(t, "gon is always awesome", out.String())
	})
}

func Benchmark_String_StringJoin(b *testing.B) {
	// Confirmed zero allocation
	var in = []string{"gon", "is", "always", "awesome"}

	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		var out gosl.Buf
		for i := 0; i < b.N; i++ {
			out = gosl.Joins(out[:0], in, ' ')
		}
		// println(out.String())
	})
}

func Test_String_StringSplit(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		var out []string
		var in = "  gon  is  always  awesome   "
		out = gosl.Splits(out, in, ' ')

		gosl.Test(t, 4, len(out))
		if len(out) == 4 {
			gosl.Test(t, "gon", out[0])
			gosl.Test(t, "is", out[1])
			gosl.Test(t, "always", out[2])
			gosl.Test(t, "awesome", out[3])
		} else {
			t.Errorf("unexpected result")
		}
	})
}

func Benchmark_String_StringSplit(b *testing.B) {
	// Confirmed zero allocation
	var out []string
	var in = "  gon  is  always  awesome   "
	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out = out[:0]
			out = gosl.Splits(out, in, ' ')
		}
		// gosl.Strings(out).Print()
	})
}

func Test_String_StringTrim(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := gosl.Trims("  1start hello  haha end  ", true, true)
		gosl.Test(t, "1start hello  haha end", out)

		out = gosl.Trims("  2start hello  haha end  ", false, true)
		gosl.Test(t, "  2start hello  haha end", out)

		out = gosl.Trims("  3start hello  haha end  ", true, false)
		gosl.Test(t, "3start hello  haha end  ", out)

		out = gosl.Trims("  4start hello  haha end  ", false, false)
		gosl.Test(t, "  4start hello  haha end  ", out)
	})
}

func Benchmark_String_StringTrim(b *testing.B) {
	p := func(s string) {
		// println("<<"+s+">>")
	}

	b.Run("Left+Right", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = out.WriteString(gosl.Trims("  1start hello  haha end  ", true, true))
		}
		p(out.String())
	})
	b.Run("Left", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = out.WriteString(gosl.Trims("  1start hello  haha end  ", true, false))
		}
		p(out.String())
	})
	b.Run("Right", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = out.WriteString(gosl.Trims("  1start hello  haha end  ", false, true))
		}
		p(out.String())
	})
	b.Run("None", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = out.WriteString(gosl.Trims("  1start hello  haha end  ", false, false))
		}
		p(out.String())
	})
}

func Test_String_HasPrefix(t *testing.T) {
	t.Run("gon vs go", func(t *testing.T) {
		var out bool
		out = gosl.HasPrefix("gon", "go")
		gosl.Test(t, true, out)
	})
	t.Run("gon vs gon", func(t *testing.T) {
		var out bool
		out = gosl.HasPrefix("gon", "gon")
		gosl.Test(t, true, out)
	})
	t.Run("gon vs on", func(t *testing.T) {
		var out bool
		out = gosl.HasPrefix("gon", "on")
		gosl.Test(t, false, out)
	})
	t.Run("gon vs null", func(t *testing.T) {
		var out bool
		out = gosl.HasPrefix("gon", "")
		gosl.Test(t, true, out)
	})
	t.Run("null vs on", func(t *testing.T) {
		var out bool
		out = gosl.HasPrefix("", "on")
		gosl.Test(t, false, out)
	})
	t.Run("null vs null", func(t *testing.T) {
		var out bool
		out = gosl.HasPrefix("", "")
		gosl.Test(t, true, out)
	})
}

func Benchmark_String_HasPrefix(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		a := "gon"
		pfx := "go"
		res := false
		for i := 0; i < b.N; i++ {
			res = gosl.HasPrefix(a, pfx)
		}
		_ = res
		// println(res)
	})
}

func Test_String_HasSuffix(t *testing.T) {
	t.Run("gon vs go", func(t *testing.T) {
		var out bool
		out = gosl.HasSuffix("go", "gon")
		gosl.Test(t, false, out)
	})
	t.Run("gon vs gon", func(t *testing.T) {
		var out bool
		out = gosl.HasSuffix("gon", "gon")
		gosl.Test(t, true, out)
	})
	t.Run("gon vs on", func(t *testing.T) {
		var out bool
		out = gosl.HasSuffix("gon", "on")
		gosl.Test(t, true, out)
	})
	t.Run("gon vs null", func(t *testing.T) {
		var out bool
		out = gosl.HasSuffix("gon", "")
		gosl.Test(t, true, out)
	})
	t.Run("null vs on", func(t *testing.T) {
		var out bool
		out = gosl.HasSuffix("", "on")
		gosl.Test(t, false, out)
	})
	t.Run("null vs null", func(t *testing.T) {
		var out bool
		out = gosl.HasSuffix("", "")
		gosl.Test(t, true, out)
	})
}

func Benchmark_String_HasSuffix(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		a := "gon"
		sfx := "on"
		res := false
		for i := 0; i < b.N; i++ {
			res = gosl.HasSuffix(a, sfx)
		}
		_ = res
		// println(res)
	})
}
func Test_String_TrimPrefix(t *testing.T) {
	gosl.Test(t, "n", gosl.TrimPrefix("gon", "go"))
	gosl.Test(t, "gon", gosl.TrimPrefix("gon", "goz"))
	gosl.Test(t, "gon", gosl.TrimPrefix("gon", ""))
	gosl.Test(t, "", gosl.TrimPrefix("", "abc"))
}

func Benchmark_String_TrimPrefix(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		a := "gon: gon is awesome"
		pfx := "gon: "
		buf := make(gosl.Buf, 0, 512)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = buf.WriteString(gosl.TrimPrefix(a, pfx))
		}
		// println(buf.String())
	})
}

func Test_String_TrimSuffix(t *testing.T) {
	gosl.Test(t, "gotta", gosl.TrimSuffix("gottago", "go"))
	gosl.Test(t, "gottago", gosl.TrimSuffix("gottago", "goz"))
	gosl.Test(t, "gottago", gosl.TrimSuffix("gottago", ""))
	gosl.Test(t, "", gosl.TrimSuffix("", "abc"))
}

func Benchmark_String_TrimSuffix(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		a := "gon: gon is awesome"
		sfx := "awesome"
		buf := make(gosl.Buf, 0, 512)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = buf.WriteString(gosl.TrimSuffix(a, sfx))
		}
		// println(buf.String())
	})
}

func Test_String_IsNumber(t *testing.T) {
	gosl.Test(t, false, gosl.IsNumber("gon"))
	gosl.Test(t, true, gosl.IsNumber("123"))
	gosl.Test(t, false, gosl.IsNumber("go123n"))
	gosl.Test(t, false, gosl.IsNumber(""))
}

func Benchmark_String_IsNumber(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		res := false
		for i := 0; i < b.N; i++ {
			res = gosl.IsNumber("434")
		}
		_ = res
		// println(res)
	})
}

