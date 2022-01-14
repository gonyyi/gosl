// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/04/2022

package limiter_test

import (
	"github.com/gonyyi/gosl/limiter"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	l := limiter.NewLimiter(3, 5) // Create Limiter with 10 concurrent workers and 80 queues
	for i := 0; i < 20; i++ {
		func(i int) {
			l.Run(func() { // Add a job using *Limiter.Run(fn)
				println("START", i)
				time.Sleep(time.Second)
				println("END", i)
			})
		}(i)
	}

	l.Stop(true)

	for l.IsActive() { // Wait until it's completely stopped
		time.Sleep(time.Second)
		s, w, q := l.Status()
		println("Looking: ", s, w, q)
	}

	if ok := l.Close(); !ok {
		s, w, q := l.Status()
		println("Unexpected: ", s, w, q)
	}
	println("FINISHED")
}
