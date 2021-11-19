// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/8/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Mutex_MuInt(t *testing.T) {
	mi := gosl.MuInt{}
	mi = mi.Init()

	mCount := gosl.MuInt{}.Init()

	// outInt1 := 0
	// outInt2 := 0
	for i := 1; i < 101; i++ {
		mCount.Add(1)
		go func(m int) {
			// sum of 1 to 100 should be 5050
			// outInt1 = outInt1 + m
			// outInt2 += m
			mi.Add(m)
			mCount.Add(-1)
		}(i)
	}
	for {
		if mCount.Get() == 0 {
			break
		}
	}
	gosl.Test(t, 5050, mi.Get())
}

func Test_Mutex(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		// mu := make(gosl.Mutex,1)
		var mu gosl.Mutex
		mu = mu.Init()

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
	b.Run("mutex", func(b *testing.B) {
		b.ReportAllocs()
		// mu := gosl.NewMutex()
		var mu gosl.Mutex
		mu = mu.Init()

		for i := 0; i < b.N; i++ {
			mu.Lock()
			mu.Unlock()
		}
	})
	b.Run("mutex.LockFor", func(b *testing.B) {
		b.ReportAllocs()
		var mu gosl.Mutex
		mu = mu.Init()

		for i := 0; i < b.N; i++ {
			mu.LockFor(func() {})
		}
	})
}

func Test_Mutex_Once(t *testing.T) {
	t.Run("3-times", func(t *testing.T) {
		// Create a pool with max size 3
		count := 0
		_ = count

		o := make(gosl.Once, 3)
		// println(100, count, "Left:", o.Available()) // should be 3
		gosl.Test(t, 0, count)
		gosl.Test(t, 3, o.Available())

		for i := 0; i < 10; i++ {
			o.Do(func() {
				count += 1
			})
		}
		// println(110, count, "Left:", o.Available()) // should be 0
		gosl.Test(t, 3, count)
		gosl.Test(t, 0, o.Available())

		o.Reset()
		// println(120, count, "Left:", o.Available()) // should be 3
		gosl.Test(t, 3, count)
		gosl.Test(t, 3, o.Available())

		o.Do(func() { count += 1 })
		// println(130, count, "Left:", o.Available()) // should be 2
		gosl.Test(t, 4, count)
		gosl.Test(t, 2, o.Available())

		o.Reset()
		// println(131, count, "Left:", o.Available()) // should be 3
		gosl.Test(t, 4, count)
		gosl.Test(t, 3, o.Available())

		o.Do(func() { count += 1 })
		// println(131, count, "Left:", o.Available()) // should be 2
		gosl.Test(t, 5, count)
		gosl.Test(t, 2, o.Available())

		for i := 0; i < 10; i++ {
			o.Do(func() {
				count += 1
			})
		}
		// println(140, count, "Left:", o.Available()) // should be 0
		gosl.Test(t, 7, count)
		gosl.Test(t, 0, o.Available())
	})

	t.Run("basic", func(t *testing.T) {
		var count = 0
		o := make(gosl.Once, 1)
		res1 := o.Do(func() { count += 1 })
		res2 := o.Do(func() { count += 1 })

		gosl.Test(t, true, res1)  // 1st Once.Do runs, and should return true.
		gosl.Test(t, false, res2) // 2nd Once.Do shouldn't run, and it should return false.
		gosl.Test(t, 1, count)     // the 2nd function won't run, therefore the result should be 1.
	})
}

func Benchmark_Mutex_Once(b *testing.B) {
	b.Run("Once", func(b *testing.B) {
		b.ReportAllocs()
		once := make(gosl.Once, 1)
		count := 0
		for i := 0; i < b.N; i++ {
			once.Do(func() {
				count += 1
			})
		}
		// println(count) // this should be 1.
	})
}


