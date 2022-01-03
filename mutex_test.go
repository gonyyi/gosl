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
