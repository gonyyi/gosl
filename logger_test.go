// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)
func Test_Logger(t *testing.T) {
	t.Run("SetOutput", func(t *testing.T) {
		var buf gosl.Buffer
		log := gosl.Logger{}.SetOutput(&buf).Enable(false)
		log.String("You never see me")
		gosl.TestString(t, "", buf.String())
		buf.Reset()

		log = log.SetOutput(&buf)
		log.String("message from buffer")
		gosl.TestString(t, "message from buffer\n", buf.String())
	})

	t.Run("As-a-Writer", func(t *testing.T) {
		var buf gosl.Buffer
		log := gosl.Logger{}.SetOutput(&buf)
		fmt.Fprintf(log, "Hello <%s>\n", "Gon!")
		gosl.TestString(t, "Hello <Gon!>\n", buf.String())
	})
}

func Benchmark_Logger(b *testing.B) {
	var show = false

	s := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg"}
	e := []error{gosl.NewError("err-111"), gosl.NewError("err-222")}
	e[1] = nil

	b.Run("combined", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buffer{}
		l := gosl.Logger{}.SetOutput(&buf)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			l.String(s[i%7])
			l.KeyBool(s[i%7], i%3==0)
			l.KeyInt(s[i%7], i)
			l.KeyFloat64(s[i%7], float64(i)+0.1234)
			l.KeyString(s[i%7], s[i%7])
			l.KeyError(s[i%7], e[i%2])
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("String", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buffer{}
		l := gosl.Logger{}.SetOutput(&buf)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			l.String(s[i%7])
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyBool", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buffer{}
		l := gosl.Logger{}.SetOutput(&buf)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			l.KeyBool(s[i%7], i%3==0)
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyInt", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buffer{}
		l := gosl.Logger{}.SetOutput(&buf)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			l.KeyInt(s[i%7], i)
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyFloat64", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buffer{}
		l := gosl.Logger{}.SetOutput(&buf)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			l.KeyFloat64(s[i%7], float64(i)+0.1234)
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyString", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buffer{}
		l := gosl.Logger{}.SetOutput(&buf)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			l.KeyString(s[i%7], s[i%7])
		}
		if show {
			print(buf.String())
		}
	})

	b.Run("KeyError", func(b *testing.B) {
		b.ReportAllocs()
		buf := gosl.Buffer{}
		l := gosl.Logger{}.SetOutput(&buf)
		for i := 0; i < b.N; i++ {
			buf.Reset()
			l.KeyError(s[i%7], e[i%2])
		}
		if show {
			print(buf.String())
		}
	})
}
