// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/20/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

func TestLineWriter(t *testing.T) {
	w1 := make(gosl.Buf, 0, 1024)
	w2 := gosl.Discard
	
	t.Run("Close()", func(t *testing.T) {
		w1 = w1.Reset()
		var fw = fakeCloserWriter{out: &w1}
		var lw = gosl.NewLineWriter(&fw)
		//var lw = gosl.LineWriter{}.SetOutput(&fw)
		_ = lw
		lw.WriteString("abc") // this writes to buf w1

		println(fw.closeCalled)
		if err := lw.Close(); err != nil {
			println(err.Error())
		}
		println(fw.closeCalled)

	})
	
	t.Run("SetOutput(),Output(),Init()", func(t *testing.T) {
		t.Run("split:sameWriter", func(t *testing.T) {
			var lw0 = gosl.LineWriter{}.Init()
			lw1 := lw0.SetOutput(&w1)
			lw2 := lw0.SetOutput(w2)
			gosl.Test(t, true, lw1.Output() == lw2.Output())
		})
		t.Run("split:diffWriter", func(t *testing.T) {

			var lw0 = gosl.LineWriter{}
			lw1 := lw0.SetOutput(&w1)
			lw2 := lw0.SetOutput(w2)
			gosl.Test(t, false, lw1.Output() == lw2.Output())
		})
	})

	t.Run("Enable(),Enabled(),WriteString(),Write()", func(t *testing.T) {
		w1 = w1.Reset()
		var lw0 = gosl.LineWriter{}.SetOutput(&w1)
		lw0.WriteString("t1")
		gosl.Test(t, "t1\n", w1.String()) // LOGGED
		gosl.Test(t, true, lw0.Enabled())
		lw0 = lw0.Enable(false)
		gosl.Test(t, false, lw0.Enabled())
		if n, err := lw0.WriteString("t2"); n != 0 || err != nil { // THIS WON'T LOGGED
			t.Fail()
		}
		gosl.Test(t, "t1\n", w1.String()) // above didn't add because wasn't enabled

		lw0 = lw0.Enable(true)
		if n, err := lw0.WriteString("t3"); n != 2 || err != nil { // LOGGED
			t.Fail()
		}
		gosl.Test(t, true, lw0.Enabled())

		if n, err := lw0.Write([]byte("t4\n")); n != 3 || err != nil { // LOGGED
			t.Fail()
		}
		if n, err := lw0.Write([]byte("t5")); n != 2 || err != nil { // LOGGED
			t.Fail()
		}
		gosl.Test(t, "t1\nt3\nt4\nt5\n", w1.String())
		w1 = w1.Reset()

		lw0 = lw0.SetOutput(gosl.Discard)
		if n, err := lw0.Write([]byte("123")); n != 3 || err != nil {
			t.Fail()
		}
		if n, err := lw0.Write([]byte("abc\n")); n != 4 || err != nil {
			t.Fail()
		}
		if n, err := lw0.WriteString("234"); n != 3 || err != nil {
			t.Fail()
		}
		if n, err := lw0.WriteString("bcd\n"); n != 4 || err != nil {
			t.Fail()
		}

		t.Run("Enable()", func(t *testing.T) {
			// When disabled, if it returns the counter
			lw0 = lw0.Enable(false)
			if n, err := lw0.WriteString("bcd\n"); n != 0 || err != nil {
				t.Fail()
			}
			if n, err := lw0.Write([]byte("bcd\n")); n != 0 || err != nil {
				t.Fail()
			}

			lw0 = lw0.Enable(true)
			gosl.Test(t, true, lw0.Enabled())
			lw0 = lw0.SetOutput(nil) // this should enable it.
			gosl.Test(t, false, lw0.Enabled())
			lw0 = lw0.Enable(true)
			gosl.Test(t, false, lw0.Enabled())
		})

	})
	t.Run("NewLineWriter()", func(t *testing.T) {
		w := gosl.NewLineWriter(gosl.Discard)
		gosl.Test(t, true, w.Enabled())
	})
	t.Run("Close()", func(t *testing.T) {
		gosl.Test(t, nil, gosl.Close(123))
	})
}

func BenchmarkLineWriter(b *testing.B) {
	var ssNoNL, ssHasNL []string // 1s don't have newline, 2s have newline
	var bssNoNL, bssHasNL [][]byte

	ssNoNL = []string{
		"000 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"001 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"002 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"003 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"004 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
	for _, v := range ssNoNL { // Create ssHasNL with newline
		ssHasNL = append(ssHasNL, v+"\n")
	}
	for _, v := range ssNoNL { // Create bssNoNL without newline
		bssNoNL = append(bssNoNL, []byte(v))
	}
	for _, v := range ssHasNL { // Create bssHasNL with newline
		bssHasNL = append(bssHasNL, []byte(v))
	}

	lw := gosl.LineWriter{}.SetOutput(gosl.Discard)
	if lw.Enabled() == false {
		b.Fail()
	}

	b.Run("Write():NL", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lw.Write(bssHasNL[i%5])
		}
	})
	b.Run("Write():NoNL", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lw.Write(bssNoNL[i%5])
		}
	})

	b.Run("WriteString():NL", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lw.WriteString(ssHasNL[i%5])
		}
	})
	b.Run("WriteString():NoNL", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			lw.WriteString(ssNoNL[i%5])
		}
	})

	b.Run("Enabled(false)", func(b *testing.B) {
		lw = lw.Enable(false)
		b.Run("Write():NL", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				lw.Write(bssHasNL[i%5])
			}
		})
		b.Run("WriteString():NL", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				lw.WriteString(ssHasNL[i%5])
			}
		})
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
