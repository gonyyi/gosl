// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

func Test_String_Atoi(t *testing.T) {
	t.Run("too-large", func(t *testing.T) {
		// int32  : -2147483648 to 2147483647
		// int64  : -9223372036854775808 to 9223372036854775807
		p := func(t *testing.T, s string, expVal int, expOk bool) {
			n, ok := gosl.Atoi(s)
			if ok != expOk {
				t.Errorf("(E-1) Num: %s", s)
				t.Fail()
				return
			}
			if expVal != n {
				t.Errorf("(E-2) Num: %s, Exp: %d, Act: %d", s, expVal, n)
				t.Fail()
			}
		}

		if gosl.IntType == 64 {
			p(t, "9223372036854775806", 9223372036854775806, true)
			p(t, "9223372036854775807", 9223372036854775807, true) // largest positive
			p(t, "9223372036854775808", 0, false)
			p(t, "-9223372036854775807", -9223372036854775807, true) // lowest this code can handle
			p(t, "-9223372036854775808", 0, false)                   // actual lowest
			p(t, "-9223372036854775809", 0, false)

		} else { // 32 bit
			p(t, "-2147483648", 0, false)          // should too large
			p(t, "-2147483647", -2147483647, true) // ok
			p(t, "2147483647", 2147483647, true)   // ok
			p(t, "2147483648", 0, false)           // too large
		}
	})

	t.Run("Plain", func(t *testing.T) {
		gosl.Test(t, -1234567890123456789, gosl.MustAtoi("-1234567890123456789", -9999))
		gosl.Test(t, -123456789, gosl.MustAtoi("-123,456,789", -9999))
		gosl.Test(t, -123, gosl.MustAtoi("-123", -9999))
		gosl.Test(t, 0, gosl.MustAtoi("-0", -1111))
		gosl.Test(t, 0, gosl.MustAtoi("+0", -4444))
		gosl.Test(t, 0, gosl.MustAtoi("0", -9999))
		gosl.Test(t, -2222, gosl.MustAtoi("+", -2222))
		gosl.Test(t, -3333, gosl.MustAtoi("-", -3333))
		gosl.Test(t, 0, gosl.MustAtoi("0000000000000", -9999))
		gosl.Test(t, 123, gosl.MustAtoi("123", -9999))
		gosl.Test(t, 123, gosl.MustAtoi("0000000000123", -9999))
		gosl.Test(t, 123456789, gosl.MustAtoi("123,456,789", -9999))
		gosl.Test(t, 1234567890123456789, gosl.MustAtoi("1234567890123456789", -1111))
		gosl.Test(t, 1234567890123456789, gosl.MustAtoi("+1234567890123456789", -9999))
	})
}

func Benchmark_String_Atoi(b *testing.B) {
	// basic:        12.72 ns/op
	// strconv.Itoa:  7.73 ns/op
	b.Run("basic", func(b *testing.B) {
		b.ReportAllocs()
		m := 0
		_ = m
		for i := 0; i < b.N; i++ {
			_ = gosl.MustAtoi("253425234", 0)
			// _,_ = gosl.Atoi("253425234")
		}

		// println(m)
	})
}
func Test_String_Itoa(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		gosl.Test(t, "0", gosl.Itoa(0))
		gosl.Test(t, "10", gosl.Itoa(10))
		gosl.Test(t, "-10", gosl.Itoa(-10))
		gosl.Test(t, "10000000", gosl.Itoa(10000000))
		gosl.Test(t, "-10000000", gosl.Itoa(-10000000))
	})
}

func Benchmark_String_Itoa(b *testing.B) {
	// Confirmed zero allocation
	b.Run("Plain", func(b *testing.B) {
		b.ReportAllocs()
		buf := make(gosl.Buf, 0, 1024)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset().WriteString(gosl.Itoa(i))
			// _ = gosl.Itoa(i)
		}
		// println(buf.String())
	})
}

