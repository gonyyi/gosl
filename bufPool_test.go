// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := gosl.Buf(nil)
	gosl.Test(t, byte(0), buf.Last())
	buf = buf.WriteString("hello")
	gosl.Test(t, byte('o'), buf.Last())
	buf = buf.WriteString("gon")
	gosl.Test(t, byte('n'), buf.Last())
	buf = buf.Reset()
	gosl.Test(t, byte(0), buf.Last())
}

func BenchmarkBuffer(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		s := []string{"abc", "def", "ghi"}
		b.ReportAllocs()
		buf := gosl.Buf(nil)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = buf.WriteString(s[i%3])
			_ = buf.Last()
		}
	})
}

func Test_Buffer_Trim(t *testing.T) {
	buf := gosl.Buf(nil)
	buf = buf.WriteString("hello") // Buffer = "hello"
	buf = buf.WriteString(" gon")  // Buffer = "hello gon"
	buf = buf.Trim(0)
	gosl.Test(t, "hello gon", buf.String())
	buf = buf.Trim(1)
	gosl.Test(t, "hello go", buf.String())
	buf = buf.Trim(4)
	gosl.Test(t, "hell", buf.String())
	buf = buf.Trim(4)
	gosl.Test(t, "", buf.String())
	buf = buf.Trim(4)
	gosl.Test(t, "", buf.String())
}

func BenchmarkBuffer_Trim(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		s := []string{"abc", "def", "ghi"}
		b.ReportAllocs()
		buf := gosl.Buf(nil)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf = buf.WriteString(s[i%3])
			buf = buf.Trim(1)
		}
		// println(Buffer.String())
	})
}



