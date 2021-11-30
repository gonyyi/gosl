// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/30/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

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
		gosl.Test(t, 1234567890123456789, gosl.MustAtoi("+1234567890123456789", -9999))
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
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Itoa(i))
		}
		// println(Buffer.String())
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
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Itoaf(i, false))
		}
		// println(Buffer.String())
	})

	// Confirmed zero allocation
	b.Run("Comma", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Itoaf(i, true))
		}
		// println(Buffer.String())
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
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoa(float64(i) + 0.1))
		}
		// println(Buffer.String())
	})

	// Confirmed zero allocation
	b.Run("Comma", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoa(float64(i) + 0.1))
		}
		// println(Buffer.String())
	})
}

func Benchmark_String_Ftoaf(b *testing.B) {
	// Confirmed zero allocation
	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoaf(float64(i)+0.1, 6, false))
		}
		// println(Buffer.String())
	})

	// Confirmed zero allocation
	b.Run("Comma", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(gosl.Ftoaf(float64(i)+0.1, 6, false))
		}
		// println(Buffer.String())
	})
}

func TestToLower(t *testing.T) {
	tmp := "abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere"
	gosl.Test(t, "abcdefghijklmnopqrstuvwxyz-abcdefghijklmnopqrstuvwxyz-0123456789-gonyiishere",
		gosl.ToLower(tmp))
}

func BenchmarkToLower(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	_ = buf
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		tmp := "abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere"
		tmp = "GonYiIsHere123"
		for i := 0; i < b.N; i++ {
			// Buffer = Buffer.Reset().WriteString( gosl.ToLower(tmp) )
			_ = gosl.ToLower(tmp)
		}
	})
}

func TestToUpper(t *testing.T) {
	tmp := "abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere"
	gosl.Test(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GONYIISHERE",
		gosl.ToUpper(tmp))
}

func BenchmarkToUpper(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		tmp := "abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere"
		tmp = "GonYiIsHere123"
		for i := 0; i < b.N; i++ {
			_ = gosl.ToUpper(tmp)
		}
	})
}