func TestString(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)
	t.Run("IsNumber", func(t *testing.T) {
		buf = buf.Reset()
		gosl.Test(t, true, gosl.IsNumber("123"))
		gosl.Test(t, true, gosl.IsNumber("-123"))
		gosl.Test(t, true, gosl.IsNumber("+123"))
		gosl.Test(t, true, gosl.IsNumber("+00000123"))
		gosl.Test(t, true, gosl.IsNumber("-00000123"))
		gosl.Test(t, true, gosl.IsNumber("00000123"))
		gosl.Test(t, false, gosl.IsNumber("123.123"))
		gosl.Test(t, false, gosl.IsNumber("-123.123"))
	})

	t.Run("Split", func(t *testing.T) {
		var s []string
		t.Run("emptyPrefixSuffix", func(t *testing.T) {
			buf = buf.Reset()
			s = gosl.Split(s[:0], "/abc/def/", '/')
			buf = buf.WriteStrings(s, ',')
			gosl.Test(t, ",abc,def,", buf)
		})
		t.Run("emptyPrefix", func(t *testing.T) {
			buf = buf.Reset()
			s = gosl.Split(s[:0], "/abc/def/ghi", '/')
			buf = buf.WriteStrings(s, ',')
			gosl.Test(t, ",abc,def,ghi", buf)
		})
		t.Run("emptySuffix", func(t *testing.T) {
			buf = buf.Reset()
			s = gosl.Split(s[:0], "abc/def/ghi/", '/')
			buf = buf.WriteStrings(s, ',')
			gosl.Test(t, "abc,def,ghi,", buf)
		})
		t.Run("delimNotFound", func(t *testing.T) {
			buf = buf.Reset()
			s = gosl.Split(s[:0], "/abc/def/ghi", 0)
			buf = buf.WriteStrings(s, ',')
			gosl.Test(t, "/abc/def/ghi", buf)
		})
		t.Run("emptyString", func(t *testing.T) {
			buf = buf.Reset()
			s = gosl.Split(s[:0], "", '/')
			buf = buf.WriteStrings(s, ',')
			gosl.Test(t, "", buf)
		})
	})

	t.Run("Elem", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			tmp := ""
			gosl.Test(t, "", gosl.Elem(tmp, '/', -2))
			gosl.Test(t, "", gosl.Elem(tmp, '/', -1))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 0))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 1))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 2))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 3))
		})
		t.Run("delimNotFound", func(t *testing.T) {
			tmp := "abc"
			gosl.Test(t, "", gosl.Elem(tmp, '/', -2))
			gosl.Test(t, "abc", gosl.Elem(tmp, '/', -1))
			gosl.Test(t, "abc", gosl.Elem(tmp, '/', 0))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 1))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 2))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 3))
		})

		t.Run("emptySome", func(t *testing.T) {
			tmp := "/def//ghi"
			gosl.Test(t, "", gosl.Elem(tmp, '/', -5))
			gosl.Test(t, "", gosl.Elem(tmp, '/', -4))
			gosl.Test(t, "def", gosl.Elem(tmp, '/', -3))
			gosl.Test(t, "", gosl.Elem(tmp, '/', -2))
			gosl.Test(t, "ghi", gosl.Elem(tmp, '/', -1))

			gosl.Test(t, "", gosl.Elem(tmp, '/', 0))
			gosl.Test(t, "def", gosl.Elem(tmp, '/', 1))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 2))
			gosl.Test(t, "ghi", gosl.Elem(tmp, '/', 3))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 4))
		})

		t.Run("emptyFirst", func(t *testing.T) {
			tmp := "/def/ghi"

			gosl.Test(t, "", gosl.Elem(tmp, '/', -4))
			gosl.Test(t, "", gosl.Elem(tmp, '/', -3))
			gosl.Test(t, "def", gosl.Elem(tmp, '/', -2))
			gosl.Test(t, "ghi", gosl.Elem(tmp, '/', -1))

			gosl.Test(t, "", gosl.Elem(tmp, '/', 0))
			gosl.Test(t, "def", gosl.Elem(tmp, '/', 1))
			gosl.Test(t, "ghi", gosl.Elem(tmp, '/', 2))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 3))
		})

		t.Run("emptyLast", func(t *testing.T) {
			tmp := "aa/bb/cc/"

			gosl.Test(t, "", gosl.Elem(tmp, '/', -5))
			gosl.Test(t, "aa", gosl.Elem(tmp, '/', -4))
			gosl.Test(t, "bb", gosl.Elem(tmp, '/', -3))
			gosl.Test(t, "cc", gosl.Elem(tmp, '/', -2))
			gosl.Test(t, "", gosl.Elem(tmp, '/', -1))

			gosl.Test(t, "aa", gosl.Elem(tmp, '/', 0))
			gosl.Test(t, "bb", gosl.Elem(tmp, '/', 1))
			gosl.Test(t, "cc", gosl.Elem(tmp, '/', 2))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 3))
		})

		t.Run("normal", func(t *testing.T) {
			tmp := "abc/def/ghi"

			gosl.Test(t, "", gosl.Elem(tmp, '/', -4))
			gosl.Test(t, "abc", gosl.Elem(tmp, '/', -3))
			gosl.Test(t, "def", gosl.Elem(tmp, '/', -2))
			gosl.Test(t, "ghi", gosl.Elem(tmp, '/', -1))

			gosl.Test(t, "abc", gosl.Elem(tmp, '/', 0))
			gosl.Test(t, "def", gosl.Elem(tmp, '/', 1))
			gosl.Test(t, "ghi", gosl.Elem(tmp, '/', 2))
			gosl.Test(t, "", gosl.Elem(tmp, '/', 3))
		})
	})

	t.Run("Trim", func(t *testing.T) {
		t.Run("Trim", func(t *testing.T) {
			gosl.Test(t, "abc", gosl.Trim("   abc   "))
		})
		t.Run("TrimLeft", func(t *testing.T) {
			gosl.Test(t, "abc   ", gosl.TrimLeft("   abc   "))
		})
		t.Run("TrimRight", func(t *testing.T) {
			gosl.Test(t, "   abc", gosl.TrimRight("   abc   "))
		})
	})

	t.Run("Prefix", func(t *testing.T) {
		t.Run("HasPrefix", func(t *testing.T) {
			gosl.Test(t, true, gosl.HasPrefix("/abc/def", "/"))
			gosl.Test(t, true, gosl.HasPrefix("/abc/def", "/abc"))
			gosl.Test(t, true, gosl.HasPrefix("/abc/def", ""))
			gosl.Test(t, false, gosl.HasPrefix("/abc/def", "-"))
			gosl.Test(t, false, gosl.HasPrefix("/abc/def", "/abcd"))
		})

		t.Run("HasSuffix", func(t *testing.T) {
			gosl.Test(t, true, gosl.HasSuffix("/abc/def/", "/"))
			gosl.Test(t, true, gosl.HasSuffix("/abc/def/", "/def/"))
			gosl.Test(t, true, gosl.HasSuffix("/abc/def/", ""))
			gosl.Test(t, false, gosl.HasSuffix("/abc/def/", "-"))
			gosl.Test(t, false, gosl.HasSuffix("/abc/def/", "a/def/"))
		})
		t.Run("TrimPrefix", func(t *testing.T) {
			gosl.Test(t, "abc/def/", gosl.TrimPrefix("/abc/def/", "/"))
			gosl.Test(t, "def/", gosl.TrimPrefix("/abc/def/", "/abc/"))
			gosl.Test(t, "/abc/def/", gosl.TrimPrefix("/abc/def/", ""))
			gosl.Test(t, "/abc/def/", gosl.TrimPrefix("/abc/def/", "-"))
			gosl.Test(t, "/abc/def/", gosl.TrimPrefix("/abc/def/", "a/def/"))
		})
		t.Run("TrimSuffix", func(t *testing.T) {
			gosl.Test(t, "/abc/def", gosl.TrimSuffix("/abc/def/", "/"))
			gosl.Test(t, "/abc", gosl.TrimSuffix("/abc/def/", "/def/"))
			gosl.Test(t, "/abc/def/", gosl.TrimSuffix("/abc/def/", ""))
			gosl.Test(t, "/abc/def/", gosl.TrimSuffix("/abc/def/", "-"))
			gosl.Test(t, "/ab", gosl.TrimSuffix("/abc/def/", "c/def/"))
		})
	})

	t.Run("N", func(t *testing.T) {
		t.Run("FirstN", func(t *testing.T) {
			gosl.Test(t, "abc", gosl.FirstN("abcdef", 3))
			gosl.Test(t, "abcdef", gosl.FirstN("abcdef", 10))
			gosl.Test(t, "", gosl.FirstN("abcdef", 0))
			gosl.Test(t, "", gosl.FirstN("abcdef", -1))
		})
		t.Run("LastN", func(t *testing.T) {
			gosl.Test(t, "def", gosl.LastN("abcdef", 3))
			gosl.Test(t, "abcdef", gosl.LastN("abcdef", 10))
			gosl.Test(t, "", gosl.LastN("abcdef", 0))
			gosl.Test(t, "", gosl.LastN("abcdef", -1))
		})
	})

}
