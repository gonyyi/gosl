// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

func TestBytesAppends(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("BytesAppendBool", func(t *testing.T) {
		gosl.Test(t, "true", string(gosl.BytesAppendBool(buf.Reset(), true)))
		gosl.Test(t, "false", string(gosl.BytesAppendBool(buf.Reset(), false)))
	})

	t.Run("BytesAppendFloat", func(t *testing.T) {
		gosl.Test(t, "0.00123", string(gosl.BytesAppendFloat(buf.Reset(), 0.0012345, 5)))
		gosl.Test(t, "-0.00123", string(gosl.BytesAppendFloat(buf.Reset(), -0.0012345, 5)))

		gosl.Test(t, "123.12345", string(gosl.BytesAppendFloat(buf.Reset(), 123.123456789, 5)))
		gosl.Test(t, "-123.12345", string(gosl.BytesAppendFloat(buf.Reset(), -123.123456789, 5)))

		gosl.Test(t, "123.00123", string(gosl.BytesAppendFloat(buf.Reset(), 123.0012345, 5)))
		gosl.Test(t, "-123.00123", string(gosl.BytesAppendFloat(buf.Reset(), -123.0012345, 5)))

		// gosl.Test(t, "123.000001234", string(gosl.BytesAppendFloat(buf.Reset(), 123.0000012345, 9)))
		// gosl.Test(t, "-123.000001234", string(gosl.BytesAppendFloat(buf.Reset(), -123.0000012345, 9)))
		// gosl.Test(t, "123.000001234500", string(gosl.BytesAppendFloat(buf.Reset(), 123.0000012345, 12)))
		// gosl.Test(t, "-123.000001234500", string(gosl.BytesAppendFloat(buf.Reset(), -123.0000012345, 12)))
	})

	t.Run("BytesAppendInt", func(t *testing.T) {
		gosl.Test(t, "123", string(gosl.BytesAppendInt(buf.Reset(), 123)))
		gosl.Test(t, "-123", string(gosl.BytesAppendInt(buf.Reset(), -123)))
	})

	t.Run("BytesAppendPrefix", func(t *testing.T) {
		// AppendPrefix will not append if prefix already exists
		buf = buf.Reset()
		gosl.Test(t, "Gon:Hi", string(gosl.BytesAppendPrefix(buf.WriteString("Hi"), []byte("Gon:")...)))
		buf = buf.Reset()
		gosl.Test(t, "Gon:Hi", string(gosl.BytesAppendPrefix(buf.WriteString("Gon:Hi"), []byte("Gon:")...)))
	})

	t.Run("BytesAppendPrefixString", func(t *testing.T) {
		// AppendPrefix will not append if prefix already exists
		buf = buf.Reset()
		gosl.Test(t, "Gon:Hi", string(gosl.BytesAppendPrefixString(buf.WriteString("Hi"), "Gon:")))
		buf = buf.Reset()
		gosl.Test(t, "Gon:Hi", string(gosl.BytesAppendPrefixString(buf.WriteString("Gon:Hi"), "Gon:")))
	})

	t.Run("BytesAppendSize", func(t *testing.T) {
		buf = buf.Reset()
		buf = gosl.BytesAppendSize(buf, 123456*1024*1024, 2)
		gosl.Test(t, "120.56GB", buf.String())

		buf = buf.Reset()
		buf = gosl.BytesAppendFloat(buf, float64(123456*1024*1024)/float64(gosl.GB), 2)
		gosl.Test(t, "120.56", buf.String())

		buf = buf.Reset()
		buf = gosl.BytesAppendInt(buf, int(123456*1024*1024/gosl.GB))
		gosl.Test(t, "120", buf.String())
	})

	t.Run("BytesAppendSizeIn", func(t *testing.T) {
		buf = buf.Reset()
		buf = gosl.BytesAppendSizeIn(buf, 123456*1024*1024, gosl.MB, 2)
		gosl.Test(t, "123456.00MB", buf.String())
	})

	t.Run("BytesAppendSuffix", func(t *testing.T) {
		// AppendPrefix will not append if prefix already exists
		buf = buf.Reset()
		gosl.Test(t, "/gon/yi/", string(gosl.BytesAppendSuffix(buf.WriteString("/gon/yi"), '/')))
		buf = buf.Reset()
		gosl.Test(t, "/gon/yi/", string(gosl.BytesAppendSuffix(buf.WriteString("/gon/yi/"), '/')))
	})

	t.Run("BytesAppendSuffixString", func(t *testing.T) {
		// AppendPrefix will not append if prefix already exists
		buf = buf.Reset()
		gosl.Test(t, "/gon/yi/", string(gosl.BytesAppendSuffixString(buf.WriteString("/gon/yi"), "/")))
		buf = buf.Reset()
		gosl.Test(t, "/gon/yi/", string(gosl.BytesAppendSuffixString(buf.WriteString("/gon/yi/"), "/")))
	})

	t.Run("Hex", func(t *testing.T) {
		// BytesToHex
		// HexToBytes

		t.Run("BytesToHex", func(t *testing.T) {
			buf = buf.Reset()
			buf = gosl.BytesToHex(buf, []byte("Hello Gon"))
			gosl.Test(t, "48656c6c6f20476f6e", buf.String())
		})
		t.Run("HexToBytes", func(t *testing.T) {
			buf = buf.Reset()
			buf, _ = gosl.HexToBytes(buf, []byte("48656c6c6f20476f6e"))
			gosl.Test(t, "Hello Gon", buf.String())
		})
	})
}

