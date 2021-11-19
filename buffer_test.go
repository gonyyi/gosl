// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := gosl.NewBuffer(nil)
	gosl.Test(t, byte(0), buf.Last())
	buf.WriteString("hello")
	gosl.Test(t, byte('o'), buf.Last())
	buf.WriteString("gon")
	gosl.Test(t, byte('n'), buf.Last())
	buf.Reset()
	gosl.Test(t, byte(0), buf.Last())
}

func BenchmarkBuffer(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		s := []string{"abc", "def", "ghi"}
		b.ReportAllocs()
		buf := gosl.NewBuffer(nil)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(s[i%3])
			_ = buf.Last()
		}
	})
}

func Test_Buffer_Trim(t *testing.T) {
	buf := gosl.NewBuffer(nil)
	buf.WriteString("hello") // buf = "hello"
	buf.WriteString(" gon")  // buf = "hello gon"
	buf.Trim(0)
	gosl.Test(t, "hello gon", buf.String())
	buf.Trim(1)
	gosl.Test(t, "hello go", buf.String())
	buf.Trim(4)
	gosl.Test(t, "hell", buf.String())
	buf.Trim(4)
	gosl.Test(t, "", buf.String())
	buf.Trim(4)
	gosl.Test(t, "", buf.String())
}

func Benchmark_Buffer_Trim(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		s := []string{"abc", "def", "ghi"}
		b.ReportAllocs()
		buf := gosl.NewBuffer(nil)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			buf.WriteString(s[i%3])
			buf.Trim(1)
		}
		// println(buf.String())
	})
}



