// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Err_IfPanic(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("panicMsg=error", func(t *testing.T) {
		f := func() {
			defer gosl.IfPanic("crazy", func(e error) {
				buf = buf.WriteString("Func1:Panic -> ")
				buf = buf.WriteString(e.Error())
			})
			panic(gosl.NewError("SomeError"))
		}
		f()
		gosl.Test(t, "Func1:Panic -> SomeError", buf.String())
	})

	t.Run("panicMsg=string", func(t *testing.T) {
		buf = buf.Reset()
		f := func() {
			defer gosl.IfPanic("crazy", func(e error) {
				buf = buf.WriteString("Func1:Panic -> ")
				buf = buf.WriteString(e.Error())
			})
			panic("Something")
		}
		f()
		gosl.Test(t, "Func1:Panic -> Something", buf.String())
	})

	t.Run("panicMsg=stringer", func(t *testing.T) {
		buf = buf.Reset()
		f := func() {
			defer gosl.IfPanic("crazy", func(e error) {
				buf = buf.WriteString("Func1:Panic -> ")
				buf = buf.WriteString(e.Error())
			})

			// Use gosl.Buffer as it's also a stringer
			s := make(gosl.Buf, 10)
			s = s.Reset().WriteString("stringer")
			panic(s)
		}
		f()
		gosl.Test(t, "Func1:Panic -> stringer", buf.String())
	})

	t.Run("panicMsg=unsupp", func(t *testing.T) {
		buf = buf.Reset()
		f := func() {
			defer gosl.IfPanic("crazy", func(e error) {
				buf = buf.WriteString("Func1:Panic -> ")
				buf = buf.WriteString(e.Error())
			})
			panic(123)
		}
		f()
		gosl.Test(t, "Func1:Panic -> unsupported panic info", buf.String())
	})

}

func Benchmark_Err_IfPanic(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			gosl.IfPanic("test", nil) // 2.6
		}
	})
}

func Test_Err_IfErr(t *testing.T) {
	var testErr = gosl.NewError("(errCode1) test error")
	_ = testErr

	gosl.IfErr("err1", nil)
	// gosl.IfErr("err2", testErr)
}

func Benchmark_Err_IfErr(b *testing.B) {
	errs := []error{
		gosl.NewError("err1"),
		gosl.NewError("err2"),
		gosl.NewError("err3"),
	}
	errs[1] = nil // set err2 to be nil error.

	// At this test, 2 out of 3 will be printed as errs[1] is nil..
	b.Run("err=2,nil=1", func(b *testing.B) {
		// Skip as this will print to os.Stdout
		b.SkipNow()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// Benchmark_err_IfErr/nil-8         	  865792	      1384 ns/op	       0 B/op	       0 allocs/op
			gosl.IfErr("errTest", errs[i%3])
		}
	})

	errs[0] = nil
	errs[2] = nil

	// At this test, none will have printed as all the errors are nil.
	b.Run("nil=3", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// Benchmark_err_IfErr/nil-8         	  865792	      1384 ns/op	       0 B/op	       0 allocs/op
			gosl.IfErr("errTest", errs[i%3])
		}
	})
}


