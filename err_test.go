// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 12/13/2021

package gosl_test

import (
	"errors"
	"fmt"
	"github.com/gonyyi/gosl"
	"testing"
)

func TestIsError(t *testing.T) {
	e1 := errors.New("e1")
	e21 := fmt.Errorf("m2: %w", e1)
	e321 := gosl.WrapError("m3", e21)
	e4321 := fmt.Errorf("m4: %w", e321)
	f := gosl.NewError("e1")

	gosl.Test(t, true, gosl.IsError(e4321, e1))
	gosl.Test(t, true, gosl.IsError(e4321, e21))
	gosl.Test(t, true, gosl.IsError(e4321, e321))
	gosl.Test(t, true, gosl.IsError(e4321, e4321))
	gosl.Test(t, false, gosl.IsError(e4321, f))
	gosl.Test(t, false, gosl.IsError(e1, f))
}

func TestUnwrapError(t *testing.T) {
	e1 := errors.New("e1")
	e21 := fmt.Errorf("m2: %w", e1)
	e321 := gosl.WrapError("m3", e21)
	e4321 := fmt.Errorf("m4: %w", e321)

	gosl.Test(t, "m4: m3: m2: e1", e4321.Error())

	u321 := gosl.UnwrapError(e4321)
	gosl.Test(t, "m3: m2: e1", u321.Error())

	u21 := gosl.UnwrapError(u321)
	gosl.Test(t, "m2: e1", u21.Error())

	u1 := gosl.UnwrapError(u21)
	gosl.Test(t, "e1", u1.Error())

	x321 := errors.Unwrap(e4321)
	gosl.Test(t, "m3: m2: e1", x321.Error())

	x21 := errors.Unwrap(e321)
	gosl.Test(t, "m2: e1", x21.Error())

	x1 := errors.Unwrap(e21)
	gosl.Test(t, "e1", x1.Error())
	gosl.Test(t, e1, x1)
}

func TestWrapError(t *testing.T) {
	e1 := gosl.NewError("e1")
	e2 := errors.New("e2")
	e3 := fmt.Errorf("e3")
	_, _, _ = e1, e2, e3

	gosl.Test(t, e1.Error(), "e1")
	gosl.Test(t, e2.Error(), "e2")
	gosl.Test(t, e3.Error(), "e3")

	// When wrap, it will join

	e12 := gosl.WrapError(e1.Error(), e2) // e1:e2
	e21 := gosl.WrapError(e2.Error(), e1)
	gosl.Test(t, e12.Error(), "e1: e2")
	gosl.Test(t, e21.Error(), "e2: e1")
	gosl.Test(t, false, e12 == e21)

	{
		e12_m1_a := gosl.UnwrapError(e12)
		e12_m1_b := errors.Unwrap(e12)
		gosl.Test(t, true, e12_m1_a == e2)
		gosl.Test(t, true, e12_m1_b == e2)
		gosl.Test(t, true, e12_m1_a.Error() == e2.Error())
		gosl.Test(t, true, e12_m1_b.Error() == e2.Error())

		e21_m2_a := gosl.UnwrapError(e21)
		e21_m2_b := errors.Unwrap(e21)
		gosl.Test(t, true, e21_m2_a == e1)
		gosl.Test(t, true, e21_m2_b == e1)
		gosl.Test(t, true, e21_m2_a.Error() == e1.Error())
		gosl.Test(t, true, e21_m2_b.Error() == e1.Error())
	}
}

func TestNewError(t *testing.T) {
	err1 := gosl.NewError("") // this should be nil
	err2 := gosl.NewError("some error")
	gosl.Test(t, true, err1 == nil)
	gosl.Test(t, false, err1 == err2)
	gosl.Test(t, false, err2 == nil)
}

func Test_Err_IfPanic(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)

	t.Run("panicMsg=error", func(t *testing.T) {
		f := func() {
			defer gosl.IfPanic("crazy", func(e interface{}) {
				buf = buf.WriteString("Func1: Panic -> ")
				if err, ok := e.(error); ok {
					buf = buf.WriteString(err.Error())
				}
			})
			panic(gosl.NewError("SomeError"))
		}
		f()
		gosl.Test(t, "Func1: Panic -> SomeError", buf.String())
	})

	t.Run("panicMsg=string", func(t *testing.T) {
		buf = buf.Reset()
		f := func() {
			defer gosl.IfPanic("crazy", func(e interface{}) {
				buf = buf.WriteString("Func1: Panic -> ")
				if err, ok := e.(error); ok {
					buf = buf.WriteString(err.Error())
				} else if s, ok := e.(string); ok {
					buf = buf.WriteString(s)
				}
			})
			panic("Something")
		}
		f()
		gosl.Test(t, "Func1: Panic -> Something", buf.String())
	})

	t.Run("panicMsg=stringer", func(t *testing.T) {
		buf = buf.Reset()
		f := func() {
			defer gosl.IfPanic("crazy", func(e interface{}) {
				buf = buf.WriteString("Func1: Panic -> ")
				if err, ok := e.(error); ok {
					buf = buf.WriteString(err.Error())
				} else if err, ok := e.(string); ok {
					buf = buf.WriteString(err)
				} else if err, ok := e.(interface{String()string}); ok {
					buf = buf.WriteString(err.String())
				}
			})

			// Use gosl.Buffer as it's also a stringer
			s := make(gosl.Buf, 10)
			s = s.Reset().WriteString("stringer")
			panic(s)
		}
		f()
		gosl.Test(t, "Func1: Panic -> stringer", buf.String())
	})

	t.Run("panicMsg=unsupp", func(t *testing.T) {
		buf = buf.Reset()
		f := func() {
			defer gosl.IfPanic("crazy", func(e interface{}) {
				buf = buf.WriteString("Func1: Panic -> ")
				if err, ok := e.(error); ok {
					buf = buf.WriteString(err.Error())
				} else if err, ok := e.(string); ok {
					buf = buf.WriteString(err)
				} else {
					buf = buf.WriteString("unsupported panic info")
				}
			})
			panic(123)
		}
		f()
		gosl.Test(t, "Func1: Panic -> unsupported panic info", buf.String())
	})

	t.Run("funcNotGiven", func(t *testing.T) {
		f := func() {
			defer gosl.IfPanic("crazy", nil)
			panic(123)
		}
		f()
	})
}

func BenchmarkIfPanic(b *testing.B) {
	b.Run("basic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			gosl.IfPanic("test", nil) // 2.6
		}
	})
}

func TestIfErr(t *testing.T) {
	var testErr = gosl.NewError("(errCode1) test error")
	_ = testErr

	gosl.IfErr("err1", nil)
	gosl.IfErr("err2", testErr)
}

func BenchmarkIfErr(b *testing.B) {
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
