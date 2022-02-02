// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 02/01/2022

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
	"time"
)

func TestLvWriter(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("Custom LvLevel", func(t *testing.T) {
		type MyLv = uint8
		const (
			MyTrace MyLv = iota
			MyDebug
			MyInfo
			MyOkay
			MyWarn
			MyError
			MyHell
		)
		buf = buf.Reset()
		//w := gosl.NewLvWriter(&buf, uint8(clv1))
		w := gosl.NewLvWriter(&buf, MyWarn)

		w.Lv(MyTrace).WriteString("MyTRACE")
		w.Lv(MyDebug).WriteString("MyDEBUG")
		w.Lv(MyInfo).WriteString("MyINFO")
		w.Lv(MyOkay).WriteString("MyOKAY")
		w.Lv(MyWarn).WriteString("MyWARN") // Since level is set to MyWarn, everything here on will be printed.
		w.Lv(MyError).WriteString("MyERROR")
		w.Lv(MyHell).WriteString("MyHELL")
		gosl.Test(t, "MyWARN\nMyERROR\nMyHELL\n", buf.String())
	})

	t.Run("Basic", func(t *testing.T) {
		buf = buf.Reset()
		w := gosl.LvWriter{}.SetOutput(&buf) // if level is not set, it will be 0.
		if w.Output() != &buf {
			t.Errorf("Output() should return &buf here")
		}
		if w.Enabled() == false {
			t.Errorf("Enabled() should return true")
		}

		w = w.Enable(false) // now, it's disabled, but writer is still there
		if w.Enabled() != false {
			t.Errorf("Enabled() should return false")
		}
		if w.Output() != &buf {
			t.Errorf("Output() should be same as &buf")
		}

		w = w.SetOutput(nil)
		if w.Output() != nil {
			t.Errorf("Output() should be nil")
		}
		if w.Enabled() == true {
			t.Errorf("Enabled() should return false")
		}

		buf.Reset()
		w = w.SetOutput(&buf).SetLevel(gosl.LvInfo)
		w.WriteString("abc") // if level not set, it will be always printed
		if buf.String() != "abc\n" {
			t.Errorf("Unexpected: <%s>", buf.String())
		}

		buf = buf.Reset()
		w.Trace().WriteString("not this")
		if buf.String() != "" {
			t.Errorf("Unexpected: <%s>", buf.String())
		}

		if err := w.Close(); err != nil {
			t.Errorf("Close() returned an error")
		}
	})

	t.Run("PredefinedLevel", func(t *testing.T) {
		w := gosl.NewLvWriter(&buf, gosl.LvInfo) // Set minimum level as LvInfo

		run := func() {
			w.Trace().WriteString("TRACE-1\n") // will NOT be written
			w.Debug().WriteString("DEBUG-1\n") // will NOT be written
			w.Info().WriteString("INFO-1\n")   // will be written
			w.Warn().WriteString("WARN-1\n")   // will be written
			w.Error().WriteString("ERROR-1\n") // will be written
			w.Fatal().WriteString("FATAL-1\n") // will be written
		}

		run()
		if buf.String() != "INFO-1\nWARN-1\nERROR-1\nFATAL-1\n" {
			t.Errorf("Unexpected: <%s>", buf.String())
		}

		buf = buf.Reset()
		w = w.SetLevel(gosl.LvWarn) // Change level to LvWarn

		run()
		if buf.String() != "WARN-1\nERROR-1\nFATAL-1\n" {
			t.Errorf("Unexpected: <%s>", buf.String())
		}
	})

	t.Run("CustomLevel", func(t *testing.T) {
		var (
			INFO uint8 = 0
			WARN uint8 = 3
			ERRR uint8 = 7
			// More...
		)

		w := gosl.NewLvWriter(&buf, WARN) // Writer will print for WARN or above
		buf = buf.Reset()

		w.Lv(INFO).WriteString("INFO-1\n") // NOT Written
		w.Lv(WARN).WriteString("WARN-1\n") // WRITTEN
		w.Lv(ERRR).WriteString("ERRR-1\n") // WRITTEN

		w = w.SetLevel(ERRR)               // change the level to ERRR
		w.Lv(WARN).WriteString("WARN-2\n") // NOT Written
		w.Lv(ERRR).WriteString("ERRR-2\n") // WRITTEN

		if buf.String() != "WARN-1\nERRR-1\nERRR-2\n" {
			t.Errorf("Unexpected: <%s>", buf.String())
		}
	})
}

