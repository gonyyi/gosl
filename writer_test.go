// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

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
		// Create a Buf that will store fakeCloserWriter
		var trigger = false

		// Create the first fake writer fw1, and add functions
		fw1 := &fakeCloserWriter{
			write: func(p []byte) (int, error) {
				return 0, nil
			},
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
