// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/03/2022

package gosl_test

import (
	"testing"

	"github.com/gonyyi/gosl"
)

func Test_Mutex_MuInt(t *testing.T) {
	var mi = gosl.NewMuInt()
	var mCount = gosl.NewMuInt()

	mCount.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(m int) {
			// sum of 1 to 1000 should be 500500
			mi.Add(m)      // final value goes here
			mCount.Add(-1) // to make sure all goroutines are ran
		}(i + 1)
	}
	mCount.Wait(0)
	gosl.Test(t, 500500, mi.Get())
}

func Test_Mutex(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var mu gosl.Mutex = gosl.NewMutex()

		gosl.Test(t, false, mu.Locked())
		mu.Lock()
		gosl.Test(t, true, mu.Locked())
		mu.Unlock()
		gosl.Test(t, false, mu.Locked())

		mu.LockFor(func() {
			gosl.Test(t, true, mu.Locked())
		})
		gosl.Test(t, false, mu.Locked())

		func() {
			mu.Lock()
			defer mu.Unlock()
			gosl.Test(t, true, mu.Locked())
		}()
		gosl.Test(t, false, mu.Locked())
	})

}

func Benchmark_Mutex(b *testing.B) {
	b.Run("Mutex", func(b *testing.B) {
		var count int

		b.Run("LockUnlock", func(b *testing.B) {
			b.ReportAllocs()
			mu := gosl.NewMutex()
			count = 0
			for i := 0; i < b.N; i++ {
				mu.Lock()
				count += 1
				mu.Unlock()
			}
			// println(b.N, count)
		})
		b.Run("LockFor", func(b *testing.B) {
			b.ReportAllocs()
			mu := gosl.NewMutex()
			count = 0
			for i := 0; i < b.N; i++ {
				mu.LockFor(func() {
					count += 1
				})
			}
			// println(b.N, count)
		})
		b.Run("LockIfNot", func(b *testing.B) {
			b.ReportAllocs()
			mu := gosl.NewMutex()
			count = 0
			for i := 0; i < b.N; i++ {
				if mu.LockIfNot() {
					count += 1
					mu.Unlock()
				}
			}
			// println(b.N, count)
		})
	})

	b.Run("MuInt", func(b *testing.B) {
		b.ReportAllocs()
		mi := gosl.NewMuInt()

		mi.Set(b.N)
		for i := 0; i < b.N; i++ {
			mi.Add(-1)
		}
	})
}


func Test_Pool(t *testing.T) {
	type fake struct{ Name string }

	t.Run("reusing test", func(t *testing.T) {
		// Create a pool
		newItems := 0
		p := gosl.Pool{New: func() interface{} {
			newItems += 1
			return &fake{Name: "fresh"}
		}}
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

		gosl.Test(t, 1, newItems)
	})

	t.Run("reusing+new", func(t *testing.T) {
		// Create a pool
		newItems := 0
		p := gosl.Pool{New: func() interface{} {
			newItems += 1
			return &fake{Name: "fresh"}
		}}
		p = p.Init(3)

		// 1ST ITEM
		item1 := p.Get().(*fake)
		gosl.Test(t, "fresh", item1.Name)
		item1.Name = "gon1"
		gosl.Test(t, "gon1", item1.Name)

		// 2ND ITEM -- expects the item to have "gon1"
		// as it wasn't reset at the beginning.
		item2 := p.Get().(*fake)
		item3 := p.Get().(*fake)
		item4 := p.Get().(*fake)
		item5 := p.Get().(*fake)

		p.Put(item1)
		p.Put(item2)
		p.Put(item3)
		p.Put(item4)
		p.Put(item5)

		gosl.Test(t, 5, newItems)
	})

	t.Run("missingNewFunc", func(t *testing.T) {
		// Create a pool
		p := gosl.Pool{}
		p = p.Init(2)

		// 1ST ITEM
		item1 := p.Get()

		gosl.Test(t, true, item1 == nil) // this is to check if p.Get() won't panic
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
		p = p.Init(1)
		for i := 0; i < b.N; i++ {
			item := p.Get()
			p.Put(item)
		}
	})

	b.Run("x256", func(b *testing.B) {
		b.ReportAllocs()

		type fake struct{ Name string }
		p := gosl.Pool{New: func() interface{} { return &fake{Name: "fresh"} }}.Init(256)

		for i := 0; i < b.N; i++ {
			item := p.Get()
			// println(item.(*fake).Name)
			p.Put(item)
		}
	})
}