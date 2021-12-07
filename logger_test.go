// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/30/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)

func TestNewLogger(t *testing.T) {
	buf := make(gosl.Buf, 0, 512)
	l := gosl.NewLogger(&buf)
	l.String("test123")
	gosl.Test(t, "test123\n", buf.String())

	buf = buf.Reset()
	l = gosl.NewLogger(&buf).SetPrefix("GON.")
	l.String("test123")
	gosl.Test(t, "GON.test123\n", buf.String())

	buf = buf.Reset()
	l = l.SetPrefix("")
	l.String("test123")
	gosl.Test(t, "test123\n", buf.String())

	buf = buf.Reset()
	l = l.SetOutput(nil)
	l.String("test123")
	gosl.Test(t, "", buf.String())

	l = l.Enable(true)
	l.String("test1234")
	gosl.Test(t, "", buf.String())

	l = l.SetOutput(&buf)
	err1 := gosl.NewError("testError1")
	l.IfErr("test", err1)
	gosl.Test(t, "test -> (err) testError1\n", buf.String())

	buf = buf.Reset()
	l.IfErr("test", nil)
	gosl.Test(t, "", buf.String())

	buf = buf.Reset()
	l.KeyBool("1", true)
	l.KeyBool("0", false)
	gosl.Test(t, "1: true\n0: false\n", buf.String())

	buf = buf.Reset()
	l.KeyError("1", err1)
	l.KeyError("0", nil)
	gosl.Test(t, "1 -> (err) testError1\n0 -> <nil>\n", buf.String())

	buf = buf.Reset()
	l = l.Enable(false)
	l.Write([]byte("test"))
	gosl.Test(t, "", buf.String())

	l = l.SetOutput(&buf).SetPrefix("gon.")
	l.IfErr("test", err1)
	gosl.Test(t, "gon.test -> (err) testError1\n", buf.String())

	l = l.Enable(true)

	gosl.Test(t, true, l.Enabled())
	err1 = l.Close()
	gosl.Test(t, nil, err1)
}

func Test_Logger_SetPrefix(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var buf gosl.Buf
		log := gosl.Logger{}.SetOutput(&buf).SetPrefix("t1.")
		log.String("hello1")
		log2 := log.SetPrefix("t2.")
		log2.String("hello2")
		log.String("bye1")
		log2.String("bye2")
		gosl.Test(t, "t1.hello1\nt2.hello2\nt1.bye1\nt2.bye2\n", buf.String())
	})

	t.Run("all", func(t *testing.T) {
		var buf gosl.Buf
		log := gosl.Logger{}.SetOutput(&buf).SetPrefix("t1.")
		log.String("hello1")
		log.KeyString("s", "val1")
		log.KeyBool("b", false)
		log.KeyInt("i", 123)
		log.KeyFloat64("f", 123.123)
		log.KeyError("key", gosl.NewError("err"))
		log.Write([]byte("bytes"))
		gosl.Test(t, "t1.hello1\nt1.s: \"val1\"\nt1.b: false\nt1.i: 123\nt1.f: 123.12\nt1.key -> (err) err\nt1.bytes", buf.String())
		// t1.hello1
		// t1.s: "val1"
		// t1.b: false
		// t1.i: 123
		// t1.f: 123.12
		// t1.key -> (err) err
		// t1.bytes
	})
}

func Test_Logger(t *testing.T) {
	t.Run("SetOutput", func(t *testing.T) {
		var buf gosl.Buf
		log := gosl.Logger{}.SetOutput(&buf).Enable(false)
		log.String("You never see me")
		gosl.Test(t, "", buf.String())
		buf.Reset()

		log = log.SetOutput(&buf)
		log.String("message from Buffer")
		gosl.Test(t, "message from Buffer\n", buf.String())
	})

	t.Run("As-a-Writer", func(t *testing.T) {
		var buf gosl.Buf
		log := gosl.Logger{}.SetOutput(&buf)
		fmt.Fprintf(log, "Hello <%s>\n", "Gon!")
		gosl.Test(t, "Hello <Gon!>\n", buf.String())
	})
}

func Benchmark_Logger(b *testing.B) {
	var show = false

	s := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg"}
	e := []error{gosl.NewError("err-111"), gosl.NewError("err-222")}
	e[1] = nil

	b.Run("combined: enabled", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b1.")
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.String(s[i%7])
			l.KeyBool(s[i%7], i%3 == 0)
			l.KeyInt(s[i%7], i)
			l.KeyFloat64(s[i%7], float64(i)+0.1234)
			l.KeyString(s[i%7], s[i%7])
			l.KeyError(s[i%7], e[i%2])
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("combined: disabled", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b2.")
		l = l.Enable(false)
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.String(s[i%7])
			l.KeyBool(s[i%7], i%3 == 0)
			l.KeyInt(s[i%7], i)
			l.KeyFloat64(s[i%7], float64(i)+0.1234)
			l.KeyString(s[i%7], s[i%7])
			l.KeyError(s[i%7], e[i%2])
		}
		if 1 == 1 {
			print(buf.String())
		}
	})

	b.Run("String", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b3.")
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.String(s[i%7])
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyBool", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b4.")
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.KeyBool(s[i%7], i%3 == 0)
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyInt", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b5.")
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.KeyInt(s[i%7], i)
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyFloat64", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b6.")
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.KeyFloat64(s[i%7], float64(i)+0.1234)
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyString", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b7.")
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.KeyString(s[i%7], s[i%7])
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyError", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buf{}
		l := gosl.Logger{}.SetOutput(&buf).SetPrefix("b8.")
		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.KeyError(s[i%7], e[i%2])
		}
		if show {
			print(buf.String())
		}
	})
}


func TestLogger_KeySize(t *testing.T) {
	buf := gosl.Buf{}
	l := gosl.Logger{}.SetOutput(&buf)
	l.KeySize("buffer", 23456234523)
	gosl.Test(t, "buffer: 21.84GB\n", buf.String())
}

func BenchmarkLogger_KeySize(b *testing.B) {
	show := false
	b.ReportAllocs()
	buf := gosl.Buf{}
	l := gosl.Logger{}.SetOutput(&buf)
	for i := 0; i < b.N; i++ {
		buf = buf.Reset()
		l.KeySize("buf", int64(i))
	}
	if show {
		print(buf.String())
	}
}
