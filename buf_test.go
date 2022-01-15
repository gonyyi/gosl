// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/gonyyi/gosl"
)

func BenchmarkBuf(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	b.Run("Write()", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			buf.Write([]byte("test if this can write to buf"))
			buf = buf.WriteInt(i)
		}
	})
	// buf.Println()
}

func TestNewBuffer(t *testing.T) {
	buf := gosl.NewBuffer(10)
	buf.WriteString("test")
	gosl.Test(t, "test", buf.String())
}

func BenchmarkNewBuffer(b *testing.B) {
	b.ReportAllocs()
	buf := gosl.NewBuffer(20)
	for i := 0; i < b.N; i++ {
		buf.Reset().WriteString("test").WriteInt(i)
	}
}

func TestBuf(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)
	t.Run("Read()", func(t *testing.T) {
		tstr := "123456789_123456789_"

		t.Run("basic", func(t *testing.T) {
			buf = buf.Set(tstr)
			gosl.EOF = io.EOF
			out, err := io.ReadAll(&buf)
			gosl.Test(t, nil, err)
			gosl.Test(t, tstr, string(out))
			gosl.Test(t, 0, buf.Len())
		})

		t.Run("indivisible", func(t *testing.T) {
			b1 := buf.Set(tstr)
			b2 := bytes.NewBufferString(tstr)
			_, _ = b1, b2

			tmp1 := make([]byte, 9)
			tmp2 := make([]byte, 9)
			for {
				n1, err1 := io.ReadFull(&b1, tmp1)
				n2, err2 := io.ReadFull(b2, tmp2)
				gosl.Test(t, n1, n2)
				gosl.Test(t, err1, err2)
				if err1 != nil || err2 != nil {
					break
				}
			}
		})

		t.Run("divisible", func(t *testing.T) {
			b1 := buf.Set(tstr)
			b2 := bytes.NewBufferString(tstr)
			_, _ = b1, b2

			// Divisible size buffer
			tmp1 := make([]byte, 10)
			tmp2 := make([]byte, 10)
			for {
				n1, err1 := io.ReadFull(&b1, tmp1)
				n2, err2 := io.ReadFull(b2, tmp2)
				gosl.Test(t, n1, n2)
				gosl.Test(t, err1, err2)
				if err1 != nil || err2 != nil {
					break
				}
			}
		})

		t.Run("larger buffer than Buf", func(t *testing.T) {
			b1 := buf.Set(tstr)
			b2 := bytes.NewBufferString(tstr)
			_, _ = b1, b2

			// Divisible size buffer
			tmp1 := make([]byte, 30)
			tmp2 := make([]byte, 30)
			for {
				n1, err1 := io.ReadFull(&b1, tmp1)
				n2, err2 := io.ReadFull(b2, tmp2)
				gosl.Test(t, n1, n2)
				gosl.Test(t, err1, err2)
				if err1 != nil || err2 != nil {
					break
				}
			}
		})
	})

	t.Run("CheckPanic", func(t *testing.T) {
		var tmp gosl.Buf
		gosl.Test(t, 0, tmp.Len())
		gosl.Test(t, "", tmp.String())
		tmp = tmp.WriteString("gon")
		gosl.Test(t, 3, tmp.Len())
		gosl.Test(t, "gon", tmp.String())
		tmp = tmp.Reset().WriteFloat(1.23456, 3)
		gosl.Test(t, "1.234", tmp.String())

		tmp = tmp.AppendPrefix('-')
		gosl.Test(t, "-1.234", tmp.String())
		tmp = tmp.AppendPrefix('-') // this will be ignored since already has one
		gosl.Test(t, "-1.234", tmp.String())

		tmp = tmp.AppendSuffix(';')
		gosl.Test(t, "-1.234;", tmp.String())
		tmp = tmp.AppendSuffix(';') // this will be ignored since already has one
		gosl.Test(t, "-1.234;", tmp.String())

		tmp = tmp.AppendPrefixString(" > ")
		gosl.Test(t, " > -1.234;", tmp.String())
		tmp = tmp.AppendPrefixString(" > ") // this will be ignored since already has one
		gosl.Test(t, " > -1.234;", tmp.String())

		tmp = tmp.AppendSuffixString(" => ")
		gosl.Test(t, " > -1.234; => ", tmp.String())
		tmp = tmp.AppendSuffixString(" => ") // this will be ignored since already has one
		gosl.Test(t, " > -1.234; => ", tmp.String())
	})

	t.Run("Copy", func(t *testing.T) {
		buf1 := make(gosl.Buf, 0, 1024).Set("OK")

		buf2 := buf1.Copy()
		gosl.Test(t, buf1.String(), buf2.String())

		buf2 = buf2.WriteString(", GON!!")
		gosl.Test(t, false, buf1.String() == buf2.String())
		buf1 = buf1.WriteString(", GON!!")
		gosl.Test(t, true, buf1.String() == buf2.String())
	})

	t.Run("Elem", func(t *testing.T) {
		buf = buf.Set("/abc/def/ghi/")
		gosl.Test(t, "def", buf.Elem('/', 2))
	})

	t.Run("Equal", func(t *testing.T) {
		buf = buf.Set("ok123")
		buf2 := []byte("ok123") // pointer should be different
		gosl.Test(t, true, buf.Equal(buf2))
	})

	// t.Run("EqualString", func(t *testing.T) {
	// 	buf = buf.Set("ok123")
	// 	gosl.Test(t, true, buf.EqualString("ok123"))
	// 	gosl.Test(t, false, buf.EqualString("ok1234"))
	// })

	t.Run("Prefix", func(t *testing.T) {
		t.Run("HasPrefix", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, true, buf.HasPrefix('/', 'a'))
		})
		t.Run("TrimPrefix", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, "bc/def/", buf.TrimPrefix('/', 'a'))
		})
		t.Run("HasPrefixString", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, true, buf.HasPrefixString("/a"))
		})
		t.Run("TrimPrefixString", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, "bc/def/", buf.TrimPrefixString("/a"))
		})
	})

	t.Run("Suffix", func(t *testing.T) {
		t.Run("HasSuffix", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, true, buf.HasSuffix('/'))
		})
		t.Run("TrimSuffix", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, "/abc/de", buf.TrimSuffix('f', '/'))
		})
		t.Run("HasSuffixString", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, true, buf.HasSuffixString("/def/"))
		})
		t.Run("TrimSuffixString", func(t *testing.T) {
			buf = buf.Set("/abc/def/")
			gosl.Test(t, "/abc/", buf.TrimSuffixString("def/"))
		})
	})

	t.Run("Index", func(t *testing.T) {
		//                0123456789
		buf = buf.Set("abcdefgcex")
		tmpEmpty := make(gosl.Buf, 0, 1024)
		t.Run("Index", func(t *testing.T) {
			gosl.Test(t, 2, buf.Index('c'))
			gosl.Test(t, 2, buf.Index('c', 'd'))
			gosl.Test(t, 0, buf.Index('a', 'b'))
			gosl.Test(t, 6, buf.Index('g', 'c'))
			gosl.Test(t, 7, buf.Index('c', 'e'))
			gosl.Test(t, 9, buf.Index('x'))
			gosl.Test(t, -1, buf.Index())
			gosl.Test(t, -1, tmpEmpty.Index('x'))
			gosl.Test(t, -1, tmpEmpty.Index())
		})
		t.Run("IndexString", func(t *testing.T) {
			gosl.Test(t, 2, buf.IndexString("c"))
			gosl.Test(t, 2, buf.IndexString("cd"))
			gosl.Test(t, 0, buf.IndexString("ab"))
			gosl.Test(t, 6, buf.IndexString("gc"))
			gosl.Test(t, 7, buf.IndexString("ce"))
			gosl.Test(t, 9, buf.IndexString("x"))
			gosl.Test(t, -1, buf.IndexString(""))
			gosl.Test(t, -1, tmpEmpty.IndexString("x"))
			gosl.Test(t, -1, tmpEmpty.IndexString(""))
		})
	})

	t.Run("Insert", func(t *testing.T) {
		t.Run("Insert", func(t *testing.T) {
			buf = buf.Set("/abc/123")
			buf = buf.Insert(4, '/', 'd', 'e', 'f')
			gosl.Test(t, "/abc/def/123", buf.String())

			buf = buf.Set("/abc/123")
			buf = buf.Insert(20, '/', 'd', 'e', 'f')
			gosl.Test(t, "/abc/123/def", buf.String())

			buf = buf.Set("/abc/123")
			buf = buf.Insert(0, '/', 'd', 'e', 'f')
			gosl.Test(t, "/def/abc/123", buf.String())

			buf = buf.Set("/abc/123")
			buf = buf.Insert(-10, '/', 'd', 'e', 'f')
			gosl.Test(t, "/def/abc/123", buf.String())
		})

		t.Run("InsertString", func(t *testing.T) {
			buf = buf.Set("/abc/123")
			buf = buf.InsertString(4, "/def")
			gosl.Test(t, "/abc/def/123", buf.String())

			buf = buf.Set("/abc/123")
			buf = buf.InsertString(20, "/def")
			gosl.Test(t, "/abc/123/def", buf.String())

			buf = buf.Set("/abc/123")
			buf = buf.InsertString(0, "/def")
			gosl.Test(t, "/def/abc/123", buf.String())

			buf = buf.Set("/abc/123")
			buf = buf.InsertString(-10, "/def")
			gosl.Test(t, "/def/abc/123", buf.String())
		})

	})

	t.Run("Replace", func(t *testing.T) {
		buf = buf.Set("/abc_def/")
		gosl.Test(t, "/abc def/", buf.Replace('_', ' '))
	})

	t.Run("Reverse", func(t *testing.T) {
		buf = buf.Set("123456")
		gosl.Test(t, "654321", buf.Reverse())
	})

	t.Run("Shift", func(t *testing.T) {
		gosl.Test(t, "012345", buf.Set("012345").Shift(2, 3, -3)) // outside the range
		gosl.Test(t, "234015", buf.Set("012345").Shift(2, 3, -2))
		gosl.Test(t, "023415", buf.Set("012345").Shift(2, 3, -1))
		gosl.Test(t, "015234", buf.Set("012345").Shift(2, 3, 1))
		gosl.Test(t, "012345", buf.Set("012345").Shift(2, 3, 2)) // outside the range
		gosl.Test(t, "012345", buf.Set("012345").Shift(2, 3, 0))
	})

	t.Run("LowerUpper", func(t *testing.T) {
		gosl.Test(t, "hello gon", buf.Set("Hello Gon").ToLower())
		gosl.Test(t, "HELLO GON", buf.Set("Hello Gon").ToUpper())
	})

	t.Run("Buf-Unique", func(t *testing.T) {
		t.Run("(*Buf).Buf", func(t *testing.T) {
			buf = buf.WriteString("abc")
			var act []byte
			act = buf.Bytes()
			gosl.Test(t, true, buf.String() == string(act))
		})

		t.Run("(*Buf).Cap", func(t *testing.T) {
			buf = buf.Set("abc")
			gosl.Test(t, cap(buf), buf.Cap())
		})

		t.Run("(*Buf).Len", func(t *testing.T) {
			buf = buf.Set("abc")
			gosl.Test(t, len(buf), buf.Len())
		})

		t.Run("(*Buf).Println", func(t *testing.T) {
			t.Skip("this test will print to stdout")
			buf = buf.Set("abc")
			// buf.Println()
		})

		t.Run("(*Buf).Reset", func(t *testing.T) {
			buf = buf.Set("abc")
			buf = buf.Reset()
			gosl.Test(t, "", buf.String())
		})

		t.Run("(*Buf).Set", func(t *testing.T) {
			buf = buf.Set("abc")
			gosl.Test(t, "def", buf.Set("def").String())
		})

		t.Run("(*Buf).String", func(t *testing.T) {
			buf = buf.Set("abcd")
			gosl.Test(t, "abcd", buf.String())
		})

		t.Run("(*Buf).Write", func(t *testing.T) {
			buf = buf.Set("abc")
			n, err := buf.Write([]byte("hi"))
			if err != nil {
				t.Errorf(err.Error())
			}
			if n != 2 {
				t.Fail()
			}
			gosl.Test(t, "abchi", buf.String())
		})

		t.Run("(*Buf).WriteBytes", func(t *testing.T) {
			buf = buf.Set("abc")
			buf = buf.WriteBytes('1', '2')
			gosl.Test(t, "abc12", buf.String())
		})

		t.Run("(*Buf).WriteString", func(t *testing.T) {
			buf = buf.Set("abc")
			buf = buf.WriteString("123")
			gosl.Test(t, "abc123", buf.String())
		})

		t.Run("(*Buf).WriteTo", func(t *testing.T) {
			tmpBuf := make(gosl.Buf, 0, 512)

			buf = buf.Set("abc")
			buf = buf.WriteString("123") // buf = abc123

			n, err := buf.WriteTo(&tmpBuf)
			if n != 6 {
				t.Fail()
			}
			if err != nil {
				t.Error(err.Error())
			}
			gosl.Test(t, buf.String(), tmpBuf.String())
		})
	})

	t.Run("Buf-FnBytes", func(t *testing.T) {
		// This is a brief check, detailed one can be found from <fBytes_test.go>
		gosl.Test(t, "truefalse", buf.Reset().WriteBool(true).WriteBool(false))

	})
}

func BenchmarkBuffer(b *testing.B) {
	sample := make(gosl.Buf, 0, 1024)

	b.Run("T1", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := gosl.GetBuffer()
			buf.Buf = buf.Buf.Set("test yo!").WriteInt(i)
			sample = sample.Reset().WriteBytes(buf.Buf...)
			gosl.PutBuffer(buf)
		}
		// sample.Println()
	})
}

func TestBuffer(t *testing.T) {
	b1 := gosl.GetBuffer()
	t.Run("Bytes()", func(t *testing.T) {
		b1.Set("test1")
		tmp := b1.Bytes()
		gosl.Test(t, true, gosl.BytesEqual(b1.Buf, tmp))
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
		b1.WriteBytes('1', '2', '3')
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