func TestBytesEquals(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024).WriteString("gon")
	bufEq := make(gosl.Buf, 0, 1024).WriteString("gon")
	bufDf := make(gosl.Buf, 0, 1024).WriteString("yi")

	t.Run("BytesEqual", func(t *testing.T) {
		gosl.Test(t, true, gosl.BytesEqual(buf, bufEq))
		gosl.Test(t, false, gosl.BytesEqual(buf, bufDf))
		gosl.Test(t, false, gosl.BytesEqual(buf, nil))
	})

	// t.Run("BytesEqualString", func(t *testing.T) {
	// 	gosl.Test(t, true, gosl.BytesEqualString(buf, bufEq.String()))
	// 	gosl.Test(t, false, gosl.BytesEqualString(buf, bufDf.String()))
	// 	gosl.Test(t, false, gosl.BytesEqualString(buf, ""))
	// })
}

func TestBytesFilterAny(t *testing.T) {
	var buf gosl.Buf

	buf = buf.Set("a1b2c3d4")
	buf = gosl.BytesFilterAny(buf, "1234567890", true)
	gosl.Test(t, "1234", buf.String())

	buf = buf.Set("a1b2c3d4")
	buf = gosl.BytesFilterAny(buf, "1234567890", false)
	gosl.Test(t, "abcd", buf.String())

}

func TestBytesHases(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024).WriteString("/gon:")

	t.Run("BytesHasPrefix", func(t *testing.T) {
		gosl.Test(t, true, gosl.BytesHasPrefix(buf, '/'))
		gosl.Test(t, false, gosl.BytesHasPrefix(buf, ':'))
	})
	t.Run("BytesHasPrefixString", func(t *testing.T) {
		gosl.Test(t, true, gosl.BytesHasPrefixString(buf, "/"))
		gosl.Test(t, false, gosl.BytesHasPrefixString(buf, ":"))
	})

	t.Run("BytesHasSuffix", func(t *testing.T) {
		gosl.Test(t, true, gosl.BytesHasSuffix(buf, ':'))
		gosl.Test(t, false, gosl.BytesHasSuffix(buf, '/'))
	})
	t.Run("BytesHasSuffixString", func(t *testing.T) {
		gosl.Test(t, true, gosl.BytesHasSuffixString(buf, ":"))
		gosl.Test(t, false, gosl.BytesHasSuffixString(buf, "/"))
	})
}

func TestBytesIndexes(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024).WriteString("/gon:")

	t.Run("BytesIndex", func(t *testing.T) {
		gosl.Test(t, 1, gosl.BytesIndex(buf, 'g', 'o'))
		gosl.Test(t, -1, gosl.BytesIndex(buf, 'g', 'n'))
	})

	t.Run("BytesIndexString", func(t *testing.T) {
		gosl.Test(t, 1, gosl.BytesIndexString(buf, "go"))
		gosl.Test(t, -1, gosl.BytesIndexString(buf, "gn"))
	})
}

