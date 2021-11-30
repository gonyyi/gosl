// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/30/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestNewBytesFilter(t *testing.T) {
	buf := make(gosl.Buf, 0, 512)
	b := []byte("0123456789")

	f := gosl.NewBytesFilter(true, b) // only allow numbers
	buf = f([]byte("abc123"))

	gosl.Test(t, "123", buf.String())

	f = gosl.NewBytesFilter(false, b) // disallow numbers
	buf = f([]byte("abc123"))

	gosl.Test(t, "abc", buf.String())
}

func TestBytesInsert(t *testing.T) {
	buf := make(gosl.Buf, 0, 512)
	buf = buf.WriteString("GON123")
	buf = gosl.BytesInsert(buf, 3, []byte(" says "))
	gosl.Test(t, "GON says 123", buf.String())

	buf = buf.Reset().WriteString("GON123")
	buf = gosl.BytesInsert(buf, 6, []byte(" here"))
	gosl.Test(t, "GON123 here", buf.String())

	buf = buf.Reset().WriteString("GON123")
	buf = gosl.BytesInsert(buf, 7, []byte("here"))
	gosl.Test(t, "GON123here", buf.String())
}

func TestBytesToUpper(t *testing.T) {
	tmp := []byte("abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere")
	gosl.Test(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GONYIISHERE",
		string(gosl.BytesToUpper(tmp)))
}

func BenchmarkBytesToUpper(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		buf := []byte("abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere")
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesToUpper(buf)
		}
		// println(string(Buffer))
	})
}

func TestBytesToLower(t *testing.T) {
	tmp := []byte("abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere")
	gosl.Test(t, "abcdefghijklmnopqrstuvwxyz-abcdefghijklmnopqrstuvwxyz-0123456789-gonyiishere",
		string(gosl.BytesToLower(tmp)))
}

func BenchmarkBytesToLower(b *testing.B) {
	b.Run("t1", func(b *testing.B) {
		b.ReportAllocs()
		buf := []byte("abcdefghijklmnopqrstuvwxyz-ABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789-GonYiIsHere")
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesToLower(buf)
		}
		// println(string(Buffer))
	})
}
