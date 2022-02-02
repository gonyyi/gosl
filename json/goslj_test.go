package goslj_test

import (
	"github.com/gonyyi/gosl"
	goslj "github.com/gonyyi/gosl/json"
	"os"
	"testing"
)

var testIn = `gon is	happy here	한글이름: 이건용` +
	"\n" +
	`so no need to 		worry {} () .: \ !@#$%^&*()_+{}[];':",<.>/?"'`
var testOut = `{"name":"gon is\thappy here\t한글이름: 이건용\nso no need to \t\tworry {} () .: \\ !@#$%^&*()_+{}[];':\",<.>/?\"'"}`

func ExampleJSON_Main() {
	jp := goslj.NewPool(20) // create a JSON pool with 20 objects

	j1 := jp.Get() // get JSON from the pool
	j2 := jp.Get()
	j3 := jp.Get()

	j1.Start(). // JSON always starts with `Start()`, and ends with `End()`
			String("city", "conway").
			String("state", "arkansas").
			Int("zip", 72034).
			End() // end
	j2.Start().
		String("name", "gonn corp").
		Int("tin", 123456789).
		Int("income", 123456).
		End()
	j3.Start().
		String("name", "Gon Yi").
		Int("age", 100).
		Sub("address", j1).  // add j1 into j
		Sub("employer", j2). // add j2 into j
		End().
		Write(os.Stdout) // print to screen
	j1.Putback() // return to pool
	j2.Putback()
	j3.Putback()

	// Output:
	// {"name":"Gon Yi","age":100,"address":{"city":"conway","state":"arkansas","zip":72034},"employer":{"name":"gonn corp","tin":123456789,"income":123456}}
}

func ExampleJSON() {
	goslj.NewJSON(1024).Start().
		String("name", testIn).
		End().Write(os.Stdout)
	// Output:
	// {"name":"gon is\thappy here\t한글이름: 이건용\nso no need to \t\tworry {} () .: \\ !@#$%^&*()_+{}[];':\",<.>/?\"'"}
}

func TestJSON(t *testing.T) {
	buf := make(gosl.Buf, 0, 1024)
	j := goslj.NewJSON(1024)
	j.Start().String("name", testIn).End().Write(&buf)
	gosl.Test(t, testOut, buf.String())
}

func BenchmarkJSON(b *testing.B) {
	buf := make(gosl.Buf, 0, 1024)
	j := goslj.NewJSON(1024)
	jp := goslj.NewPool(20)
	discard := &gosl.Discard{}
	_, _ = buf, discard

	b.Run("simple", func(b *testing.B) {
		// BenchmarkJSON/simple-12         	21592483	        51.36 ns/op	       0 B/op	       0 allocs/op
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			//buf = buf.Reset()
			j.Reset().Start().String("name", "Gon Yi").Int("age", 100).End().Write(discard)
		}
	})

	b.Run("complex", func(b *testing.B) {
		// BenchmarkJSON/complex-12         	 4422925	       250.2 ns/op	       0 B/op	       0 allocs/op
		b.ReportAllocs()
		j1 := jp.Get() // since this isn't testing for pool, avoid get/put during for-loop iteration
		j2 := jp.Get()
		for i := 0; i < b.N; i++ {
			//buf = buf.Reset()
			j1.Reset().Start().
				String("city", "conway").
				String("state", "arkansas").
				Int("zip", 72034).
				End()
			j2.Reset().Start().
				String("name", "gonn corp").
				Int("tin", 123456789).
				Int("income", 123456).
				End()
			j.Reset().Start().
				String("name", "Gon Yi").
				Int("age", 100).
				Sub("address", j1).  // add j1 into j
				Sub("employer", j2). // add j2 into j
				End().Write(discard)
		}
		j1.Putback()
		j2.Putback()
		// {
		//    "name": "Gon Yi",
		//    "age": 100,
		//    "address": {
		//        "city": "conway",
		//        "state": "arkansas",
		//        "zip": 72034
		//    },
		//    "employer": {
		//        "name": "gonn corp",
		//        "tin": 123456789,
		//        "income": 123456
		//    }
		// }
	})

	b.Run("pool+simple", func(b *testing.B) {
		// BenchmarkJSON/simple+pool-12         	11616590	        93.72 ns/op	       0 B/op	       0 allocs/op
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			//buf = buf.Reset()
			jp.Get().Start().String("name", "Gon Yi").Int("age", 100).End().Write(discard).Putback()
		}
	})

	b.Run("pool+complex", func(b *testing.B) {
		// BenchmarkJSON/complex+pool-12         	 2959813	       390.1 ns/op	       0 B/op	       0 allocs/op
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			//buf = buf.Reset()
			j1 := jp.Get().Start().
				String("city", "conway").
				String("state", "arkansas").
				Int("zip", 72034).
				End()
			j2 := jp.Get().Start().
				String("name", "gonn corp").
				Int("tin", 123456789).
				Int("income", 123456).
				End()
			j0 := jp.Get().Start().
				String("name", "Gon Yi").
				Int("age", 100).
				Sub("address", j1).
				Sub("employer", j2).
				End().Write(discard)

			// return to pool
			//_ = j0
			j0.Putback()
			j1.Putback()
			j2.Putback()
		}
	})
	//println(jp.Stats())
	//buf.Println()
}
