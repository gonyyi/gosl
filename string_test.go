// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/16/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_String_StringJoin(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		var in = []string{"gon", "is", "always", "awesome"}
		var out gosl.Buf
		out = gosl.StringsJoin(out, in, ' ')
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
			out = gosl.StringsJoin(out[:0], in, ' ')
		}
		// println(out.String())
	})
}

func Test_String_StringSplit(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		var out []string
		var in = "  gon  is  always  awesome   "
		out = gosl.StringSplit(out, in, ' ')

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
			out = gosl.StringSplit(out, in, ' ')
		}
		// gosl.Strings(out).Print()
	})
}
func Test_String_Atoi(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		gosl.Test(t, -1234567890123456789, gosl.MustAtoi("-1234567890123456789", -9999))
		gosl.Test(t, -123456789, gosl.MustAtoi("-123,456,789", -9999))
		gosl.Test(t, -123, gosl.MustAtoi("-123", -9999))
		gosl.Test(t, 0, gosl.MustAtoi("-0", -9999))
		gosl.Test(t, 0, gosl.MustAtoi("0", -9999))
		gosl.Test(t, 123, gosl.MustAtoi("123", -9999))
		gosl.Test(t, 123, gosl.MustAtoi("0000000000123", -9999))
		gosl.Test(t, 123456789, gosl.MustAtoi("123,456,789", -9999))
		gosl.Test(t, 1234567890123456789, gosl.MustAtoi("1234567890123456789", -9999))
	})
}

func Benchmark_String_Atoi(b *testing.B) {
	// basic:        12.72 ns/op
	// strconv.Itoa:  7.73 ns/op
	b.Run("basic", func(b *testing.B) {
		b.ReportAllocs()
		m := 0
		_ = m
		for i := 0; i < b.N; i++ {
			_ = gosl.MustAtoi("253425234", 0)
			// _,_ = gosl.Atoi("253425234")
		}

		// println(m)
	})
}
func Test_String_Itoa(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		gosl.Test(t, "0", gosl.Itoa(0))
		gosl.Test(t, "10", gosl.Itoa(10))
		gosl.Test(t, "-10", gosl.Itoa(-10))
		gosl.Test(t, "10000000", gosl.Itoa(10000000))
		gosl.Test(t, "-10000000", gosl.Itoa(-10000000))
	})
}

func Benchmark_String_Itoa(b *testing.B) {
	// Confirmed zero allocation
	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.NewBuffer(make([]byte, 0, 1024))
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Itoa(i))
		}
		// println(Buf.String())
	})
}

func Test_String_Itoaf(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		gosl.Test(t, "0", gosl.Itoaf(0, false))
		gosl.Test(t, "10", gosl.Itoaf(10, false))
		gosl.Test(t, "-10", gosl.Itoaf(-10, false))
		gosl.Test(t, "10000000", gosl.Itoaf(10000000, false))
		gosl.Test(t, "-10000000", gosl.Itoaf(-10000000, false))
	})
	t.Run("Comma", func(t *testing.T) {
		gosl.Test(t, "0", gosl.Itoaf(0, true))
		gosl.Test(t, "10", gosl.Itoaf(10, true))
		gosl.Test(t, "-10", gosl.Itoaf(-10, true))
		gosl.Test(t, "10,000,000", gosl.Itoaf(10000000, true))
		gosl.Test(t, "-10,000,000", gosl.Itoaf(-10000000, true))
	})
}

func Benchmark_String_Itoaf(b *testing.B) {
	// Confirmed zero allocation
	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.NewBuffer(make([]byte, 0, 1024))
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Itoaf(i, false))
		}
		// println(Buf.String())
	})

	// Confirmed zero allocation
	b.Run("Comma", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.NewBuffer(make([]byte, 0, 1024))
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Itoaf(i, true))
		}
		// println(Buf.String())
	})
}

