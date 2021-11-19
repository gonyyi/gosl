// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Pool(t *testing.T) {
	type fake struct{ Name string }

	t.Run("reusing test", func(t *testing.T) {
		// Create a pool
		p := gosl.Pool{New: func() interface{} { return &fake{Name: "fresh"} }}
		p = p.Init(3)

		// 1ST ITEM
		item1 := p.Get().(*fake)
		gosl.Test(t, "fresh", item1.Name)
		item1.Name = "gon1"
		gosl.Test(t, "gon1", item1.Name)
		p.Put(item1)

		// 2ND ITEM -- expects the item to have "gon1"
		// as it wasn't reset at the beginning.
		item2 := p.Get().(*fake)
		gosl.Test(t, "gon1", item2.Name)
		p.Put(item2)

	})

	t.Run("concurrent", func(t *testing.T) {
		// Create a pool with max size 3
		p := gosl.Pool{New: func() interface{} { return &fake{Name: "fresh"} }}.Init(3)
		// p.Init(3)

		// 1ST ITEM
		item1 := p.Get().(*fake)
		gosl.Test(t, "fresh", item1.Name)
		item1.Name = "gon1"
		gosl.Test(t, "gon1", item1.Name)
		// this time, after calling the first item, it didn't return.
		// p.Put(item)

		// 2ND ITEM -- since 1st item didn't return the "fake",
		// it will create new fresh
		item2 := p.Get().(*fake)
		gosl.Test(t, "fresh", item2.Name)
		item2.Name = "gon2"

		// this time, after calling the first item, it didn't return.
		// since item1 was returned first and then item 2, next one will
		// have item1's last known value which is "gon1"
		p.Put(item1)
		p.Put(item2)

		// 3RD ITEM
		item3 := p.Get().(*fake)
		gosl.Test(t, "gon1", item3.Name)
		p.Put(item3)

		// 4TH ITEM -- although item3 was returned, it will be behind item2
		// which was returned earlier. Therefore item2's result
		item4 := p.Get().(*fake)
		gosl.Test(t, "gon2", item4.Name)
		p.Put(item4)
	})
}

func Benchmark_Pool(b *testing.B) {

	b.Run("x1", func(b *testing.B) {
		b.ReportAllocs()

		type fake struct{ Name string }
		// p := gosl.NewPool(1, func() interface{} { return &fake{Name: "fresh"} })
		p := gosl.Pool{
			New: func() interface{} { return &fake{Name: "fresh"} },
		}
		p.Init(1)
		for i := 0; i < b.N; i++ {
			item := p.Get()
			p.Put(item)
		}
	})

	b.Run("x256(default)", func(b *testing.B) {
		b.ReportAllocs()

		type fake struct{ Name string }
		p := gosl.Pool{New: func() interface{} { return &fake{Name: "fresh"} }}

		for i := 0; i < b.N; i++ {
			item := p.Get()
			p.Put(item)
		}
	})
}


