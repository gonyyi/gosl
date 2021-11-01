// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Conv_Itoa(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		gosl.TestString(t, "0", gosl.Itoa(0))
		gosl.TestString(t, "10", gosl.Itoa(10))
		gosl.TestString(t, "-10", gosl.Itoa(-10))
		gosl.TestString(t, "10000000", gosl.Itoa(10000000))
		gosl.TestString(t, "-10000000", gosl.Itoa(-10000000))
	})
}

func Benchmark_Conv_Itoa(b *testing.B) {
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

func Test_Conv_Itoaf(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		gosl.TestString(t, "0", gosl.Itoaf(0, false))
		gosl.TestString(t, "10", gosl.Itoaf(10, false))
		gosl.TestString(t, "-10", gosl.Itoaf(-10, false))
		gosl.TestString(t, "10000000", gosl.Itoaf(10000000, false))
		gosl.TestString(t, "-10000000", gosl.Itoaf(-10000000, false))
	})
	t.Run("Comma", func(t *testing.T) {
		gosl.TestString(t, "0", gosl.Itoaf(0, true))
		gosl.TestString(t, "10", gosl.Itoaf(10, true))
		gosl.TestString(t, "-10", gosl.Itoaf(-10, true))
		gosl.TestString(t, "10,000,000", gosl.Itoaf(10000000, true))
		gosl.TestString(t, "-10,000,000", gosl.Itoaf(-10000000, true))
	})
}

func Benchmark_Conv_Itoaf(b *testing.B) {
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

func Test_Conv_Ftoa(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		gosl.TestString(t, "1.00", gosl.Ftoa(1))

		gosl.TestString(t, "-1.00", gosl.Ftoa(-1))

		gosl.TestString(t, "10.00", gosl.Ftoa(10))
		gosl.TestString(t, "100.00", gosl.Ftoa(100))
		gosl.TestString(t, "1000.00", gosl.Ftoa(1000))
		gosl.TestString(t, "10000.00", gosl.Ftoa(10000))
		gosl.TestString(t, "100000.00", gosl.Ftoa(100000))
		gosl.TestString(t, "1000000.00", gosl.Ftoa(1000000))

		gosl.TestString(t, "-10.00", gosl.Ftoa(-10))
		gosl.TestString(t, "-100.00", gosl.Ftoa(-100))
		gosl.TestString(t, "-1000.00", gosl.Ftoa(-1000))
		gosl.TestString(t, "-10000.00", gosl.Ftoa(-10000))
		gosl.TestString(t, "-100000.00", gosl.Ftoa(-100000))
		gosl.TestString(t, "-1000000.00", gosl.Ftoa(-1000000))
	})

	t.Run("Plain", func(t *testing.T) {
		gosl.TestString(t, "0.12", gosl.Ftoa(0.123))
	})
}