func Test_String_StringTrim(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		out := gosl.StringTrim("  1start hello  haha end  ", true, true)
		gosl.Test(t, "1start hello  haha end", out)

		out = gosl.StringTrim("  2start hello  haha end  ", false, true)
		gosl.Test(t, "  2start hello  haha end", out)

		out = gosl.StringTrim("  3start hello  haha end  ", true, false)
		gosl.Test(t, "3start hello  haha end  ", out)

		out = gosl.StringTrim("  4start hello  haha end  ", false, false)
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
			out = out.WriteString(gosl.StringTrim("  1start hello  haha end  ", true, true))
		}
		p(out.String())
	})
	b.Run("Left", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = out.WriteString(gosl.StringTrim("  1start hello  haha end  ", true, false))
		}
		p(out.String())
	})
	b.Run("Right", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = out.WriteString(gosl.StringTrim("  1start hello  haha end  ", false, true))
		}
		p(out.String())
	})
	b.Run("None", func(b *testing.B) {
		b.ReportAllocs()
		var out = make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			out = out.Reset()
			out = out.WriteString(gosl.StringTrim("  1start hello  haha end  ", false, false))
		}
		p(out.String())
	})
}
func Test_String_Ftoa(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		gosl.Test(t, "1.00", gosl.Ftoa(1))

		gosl.Test(t, "-1.00", gosl.Ftoa(-1))

		gosl.Test(t, "10.00", gosl.Ftoa(10))
		gosl.Test(t, "100.00", gosl.Ftoa(100))
		gosl.Test(t, "1000.00", gosl.Ftoa(1000))
		gosl.Test(t, "10000.00", gosl.Ftoa(10000))
		gosl.Test(t, "100000.00", gosl.Ftoa(100000))
		gosl.Test(t, "1000000.00", gosl.Ftoa(1000000))

		gosl.Test(t, "-10.00", gosl.Ftoa(-10))
		gosl.Test(t, "-100.00", gosl.Ftoa(-100))
		gosl.Test(t, "-1000.00", gosl.Ftoa(-1000))
		gosl.Test(t, "-10000.00", gosl.Ftoa(-10000))
		gosl.Test(t, "-100000.00", gosl.Ftoa(-100000))
		gosl.Test(t, "-1000000.00", gosl.Ftoa(-1000000))
	})

	t.Run("Plain", func(t *testing.T) {
		gosl.Test(t, "0.12", gosl.Ftoa(0.123))
	})
}

func Test_String_Ftoaf(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		gosl.Test(t, "1", gosl.Ftoaf(1, 0, false))
		gosl.Test(t, "1.0", gosl.Ftoaf(1, 1, false))
		gosl.Test(t, "1.00", gosl.Ftoaf(1, 2, false))
		gosl.Test(t, "1.000", gosl.Ftoaf(1, 3, false))
		gosl.Test(t, "1.0000", gosl.Ftoaf(1, 4, false))
		gosl.Test(t, "1.0000", gosl.Ftoaf(1, 5, false))

		gosl.Test(t, "-1", gosl.Ftoaf(-1, 0, false))
		gosl.Test(t, "-1.0", gosl.Ftoaf(-1, 1, false))
		gosl.Test(t, "-1.00", gosl.Ftoaf(-1, 2, false))
		gosl.Test(t, "-1.000", gosl.Ftoaf(-1, 3, false))
		gosl.Test(t, "-1.0000", gosl.Ftoaf(-1, 4, false))
		gosl.Test(t, "-1.0000", gosl.Ftoaf(-1, 5, false))

		gosl.Test(t, "10", gosl.Ftoaf(10, 0, false))
		gosl.Test(t, "100.0", gosl.Ftoaf(100, 1, false))
		gosl.Test(t, "1000.00", gosl.Ftoaf(1000, 2, false))
		gosl.Test(t, "10000.000", gosl.Ftoaf(10000, 3, false))
		gosl.Test(t, "100000.0000", gosl.Ftoaf(100000, 4, false))
		gosl.Test(t, "1000000.0000", gosl.Ftoaf(1000000, 5, false))

		gosl.Test(t, "-10", gosl.Ftoaf(-10, 0, false))
		gosl.Test(t, "-100.0", gosl.Ftoaf(-100, 1, false))
		gosl.Test(t, "-1000.00", gosl.Ftoaf(-1000, 2, false))
		gosl.Test(t, "-10000.000", gosl.Ftoaf(-10000, 3, false))
		gosl.Test(t, "-100000.0000", gosl.Ftoaf(-100000, 4, false))
		gosl.Test(t, "-1000000.0000", gosl.Ftoaf(-1000000, 5, false))

		gosl.Test(t, "1", gosl.Ftoaf(1, 0, true))
		gosl.Test(t, "1.0", gosl.Ftoaf(1, 1, true))
		gosl.Test(t, "1.00", gosl.Ftoaf(1, 2, true))
		gosl.Test(t, "1.000", gosl.Ftoaf(1, 3, true))
		gosl.Test(t, "1.0000", gosl.Ftoaf(1, 4, true))
		gosl.Test(t, "1.0000", gosl.Ftoaf(1, 5, true))

		gosl.Test(t, "-1", gosl.Ftoaf(-1, 0, true))
		gosl.Test(t, "-1.0", gosl.Ftoaf(-1, 1, true))
		gosl.Test(t, "-1.00", gosl.Ftoaf(-1, 2, true))
		gosl.Test(t, "-1.000", gosl.Ftoaf(-1, 3, true))
		gosl.Test(t, "-1.0000", gosl.Ftoaf(-1, 4, true))
		gosl.Test(t, "-1.0000", gosl.Ftoaf(-1, 5, true))

		gosl.Test(t, "10", gosl.Ftoaf(10, 0, true))
		gosl.Test(t, "100.0", gosl.Ftoaf(100, 1, true))
		gosl.Test(t, "1,000.00", gosl.Ftoaf(1000, 2, true))
		gosl.Test(t, "10,000.000", gosl.Ftoaf(10000, 3, true))
		gosl.Test(t, "100,000.0000", gosl.Ftoaf(100000, 4, true))
		gosl.Test(t, "1,000,000.0000", gosl.Ftoaf(1000000, 5, true))

		gosl.Test(t, "-10", gosl.Ftoaf(-10, 0, true))
		gosl.Test(t, "-100.0", gosl.Ftoaf(-100, 1, true))
		gosl.Test(t, "-1,000.00", gosl.Ftoaf(-1000, 2, true))
		gosl.Test(t, "-10,000.000", gosl.Ftoaf(-10000, 3, true))
		gosl.Test(t, "-100,000.0000", gosl.Ftoaf(-100000, 4, true))
		gosl.Test(t, "-1,000,000.0000", gosl.Ftoaf(-1000000, 5, true))
	})

	t.Run("Plain", func(t *testing.T) {
		gosl.Test(t, "0.1", gosl.Ftoaf(0.123, 1, false))
		gosl.Test(t, "0.12", gosl.Ftoaf(0.123, 2, false))
		gosl.Test(t, "0.123", gosl.Ftoaf(0.123, 3, false))
		gosl.Test(t, "0.1230", gosl.Ftoaf(0.123, 4, false))
		gosl.Test(t, "0.1230", gosl.Ftoaf(0.123, 5, false))
	})
}

