// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

// func TestBuffer(t *testing.T) {
// 	b1 := gosl.GetBuffer()
// 	println(cap(b1.Buf))
// 	gosl.GlobalBufferSize = 2048
// 	b2 := gosl.GetBuffer()
// 	println(cap(b2.Buf))
// 	b1.Free()
// 	b3 := gosl.GetBuffer()
// 	println(cap(b3.Buf))
// 	b4 := gosl.GetBuffer()
// 	println(cap(b4.Buf))
// }

func BenchmarkBuffer(b *testing.B) {
	sample := make(gosl.Buf, 0, 1024)

	b.Run("T1", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := gosl.GetBuffer()
			buf.Buf = buf.Buf.Set("test yo!").WriteInt(i)
			sample = sample.Reset().WriteBytes(buf.Buf...)
			buf.Free()
		}
		// sample.Println()
	})
}
