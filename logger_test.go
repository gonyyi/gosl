// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

func TestLogger(t *testing.T) {
	var l gosl.Logger

	buf := make(gosl.Buf, 0, 1024) // output goes here

	t.Run("Output", func(t *testing.T) {
		l = l.SetOutput(gosl.Discard, false)
		gosl.Test(t, true, l.Output() != nil)
		l = l.SetOutput(nil, false)
		gosl.Test(t, false, l.Output() != nil)
	})

	t.Run("Close", func(t *testing.T) {
		l = l.SetOutput(gosl.Discard, false)
		l.Close()
		l = l.SetOutput(nil, false)
		l.Close()
	})

	t.Run("Output", func(t *testing.T) {
		t.Run("Newline", func(t *testing.T) {
			l = l.SetOutput(&buf, true)
			buf = buf.Reset()
			l.Write([]byte("byte1"))
			l.WriteString("string1")
			gosl.Test(t, "byte1\nstring1\n", buf.String())
			// buf.Println()
		})
		t.Run("NoNewline", func(t *testing.T) {
			l = l.SetOutput(&buf, false)
			buf = buf.Reset()
			l.Write([]byte("byte1"))
			l.WriteString("string1")
			gosl.Test(t, "byte1string1", buf.String())
			// buf.Println()
		})
	})

	t.Run("NoOutput", func(t *testing.T) {
		t.Run("Newline", func(t *testing.T) {
			l = l.SetOutput(gosl.Discard, true)
			buf = buf.Reset()
			l.Write([]byte("byte1"))
			l.WriteString("string1")
			gosl.Test(t, "", buf.String())
			// buf.Println()
		})
		t.Run("NoNewline", func(t *testing.T) {
			l = l.SetOutput(gosl.Discard, false)
			buf = buf.Reset()
			l.Write([]byte("byte1"))
			l.WriteString("string1")
			gosl.Test(t, "", buf.String())
			// buf.Println()
		})
	})
}

func BenchmarkLogger(b *testing.B) {
	var l gosl.Logger

	strs := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	bytes := [][]byte{[]byte("one"), []byte("two"), []byte("three"), []byte("four"), []byte("five"), []byte("six"), []byte("seven"), []byte("eight"), []byte("nine"), []byte("ten")}

	buf := make(gosl.Buf, 0, 1024) // output goes here

	b.Run("Output", func(b *testing.B) {
		b.Run("Newline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard, true)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard, true)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
		b.Run("NoNewline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard, false)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard, false)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
	})

	b.Run("NoOutput", func(b *testing.B) {
		b.Run("Newline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil, true)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil, true)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
		b.Run("NoNewline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil, false)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil, false)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
	})

}
