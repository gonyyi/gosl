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

func TestBuffer(t *testing.T) {
	b1 := gosl.GetBuffer()
	t.Run("Bytes()", func(t *testing.T) {
		b1.Set("test1")
		tmp := b1.Bytes()
		gosl.Test(t, true,  gosl.BytesEqual(b1.Buf, tmp))
	})
	t.Run("Cap()", func(t *testing.T) {
		b1.Set("test2")
		gosl.Test(t, true, b1.Buf.Cap() == b1.Cap())
	})
	t.Run("Len()", func(t *testing.T) {
		b1.Set("test3")
		gosl.Test(t, 5, b1.Len())
	})
	t.Run("Println()", func(t *testing.T) {
		b1.Set("test4")
		// b1.Println()
	})
	t.Run("Reset()", func(t *testing.T) {
		b1.Set("test5")
		b1.Reset()
		gosl.Test(t, 0, b1.Len())
	})
	t.Run("Set()", func(t *testing.T) {
		b1.Reset()
		b1.Set("test5")
		b1.Set("test6")
		gosl.Test(t, "test6", b1.String())
	})
	t.Run("String()", func(t *testing.T) {
		b1.Reset()
		b1.Set("test7")
		gosl.Test(t, "test7", b1.String())
	})
	t.Run("Write()", func(t *testing.T) {
		b1.Set("te")
		n, err := b1.Write([]byte("st8"))
		gosl.Test(t, nil, err)
		gosl.Test(t, 3, n)
		gosl.Test(t, "test8", b1.String())
	})
	t.Run("WriteBytes()", func(t *testing.T) {
		b1.Set("test9")
		b1.WriteBytes('-')
		b1.WriteBytes('1','2','3')
		gosl.Test(t, "test9-123", b1.String())
	})
	t.Run("WriteBool()", func(t *testing.T) {
		b1.Set("test10")
		b1.WriteBool(true)
		b1.WriteBool(false)
		gosl.Test(t, "test10truefalse", b1.String())
	})
	t.Run("WriteFloat()", func(t *testing.T) {
		b1.Set("test11-")
		b1.WriteFloat(3.141592, 2)
		gosl.Test(t, "test11-3.14", b1.String())
	})
	t.Run("WriteInt()", func(t *testing.T) {
		b1.Set("test12-")
		b1.WriteInt(1212)
		gosl.Test(t, "test12-1212", b1.String())
	})
	t.Run("WriteString()", func(t *testing.T) {
		b1.Set("test13-")
		b1.WriteString("done")
		gosl.Test(t, "test13-done", b1.String())
	})
	t.Run("WriteStrings()", func(t *testing.T) {
		b1.Set("test14-")
		b1.WriteStrings([]string{"a", "b", "c"}, ',', ' ')
		gosl.Test(t, "test14-a, b, c", b1.String())
	})
	t.Run("WriteTo()", func(t *testing.T) {
		tmpBuf := make(gosl.Buf, 0, 1024)
		b1.Set("test15")
		b1.WriteTo(&tmpBuf)
		gosl.Test(t, "test15", string(tmpBuf))
	})
}
