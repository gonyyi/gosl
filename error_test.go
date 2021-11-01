package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

var testErr = gosl.NewError("(errCode1) test error")

func Test_Err_IfErr(t *testing.T) {
	gosl.IfErr("err1", nil)
	gosl.IfErr("err2", testErr)
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
