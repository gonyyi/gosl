// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/06/2021

package gosl_test

import (
	"bytes"
	"github.com/gonyyi/gosl"
	"testing"
)

func TestBytesEqual(t *testing.T) {
	b1 := []byte("hey this is crazyyiiii")
	b2 := []byte("hey this is crazyyiiii")

	gosl.Test(t, 0, bytes.Compare(b1, b2))
	gosl.Test(t, true, gosl.BytesEqual(b1, b2))

	b2[0] = 0
	gosl.Test(t, 1, bytes.Compare(b1, b2))
	gosl.Test(t, false, gosl.BytesEqual(b1, b2))

	b1[0] = 0
	gosl.Test(t, 0, bytes.Compare(b1, b2))
	gosl.Test(t, true, gosl.BytesEqual(b1, b2))
}

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
