// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func Test_Mutex(t *testing.T) {
	mu := gosl.NewMutex()
	for i := 0; i < 100; i++ {
		go mu.Unlock()
		mu.Lock()
	}
	for i := 0; i < 100; i++ {
		go mu.Lock()
		mu.Unlock()
	}
	for i := 0; i < 1000; i++ {
		go mu.Lock()
	}
	for i := 0; i < 1000; i++ {
		mu.Unlock()
	}
}

func Benchmark_Mutex(b *testing.B) {
	b.Run("mutex", func(b *testing.B) {
		b.ReportAllocs()
		mu := gosl.NewMutex()
		for i := 0; i < b.N; i++ {
			mu.Lock()
			mu.Unlock()
		}
	})
	b.Run("mutex.LockFor", func(b *testing.B) {
		b.ReportAllocs()
		mu := gosl.NewMutex()
		for i := 0; i < b.N; i++ {
			mu.LockFor(func() {})
		}
	})
}

func Test_Mutex_Once(t *testing.T) {
	var count = 0
	o := gosl.NewOnce()
	res1 := o.Do(func() { count += 1 })
	res2 := o.Do(func() { count += 1 })

	gosl.TestBool(t, true, res1)  // 1st Once.Do runs, and should return true.
	gosl.TestBool(t, false, res2) // 2nd Once.Do shouldn't run, and it should return false.
	gosl.TestInt(t, 1, count)     // the 2nd function won't run, therefore the result should be 1.
}

func Benchmark_Mutex_Once(b *testing.B) {
	b.Run("Once", func(b *testing.B) {
		b.ReportAllocs()
		once := gosl.NewOnce()
		count := 0
		for i := 0; i < b.N; i++ {
			once.Do(func() {
				count += 1
			})
		}
		// println(count) // this should be 1.
	})
}