func TestBytesInserts(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("BytesIndex", func(t *testing.T) {
		buf = buf.Reset().WriteString("/gon:")
		gosl.Test(t, "/vm-gon:", string(gosl.BytesInsert(buf, 1, []byte("vm-")...)))
		buf = buf.Reset().WriteString("/gon:")
		gosl.Test(t, "/test/gon:", string(gosl.BytesInsert(buf, 0, []byte("/test")...)))
	})

	t.Run("BytesIndexString", func(t *testing.T) {
		buf = buf.Reset().WriteString("/gon:")
		gosl.Test(t, "/vm-gon:", string(gosl.BytesInsertString(buf, 1, "vm-")))
		buf = buf.Reset().WriteString("/gon:")
		gosl.Test(t, "/test/gon:", string(gosl.BytesInsertString(buf, 0, "/test")))
	})
}

func TestBytesTos(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("BytesToLower", func(t *testing.T) {
		buf = buf.Reset().WriteString("abcGONdef")
		gosl.BytesToLower(buf)
		gosl.Test(t, "abcgondef", buf.String())
	})

	t.Run("BytesToUpper", func(t *testing.T) {
		buf = buf.Reset().WriteString("abcGONdef")
		gosl.BytesToUpper(buf)
		gosl.Test(t, "ABCGONDEF", buf.String())
	})
}

func TestBytesTrims(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("BytesTrimPrefix", func(t *testing.T) {
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "gonyyi.com/test", string(gosl.BytesTrimPrefix(buf, []byte("http://")...)))

		// Case where prefix not found
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "http://gonyyi.com/test", string(gosl.BytesTrimPrefix(buf, []byte("https://")...)))
	})

	t.Run("BytesTrimPrefixString", func(t *testing.T) {
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "gonyyi.com/test", string(gosl.BytesTrimPrefixString(buf, "http://")))

		// Case where prefix not found
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "http://gonyyi.com/test", string(gosl.BytesTrimPrefixString(buf, "https://")))
	})

	t.Run("BytesTrimSuffix", func(t *testing.T) {
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "http://gonyyi.com", string(gosl.BytesTrimSuffix(buf, []byte("/test")...)))

		// Case where prefix not found
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "http://gonyyi.com/test", string(gosl.BytesTrimSuffix(buf, []byte("test/")...)))
	})

	t.Run("BytesTrimSuffixString", func(t *testing.T) {
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "http://gonyyi.com", string(gosl.BytesTrimSuffixString(buf, "/test")))

		// Case where prefix not found
		buf = buf.Reset().WriteString("http://gonyyi.com/test")
		gosl.Test(t, "http://gonyyi.com/test", string(gosl.BytesTrimSuffixString(buf, "/test/")))
	})
}

func TestBytesElem(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("abcd", func(t *testing.T) {
		buf = buf.Reset().WriteString("abcd")
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -3)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -2)))
		gosl.Test(t, "abcd", string(gosl.BytesElem(buf, '/', -1)))
		gosl.Test(t, "abcd", string(gosl.BytesElem(buf, '/', 0)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 1))) // doesn't exist
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 2))) // doesn't exist
	})

	t.Run("abc/def/ghi", func(t *testing.T) {
		buf = buf.Reset().WriteString("abc/def/ghi")
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -5))) // doesn't exist
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -4))) // doesn't exist
		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', -3)))
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', -2)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', -1)))

		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', 0)))
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', 1)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', 2)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 3))) // this doesn't exist
	})

	t.Run("abc/def/ghi/", func(t *testing.T) {
		buf = buf.Reset().WriteString("abc/def/ghi/")
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -6)))    // doesn't exist
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -5)))    // doesn't exist
		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', -4))) // doesn't exist
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', -3)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', -2)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -1)))

		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', 0)))
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', 1)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', 2)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 3)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 4))) // this doesn't exist
	})

	t.Run("/abc/def/ghi", func(t *testing.T) {
		buf = buf.Reset().WriteString("/abc/def/ghi")
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -6))) // doesn't exist
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -5))) // doesn't exist
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -4)))
		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', -3)))
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', -2)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', -1)))

		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 0)))
		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', 1)))
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', 2)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', 3)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 4))) // this doesn't exist
	})

	t.Run("/abc/def/ghi/", func(t *testing.T) {
		buf = buf.Reset().WriteString("/abc/def/ghi/")
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -7))) // doesn't exist
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -6))) // doesn't exist
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -5)))
		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', -4)))
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', -3)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', -2)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', -1)))

		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 0)))
		gosl.Test(t, "abc", string(gosl.BytesElem(buf, '/', 1)))
		gosl.Test(t, "def", string(gosl.BytesElem(buf, '/', 2)))
		gosl.Test(t, "ghi", string(gosl.BytesElem(buf, '/', 3)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 4)))
		gosl.Test(t, "", string(gosl.BytesElem(buf, '/', 5))) // this doesn't exist
	})
}

