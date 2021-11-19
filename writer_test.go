// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"bytes"
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Writer_MultiWriter(t *testing.T) {
	buf1 := make(gosl.Buf, 0, 1024)
	buf2 := make(gosl.Buf, 0, 1024)
	buf3 := make(gosl.Buf, 0, 1024)
	buf4 := make(gosl.Buf, 0, 1024)
	buf5 := make(gosl.Buf, 0, 1024)

	t.Run("2w", func(t *testing.T) {
		dw := gosl.NewMultiWriter(&buf1, &buf2)
		dw.Write([]byte("hello gon1"))
		gosl.Test(t, "hello gon1", buf1.String())
		gosl.Test(t, "hello gon1", buf2.String())
		gosl.Test(t, "", buf3.String())
		gosl.Test(t, "", buf4.String())
		gosl.Test(t, "", buf5.String())
	})

	buf1 = buf1.Reset()
	buf2 = buf2.Reset()
	buf3 = buf3.Reset()
	buf4 = buf4.Reset()
	buf5 = buf5.Reset()

	t.Run("5w", func(t *testing.T) {
		dw := gosl.NewMultiWriter(&buf1, &buf2, &buf3, &buf4, &buf5)
		dw.Write([]byte("hello gon5"))
		gosl.Test(t, "hello gon5", buf1.String())
		gosl.Test(t, "hello gon5", buf2.String())
		gosl.Test(t, "hello gon5", buf3.String())
		gosl.Test(t, "hello gon5", buf4.String())
		gosl.Test(t, "hello gon5", buf5.String())
	})
}

func Test_Writer_AlterWriter(t *testing.T) {
	var buf bytes.Buffer
	dw := gosl.NewAltWriter(&buf, func(p []byte) []byte {
		if bytes.Contains(p, []byte("hello")) {
			return nil
		}
		return p
	})

	dw.Write([]byte("hello gon"))
	dw.Write([]byte("hi gon"))
	gosl.Test(t, "hi gon", buf.String())
}

// fakeCloserWriter is a struct for faked Write and Close.
type fakeCloserWriter struct {
	out   gosl.Writer
	write func(p []byte) (n int, err error)
	close func() error
}

func (w *fakeCloserWriter) Write(p []byte) (n int, err error) {
	return w.write(p)
}
func (w *fakeCloserWriter) Close() error {
	return w.close()
}

func Test_Writer_WriterClose(t *testing.T) {

	t.Run("CloseWriter", func(t *testing.T) {
		// Create a Buffer that will store fakeCloserWriter
		buf := make(gosl.Buf, 0, 1024)

		// Create the first fake writer fw1, and add functions
		fw1 := &fakeCloserWriter{
			write: func(p []byte) (int, error) { return buf.Write(p) },
			close: func() error { buf = buf.WriteString("FW1:CLOSING-TIME"); return nil },
		}

		// Create prefix writers chained as below:
		// pw3 --(writes)--> pw2 --(writes)--> pw1 --(writes)--> fw1
		// The prefix writer doesn't implemented close method, however,
		// it's original form is altWriter which has close method.
		// Therefore, gosl.Close() will run Close() of altWriter,
		// and on and on, eventually to fw1.
		pw1 := gosl.NewPrefixWriter("PW1:", fw1)
		pw2 := gosl.NewPrefixWriter("PW2:", pw1)
		pw3 := gosl.NewPrefixWriter("PW3:", pw2)

		// Write something to pw
		pw3.Write([]byte("hello gon!\n"))
		pw3.Write([]byte("how are you!\n"))

		// Close pw, however since pw is regular writer, but its based
		// on an altWriter that has Close() method.
		gosl.Close(pw2)

		gosl.Test(t, "PW1:PW2:PW3:hello gon!\nPW1:PW2:PW3:how are you!\nFW1:CLOSING-TIME", buf.String())
	})

	t.Run("CloseOtherObj", func(t *testing.T) {
		buf := make(gosl.Buf, 0, 1024)

		// make some random type that has Close method
		type fakeSomething interface {
			Close() error
		}
		// create an object that meets the `fakeSomething` interface
		// write "SUCCESS!!" to `Buffer` above when `Close()` is called.
		var something fakeSomething = &fakeCloserWriter{
			close: func() error {
				buf = buf.WriteString("SUCCESS!!")
				return nil
			},
		}

		// `gosl.Close()` will trigger `Close()` of `fakeSomething`,
		// therefore writes "SUCCESS!!" to `Buffer`
		gosl.Close(something)
		gosl.Test(t, "SUCCESS!!", buf.String())
	})
}


