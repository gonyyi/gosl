// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Buf(t *testing.T) {
	buf := gosl.Buf{}
	gosl.Test(t, byte(0), buf.Last())

	buf = buf.WriteString("hello")
	gosl.Test(t,  byte('o'), buf.Last())

	buf = buf.WriteString("gon")
	gosl.Test(t, byte('n'), buf.Last())

	buf = buf.Reset()
	gosl.Test(t, byte(0), buf.Last())

	buf = buf.WriteString("test gon")
	gosl.Test(t, "test gon", buf.String())

	buf = buf.Reset()
	buf = buf.WriteString("name:").WriteString("gon").WriteByte(',')
	buf = buf.WriteString("weight:").WriteInt(190).WriteByte(',')
	buf = buf.WriteString("gpa:").WriteFloat64(1.1).WriteByte(',')
	buf = buf.WriteString("isGoodStudent:").WriteBool(false)
	// println(buf.String())

}

func Benchmark_Buf(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		s := []string{"abc", "def", "ghi"}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := make(gosl.Buf, 512)
			buf = buf.Reset()
			buf = buf.WriteString(s[i%3])
			_ = buf.Last()
		}
		// println(string(out))
	})

	b.Run("all", func(b *testing.B) {
		b.ReportAllocs()
		var out []byte
		for i := 0; i < b.N; i++ {
			buf := make(gosl.Buf, 0, 512)
			buf = buf.WriteString("id:").WriteInt(i).WriteByte(',')
			buf = buf.WriteString("name:").WriteString("gon").WriteByte(',')
			buf = buf.WriteString("weight:").WriteInt(190).WriteByte(',')
			buf = buf.WriteString("gpa:").WriteFloat64(1.1).WriteByte(',')
			buf = buf.WriteString("isGoodStudent:").WriteBool(false)
			out = append(out[:0], buf...)
		}
		// println(string(out))
	})
}


