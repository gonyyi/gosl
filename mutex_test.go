// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"testing"
)

func TestMutex(t *testing.T) {
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

func BenchmarkMutex(b *testing.B) {
	b.Run("mutex", func(b *testing.B) {
		b.ReportAllocs()
		mu := gosl.NewMutex()
		for i := 0; i < b.N; i++ {
			mu.Lock()
			mu.Unlock()
		}
	})
}
