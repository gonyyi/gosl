// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := gosl.NewBuffer(nil)
	gosl.TestByte(t, 0, buf.Last())
	buf.WriteString("hello")
	gosl.TestByte(t, 'o', buf.Last())
	buf.WriteString("gon")
	gosl.TestByte(t, 'n', buf.Last())
	buf.Reset()
	gosl.TestByte(t, 0, buf.Last())
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
	gosl.TestString(t, "hello gon", buf.String())
	buf.Trim(1)
	gosl.TestString(t, "hello go", buf.String())
	buf.Trim(4)
	gosl.TestString(t, "hell", buf.String())
	buf.Trim(4)
	gosl.TestString(t, "", buf.String())
	buf.Trim(4)
	gosl.TestString(t, "", buf.String())
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


