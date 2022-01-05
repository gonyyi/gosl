// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/05/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

func TestLogger(t *testing.T) {
	var l gosl.Logger
	l = l.SetNewline(true)

	buf := make(gosl.Buf, 0, 1024) // output goes here

	t.Run("Enable-1", func(t *testing.T) {
		buf = buf.Reset()
		l2 := gosl.NewLogger(&buf)
		l2.WriteString("abc")
		l2.WriteString("123")

		l2 = l2.Enable(false)
		l2.WriteString("456") // will not be printed
		l2.WriteString("789") // will not be printed

		l2 = l2.Enable(true)
		l2.WriteString("bcd")
		l2.WriteString("cde")

		gosl.Test(t, "abc\n123\nbcd\ncde\n", buf.String())
	})

	t.Run("Enable-2", func(t *testing.T) {
		// Logger was copied using `Enable()`, `SetOutput`, etc, it should work
		// independent to its original

		buf = buf.Reset()
		l2 := gosl.NewLogger(&buf)
		l2 = l2.Enable(false)

		l2.WriteString("l2-1") // does not print

		f1 := func() {
			lf1 := l2.Enable(true)
			lf1.WriteString("f1-1") // prints
		}

		f2 := func() {
			l2.WriteString("f2-1") // does not print
		}

		f1()
		f2()
		gosl.Test(t, "f1-1\n", buf.String())
	})

	t.Run("SetNewline", func(t *testing.T) {
		buf = buf.Reset()
		l2 := gosl.NewLogger(&buf)
		l2.WriteString("l2-1")
		l2.WriteString("l2-2")

		l2 = l2.SetNewline(false)
		l2.WriteString("l2-3")
		l2.WriteString("l2-4")

		gosl.Test(t, "l2-1\nl2-2\nl2-3l2-4", buf.String())
	})

	t.Run("Output", func(t *testing.T) {
		l = l.SetOutput(gosl.Discard)
		gosl.Test(t, true, l.Output() != nil)
		l = l.SetOutput(nil)
		gosl.Test(t, false, l.Output() != nil)
	})

	t.Run("Close", func(t *testing.T) {
		l = l.SetOutput(gosl.Discard)
		l.Close()
		l = l.SetOutput(nil)
		l.Close()
	})

	t.Run("Output=Buf", func(t *testing.T) {
		t.Run("newline", func(t *testing.T) {
			l = gosl.NewLogger(&buf)
			buf = buf.Reset()
			l.Write([]byte("byte1"))
			l.WriteString("string1")
			gosl.Test(t, "byte1\nstring1\n", buf.String())
			// buf.Println()
		})
		t.Run("NoNewline", func(t *testing.T) {
			buf = buf.Reset()
			l = gosl.NewLogger(&buf)
			l = l.SetNewline(false)
			l.Write([]byte("byte1"))
			l.WriteString("string1")
			gosl.Test(t, "byte1string1", buf.String())
			// buf.Println()
		})
	})

	t.Run("Output=Discard", func(t *testing.T) {
		t.Run("newline", func(t *testing.T) {
			buf = buf.Reset()
			l = gosl.NewLogger(gosl.Discard)
			l.Write([]byte("byte1"))
			l.WriteString("string1")
			gosl.Test(t, "", buf.String())
			// buf.Println()
		})
		t.Run("NoNewline", func(t *testing.T) {
			buf = buf.Reset()
			l = gosl.NewLogger(gosl.Discard)
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
		b.Run("newline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
		b.Run("NoNewline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(gosl.Discard)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
	})

	b.Run("NoOutput", func(b *testing.B) {
		b.Run("newline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
		b.Run("NoNewline", func(b *testing.B) {
			b.Run("bytes", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.Write(bytes[i%10])
				}
			})
			b.Run("string", func(b *testing.B) {
				b.ReportAllocs()
				l = l.SetOutput(nil)
				for i := 0; i < b.N; i++ {
					buf = buf.Reset()
					l.WriteString(strs[i%10])
				}
			})
		})
	})

}