func TestBytesEtc(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("BytesLastByte", func(t *testing.T) {
		buf = buf.Reset().WriteString("Hello there")
		gosl.Test(t, byte('e'), gosl.BytesLastByte(buf))

		buf = buf.Reset()
		gosl.Test(t, byte(0), gosl.BytesLastByte(buf))
	})

	t.Run("BytesCopy", func(t *testing.T) {
		buf = buf.Reset().WriteString("Hello there")
		buf2 := gosl.BytesCopy(buf)

		// pointer should be different
		bufBytes := buf.Bytes()
		gosl.Test(t, false, &bufBytes == &buf2)

		gosl.Test(t, "Hello there", string(buf2))

		// When original is modified
		buf = buf.InsertString(5, ",")
		gosl.Test(t, "Hello, there", string(buf))
		gosl.Test(t, "Hello there", string(buf2))

		// When copy is modified
		buf2 = buf2[:0] // clear buf2
		gosl.Test(t, "Hello, there", string(buf))
		gosl.Test(t, "", string(buf2))
	})

	t.Run("BytesReplace", func(t *testing.T) {
		buf = buf.Reset().WriteString("Hello there")
		gosl.BytesReplace(buf, ' ', '-')
		gosl.Test(t, "Hello-there", string(buf))
	})

	t.Run("BytesReverse", func(t *testing.T) {
		// return value can be ignored
		buf = buf.Reset().WriteString("12345")
		gosl.BytesReverse(buf)
		gosl.Test(t, "54321", buf.String())
	})

	t.Run("BytesShift", func(t *testing.T) {
		var ok bool

		//                01234567890
		buf = buf.Set("123-456-789")
		ok = gosl.BytesShift(buf, 4, 4, -4)
		gosl.Test(t, "456-123-789", buf.String())
		gosl.Test(t, true, ok)

		//                01234567890
		buf = buf.Set("123-456-789")
		ok = gosl.BytesShift(buf, 4, 8, -4)
		gosl.Test(t, "123-456-789", buf.String())
		gosl.Test(t, false, ok)

		//                01234567890
		buf = buf.Set("123-456-789")
		ok = gosl.BytesShift(buf, 4, 7, -4)
		gosl.Test(t, "456-789123-", buf.String())
		gosl.Test(t, true, ok)
	})
}

func BenchmarkBytesAppends(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	b.Run("BytesAppendBool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesAppendBool(buf.Reset(), true)
		}
	})

	b.Run("BytesAppendInt", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesAppendInt(buf.Reset(), i)
		}
		// buf.Println()
	})

	b.Run("BytesAppendPrefix", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesAppendPrefix(buf.Set("test"), 'O', 'K', '-')
		}
		// buf.Println()
	})

	b.Run("BytesAppendPrefixString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesAppendPrefixString(buf.Set("test"), "OK-")
		}
		// buf.Println()
	})

	b.Run("BytesAppendSuffix", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesAppendSuffix(buf.Set("test"), ':', 'O', 'K')
		}
		// buf.Println()
	})

	b.Run("BytesAppendSuffixString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = gosl.BytesAppendSuffixString(buf.Set("test"), ":OK")
		}
		// buf.Println()
	})
}

func BenchmarkBytesEquals(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024).WriteString("gon")
	bufEq := make(gosl.Buf, 0, 1024).WriteString("gon")
	bufDf := make(gosl.Buf, 0, 1024).WriteString("yi")

	var out bool
	_ = out

	b.Run("BytesEqual", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out = gosl.BytesEqual(buf, bufEq)
		}
		// println(out)
	})
	b.Run("BytesEqual:false", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out = gosl.BytesEqual(buf, bufDf)
		}
		// println(out)
	})
}

func BenchmarkBytesFilterAny(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	b.Run("keep=true", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = buf.Set("a1b2c3d4")
			buf = gosl.BytesFilterAny(buf, "1234567890", false)
		}
	})
	// buf.Println()
	b.Run("keep=false", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = buf.Set("a1b2c3d4")
			buf = gosl.BytesFilterAny(buf, "1234567890", true)
		}
	})
	// buf.Println()
}

