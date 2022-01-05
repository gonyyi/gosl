// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/04/2022

package limiter_test

import (
	"testing"

	"github.com/gonyyi/gosl/limiter"
)

func TestLimiter(t *testing.T) {
	// Create a limit with 10 jobs at a time
	// > limit := limiter.Limiter{}
	// > limit.Init(12)
	// Or simply,
	// > limit := NewLimiter(10)

	limit := limiter.Limiter{}
  
	run := func(z int) {
		// println("\nSTEP ",z," -------------")
		limit.Init(5)

		for i := 0; i < 10000; i++ {
			limit.Ready()
			go func() {
				limit.Done()
			}()
		}

		limit.Wait()
		limit.Close()
		// println("\tDone: ", limit.Started(), "/", limit.Finished(), "/", limit.Running())
		if limit.Started() != 10000 || limit.Finished() != 10000 || limit.Running() != 0 {
			t.Errorf("Limiter Test %d - Started: %d, Finished: %d, Running: %d", z, limit.Started(), limit.Finished(), limit.Running())
			t.Fail()
		}
	}

	for i:=0; i<10; i++ {
		run(i)
	}
}
