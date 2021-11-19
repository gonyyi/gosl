// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/19/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

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
		//println(string(buf))
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
		//println(string(buf))
	})
}