func BenchmarkLvWriter(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	ss := [][]byte{
		[]byte("000abcdef000abcdef000abcdef000abcdef000abcdef000abcdef000abcdef000abcdef"),
		[]byte("def111defghi111defghi111defghi111defghi111defghi111defghi111defghi111defghi"),
	}
	s := []string{"000abcdef000abcdef000abcdef000abcdef000abcdef000abcdef000abcdef000abcdef", "def111defghi111defghi111defghi111defghi111defghi111defghi111defghi111defghi"}
	_, _ = s, ss

	xs := []string{
		"000 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"001 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"002 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"003 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"004 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
	lw := gosl.LvWriter{}.SetOutput(&buf).SetLevel(0)

	b.Run("WriteString()", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lw = lw.SetOutput(&gosl.Discard{})
			//buf = buf.Reset()
			lw.WriteString(xs[i%3])
		}
	})

	b.Run("WriteAny()", func(b *testing.B) {
		lw = lw.SetLevel(gosl.LvInfo)
		b.Run("enabled", func(b *testing.B) {
			b.ReportAllocs()
			lw = lw.SetOutput(&gosl.Discard{})
			for i := 0; i < b.N; i++ {
				//buf = buf.Reset()
				lw.Fatal().WriteAny(i, ":", xs[i%3])
			}
		})

		b.Run("disabled", func(b *testing.B) {
			b.ReportAllocs()
			lw = lw.SetOutput(&gosl.Discard{})
			for i := 0; i < b.N; i++ {
				//buf = buf.Reset()
				lw.Debug().WriteAny(i, ":", xs[i%3])
			}
		})
	})

	b.Run("WriteAny()+timeFunc", func(b *testing.B) {
		lw = lw.SetLevel(gosl.LvInfo)
		now := time.Now()
		header := func(dst []byte) []byte {
			now = time.Now()
			dst = now.AppendFormat(dst, "01/02 Mon 15:04:05.000 ")
			return dst
		}
		b.Run("enabled", func(b *testing.B) {
			b.ReportAllocs()
			//lw = lw.SetOutput(&buf)
			lw = lw.SetOutput(&gosl.Discard{})
			//lw = lw.SetOutput(os.Stdout)
			for i := 0; i < b.N; i++ {
				//buf = buf.Reset()
				lw.Info().WriteAny(header, ss[i%2])
			}
			//buf.WriteTo(os.Stdout)
		})

		//buf.Println()
		b.Run("disabled", func(b *testing.B) {
			b.ReportAllocs()
			lw = lw.SetOutput(&gosl.Discard{})
			for i := 0; i < b.N; i++ {
				//buf = buf.Reset()
				lw.Debug().WriteAny(header, ss[i%2])
			}
		})
	})
	b.Run("Write(): enabled", func(b *testing.B) {
		b.ReportAllocs()

		l := gosl.NewLvWriter(&buf, gosl.LvInfo)

		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.Info().Write(ss[i%2])
		}
	})

	b.Run("Write(): disabled", func(b *testing.B) {
		b.ReportAllocs()

		l := gosl.NewLvWriter(&buf, gosl.LvFatal)

		for i := 0; i < b.N; i++ {
			buf = buf.Reset()
			l.Info().Write(ss[i%2])
		}
	})
}

// fakeCloserWriter is a struct for faked Write and Close.
type fakeCloserWriter struct {
	out         gosl.Writer
	close       func() error
	closeCalled bool
}

func (w *fakeCloserWriter) Write(p []byte) (n int, err error) {
	if w.out != nil {
		return w.out.Write(p)
	}
	return 0, nil
}
func (w *fakeCloserWriter) Close() error {
	w.closeCalled = true
	if w.close != nil {
		return w.close()
	}
	return nil
}

func Test_Writer_WriterClose(t *testing.T) {
	t.Run("CloseWriter", func(t *testing.T) {
		// Create a Buf that will store fakeCloserWriter
		var trigger = false

		// Create the first fake writer fw1, and add functions
		fw1 := &fakeCloserWriter{
			close: func() error {
				trigger = true
				return nil
			},
		}

		// Close pw, however since pw is regular writer, but its based
		// on an altWriter that has Close() method.
		gosl.Close(fw1)
		gosl.Test(t, true, trigger)
	})

	t.Run("CloseOtherObj", func(t *testing.T) {
		buf := make(gosl.Buf, 0, 1024)

		// make some random type that has Close method
		type fakeSomething interface {
			Close() error
		}
		// create an object that meets the `fakeSomething` interface
		// write "SUCCESS!!" to `Buf` above when `Close()` is called.
		var something fakeSomething = &fakeCloserWriter{
			close: func() error {
				buf = buf.WriteString("SUCCESS!!")
				return nil
			},
		}

		// `gosl.Close()` will trigger `Close()` of `fakeSomething`,
		// therefore writes "SUCCESS!!" to `Buf`
		gosl.Close(something)
		gosl.Test(t, "SUCCESS!!", buf.String())
	})
}