func Test_Conv_Ftoaf(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		gosl.TestString(t, "1", gosl.Ftoaf(1, 0, false))
		gosl.TestString(t, "1.0", gosl.Ftoaf(1, 1, false))
		gosl.TestString(t, "1.00", gosl.Ftoaf(1, 2, false))
		gosl.TestString(t, "1.000", gosl.Ftoaf(1, 3, false))
		gosl.TestString(t, "1.0000", gosl.Ftoaf(1, 4, false))
		gosl.TestString(t, "1.0000", gosl.Ftoaf(1, 5, false))

		gosl.TestString(t, "-1", gosl.Ftoaf(-1, 0, false))
		gosl.TestString(t, "-1.0", gosl.Ftoaf(-1, 1, false))
		gosl.TestString(t, "-1.00", gosl.Ftoaf(-1, 2, false))
		gosl.TestString(t, "-1.000", gosl.Ftoaf(-1, 3, false))
		gosl.TestString(t, "-1.0000", gosl.Ftoaf(-1, 4, false))
		gosl.TestString(t, "-1.0000", gosl.Ftoaf(-1, 5, false))

		gosl.TestString(t, "10", gosl.Ftoaf(10, 0, false))
		gosl.TestString(t, "100.0", gosl.Ftoaf(100, 1, false))
		gosl.TestString(t, "1000.00", gosl.Ftoaf(1000, 2, false))
		gosl.TestString(t, "10000.000", gosl.Ftoaf(10000, 3, false))
		gosl.TestString(t, "100000.0000", gosl.Ftoaf(100000, 4, false))
		gosl.TestString(t, "1000000.0000", gosl.Ftoaf(1000000, 5, false))

		gosl.TestString(t, "-10", gosl.Ftoaf(-10, 0, false))
		gosl.TestString(t, "-100.0", gosl.Ftoaf(-100, 1, false))
		gosl.TestString(t, "-1000.00", gosl.Ftoaf(-1000, 2, false))
		gosl.TestString(t, "-10000.000", gosl.Ftoaf(-10000, 3, false))
		gosl.TestString(t, "-100000.0000", gosl.Ftoaf(-100000, 4, false))
		gosl.TestString(t, "-1000000.0000", gosl.Ftoaf(-1000000, 5, false))

		gosl.TestString(t, "1", gosl.Ftoaf(1, 0, true))
		gosl.TestString(t, "1.0", gosl.Ftoaf(1, 1, true))
		gosl.TestString(t, "1.00", gosl.Ftoaf(1, 2, true))
		gosl.TestString(t, "1.000", gosl.Ftoaf(1, 3, true))
		gosl.TestString(t, "1.0000", gosl.Ftoaf(1, 4, true))
		gosl.TestString(t, "1.0000", gosl.Ftoaf(1, 5, true))

		gosl.TestString(t, "-1", gosl.Ftoaf(-1, 0, true))
		gosl.TestString(t, "-1.0", gosl.Ftoaf(-1, 1, true))
		gosl.TestString(t, "-1.00", gosl.Ftoaf(-1, 2, true))
		gosl.TestString(t, "-1.000", gosl.Ftoaf(-1, 3, true))
		gosl.TestString(t, "-1.0000", gosl.Ftoaf(-1, 4, true))
		gosl.TestString(t, "-1.0000", gosl.Ftoaf(-1, 5, true))

		gosl.TestString(t, "10", gosl.Ftoaf(10, 0, true))
		gosl.TestString(t, "100.0", gosl.Ftoaf(100, 1, true))
		gosl.TestString(t, "1,000.00", gosl.Ftoaf(1000, 2, true))
		gosl.TestString(t, "10,000.000", gosl.Ftoaf(10000, 3, true))
		gosl.TestString(t, "100,000.0000", gosl.Ftoaf(100000, 4, true))
		gosl.TestString(t, "1,000,000.0000", gosl.Ftoaf(1000000, 5, true))

		gosl.TestString(t, "-10", gosl.Ftoaf(-10, 0, true))
		gosl.TestString(t, "-100.0", gosl.Ftoaf(-100, 1, true))
		gosl.TestString(t, "-1,000.00", gosl.Ftoaf(-1000, 2, true))
		gosl.TestString(t, "-10,000.000", gosl.Ftoaf(-10000, 3, true))
		gosl.TestString(t, "-100,000.0000", gosl.Ftoaf(-100000, 4, true))
		gosl.TestString(t, "-1,000,000.0000", gosl.Ftoaf(-1000000, 5, true))
	})

	t.Run("Plain", func(t *testing.T) {
		gosl.TestString(t, "0.1", gosl.Ftoaf(0.123, 1, false))
		gosl.TestString(t, "0.12", gosl.Ftoaf(0.123, 2, false))
		gosl.TestString(t, "0.123", gosl.Ftoaf(0.123, 3, false))
		gosl.TestString(t, "0.1230", gosl.Ftoaf(0.123, 4, false))
		gosl.TestString(t, "0.1230", gosl.Ftoaf(0.123, 5, false))
	})
}

func Benchmark_Conv_Ftoa(b *testing.B) {
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

func Benchmark_Conv_Ftoaf(b *testing.B) {
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