func Benchmark_String_Ftoa(b *testing.B) {
	// Confirmed zero allocation
	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.NewBuffer(make([]byte, 0, 1024))
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoa(float64(i) + 0.1))
		}
		// println(Buf.String())
	})

	// Confirmed zero allocation
	b.Run("Comma", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.NewBuffer(make([]byte, 0, 1024))
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoa(float64(i) + 0.1))
		}
		// println(Buf.String())
	})
}

func Benchmark_String_Ftoaf(b *testing.B) {
	// Confirmed zero allocation
	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.NewBuffer(make([]byte, 0, 1024))
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoaf(float64(i)+0.1, 6, false))
		}
		// println(Buf.String())
	})

	// Confirmed zero allocation
	b.Run("Comma", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.NewBuffer(make([]byte, 0, 1024))
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoaf(float64(i)+0.1, 6, false))
		}
		// println(Buf.String())
	})
}

func Test_String_StringsFollow(t *testing.T) {
	var a = []string{"gon", "is", "here"}
	gosl.StringsFollow(a, func(s string) string {
		return " > " + s
	})
	gosl.Test(t, "[ > gon  > is  > here]", fmt.Sprint(a))
}

func Test_String_IntsFollow(t *testing.T) {
	var a = []int{1, 3, 5, 8}
	gosl.IntsFollow(a, func(i int) int {
		return i + 10
	})
	gosl.Test(t, "[11 13 15 18]", fmt.Sprint(a))
}

func Test_String_IntsJoin(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var a = []int{1, 3, 5, 8, 100, 0, -200, 23}
		buf := make(gosl.Buf, 0, 4<<10)
		buf = buf.WriteByte('[')
		buf = gosl.IntsJoin(buf, a, ',')
		buf = buf.WriteByte(']')
		gosl.Test(t, "[1,3,5,8,100,0,-200,23]", buf.String())
	})

	t.Run("null", func(t *testing.T) {
		var a = []int{}
		buf := make(gosl.Buf, 0, 4<<10)
		buf = buf.WriteByte('[')
		buf = gosl.IntsJoin(buf, a, ',')
		buf = buf.WriteByte(']')
		gosl.Test(t, "[]", buf.String())
	})
}

func Test_String_AppendStringMiddle_Right(t *testing.T) {
	var out gosl.Buf
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

func Benchmark_String_StringMiddle_Right(b *testing.B) {
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

func Test_String_AppendFit(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)
	gosl.Test(t, "hello th..", string(gosl.AppendFit(buf, "hello there", 10, '-', true)))
	gosl.Test(t, "hello ther", string(gosl.AppendFit(buf, "hello there", 10, '-', false)))
	gosl.Test(t, "", string(gosl.AppendFit(buf, "hello there", -1, '-', false)))
	gosl.Test(t, "hello-----", string(gosl.AppendFit(buf, "hello", 10, '-', false)))
}

func Benchmark_String_AppendFit(b *testing.B) {
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