func BenchmarkBytesHases(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024).WriteString("/gon:")
	var out bool
	_ = out

	b.Run("BytesHasPrefix", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out = gosl.BytesHasPrefix(buf, '/')
		}
	})
	b.Run("BytesHasPrefixString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out = gosl.BytesHasPrefixString(buf, "/")
		}
	})
	b.Run("BytesHasSuffix", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out = gosl.BytesHasSuffix(buf, '/')
		}
	})
	b.Run("BytesHasSuffixString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			out = gosl.BytesHasSuffixString(buf, "/")
		}
	})
}

func BenchmarkBytesIndexes(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024).WriteString("/gon:")

	b.Run("BytesIndex", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			gosl.BytesIndex(buf, 'g', 'o')
		}
	})

	b.Run("BytesIndexString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			gosl.BytesIndexString(buf, "go")
		}
	})
}

func BenchmarkBytesInserts(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	buf = buf.Set("/gon:")

	b.Run("BytesInsert", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// buf = buf.Set("/gon:")
			gosl.BytesInsert(buf, 1, []byte("vm-")...)
		}
	})

	b.Run("BytesInsertString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// buf = buf.Set("/gon:")
			gosl.BytesInsertString(buf, 1, "vm-")
		}
	})
}

func BenchmarkBytesTos(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024).Set("abcGONdef")

	b.Run("BytesToLower", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			gosl.BytesToLower(buf)
		}
	})
	b.Run("BytesToUpper", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			gosl.BytesToUpper(buf)
		}
	})
}
func BenchmarkBytesTrims(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	buf = buf.Set("http://gonyyi.com/test")
	buf2 := make(gosl.Buf, 0, 1024)
	_ = buf2

	b.Run("BytesTrimPrefix", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf2 = gosl.BytesTrimPrefix(buf, []byte("http://")...)
		}
		// buf.Println()
		// buf2.Println()
	})

	b.Run("BytesTrimPrefixString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf2 = gosl.BytesTrimPrefixString(buf, "http://")
		}
		// buf.Println()
		// buf2.Println()
	})

	b.Run("BytesTrimSuffix", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf2 = gosl.BytesTrimSuffix(buf, []byte("/test")...)
		}
		// buf.Println()
		// buf2.Println()
	})

	b.Run("BytesTrimSuffixString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf2 = gosl.BytesTrimSuffixString(buf, "/test")
		}
		// buf.Println()
		// buf2.Println()
	})
}

func BenchmarkBytesElem(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024).Set("/abc/ghi")
	buf2 := make(gosl.Buf, 0, 1024)
	_ = buf2

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf2 = gosl.BytesElem(buf, '/', 2)
		// buf2 = buf2.Set(string(buf2))
	}
}

func BenchmarkBytesEtc(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	buf2 := make(gosl.Buf, 0, 1024)
	buf = buf.Set("Hello there")
	_ = buf2

	b.Run("BytesLastByte", func(b *testing.B) {
		b.ReportAllocs()
		var c byte
		_ = c
		for i := 0; i < b.N; i++ {
			buf = buf.Set("Hello there")
			c = gosl.BytesLastByte(buf)
			// buf = append(buf, gosl.BytesLastByte(buf))
		}
		// println(string(c))
	})

	// BytesCopy is a deep copy and will need alloc
	b.Run("BytesCopy", func(b *testing.B) {
		b.SkipNow()

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = buf.Set("Hello there")
			buf2 = gosl.BytesCopy(buf)
		}
		// println(string(c))
	})

	b.Run("BytesReplace", func(b *testing.B) {
		b.ReportAllocs()
		//                01234567890
		buf = buf.Set("Hello there")
		for i := 0; i < b.N; i++ {
			if buf.Index(' ') > 0 {
				gosl.BytesReplace(buf, ' ', '-')
			} else {
				gosl.BytesReplace(buf, '-', ' ')
			}
		}
		// buf.Println()
	})

	b.Run("BytesReverse", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			for j := 0; j < i%3; j++ {
				gosl.BytesReverse(buf)
			}
		}
		// buf.Println()
	})

	b.Run("BytesShift", func(b *testing.B) {
		b.ReportAllocs()
		//                01234567890
		buf = buf.Set("Hello there")
		for i := 0; i < b.N; i++ {
			gosl.BytesShift(buf, 5, 6, -2)
		}
		// buf.Println()
	})
}
