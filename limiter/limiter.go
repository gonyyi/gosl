// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/04/2022

// Limiter is a candidate for GoSL

package limiter

var DefaultLimiterConcurrency = 3

// NewLimiter will create
func NewLimiter(concurrency int) *Limiter {
	l := &Limiter{}
	l.Init(concurrency)
	return l
}

// Limiter is a struct
type Limiter struct {
	limiter      chan struct{}
	count        chan int8
	closeRequest chan struct{}
	closeReady   chan struct{}

	started  int
	finished int
	running  int

	initialized bool
}

// Init will initialize Limiter.
func (l *Limiter) Init(concurrency int) (ok bool) {
	if l.initialized == false {
		l.limiter = make(chan struct{}, concurrency)
		l.count = make(chan int8, concurrency)
		l.closeRequest = make(chan struct{}, 1)
		l.closeReady = make(chan struct{}, 1)
		l.started, l.finished, l.running = 0, 0, 0
		l.monitor()
		l.initialized = true
		return true
	}
	return false
}

// monitor will monitor the stats
func (l *Limiter) monitor() {
	var tmp int8
	go func() {
	loop:
		for {
			select {
			case tmp = <-l.count:
				switch {
				case tmp < 0:
					l.finished += 1
					l.running -= 1
				case tmp > 0:
					l.started += 1
					l.running += 1
				}
			case <-l.closeRequest:
				break loop
			}
		}

		close(l.limiter)
		close(l.count)

		l.closeReady <- struct{}{}
	}()
}

// Started will return a number of tasks started
func (l *Limiter) Started() int { return l.started }

// Finished will return a number of tasks finished
func (l *Limiter) Finished() int { return l.finished }

// Running will return a number of tasks currently running
func (l *Limiter) Running() int { return l.running }

// Wait will wait until the job finishes
func (l *Limiter) Wait() {
	for {
		if l.started == l.finished && l.running == 0 {
			break
		}
	}
}

// Close will wait all the jobs are finished, and then closeReady the channels.
func (l *Limiter) Close() {
	// send a closeReady request signal (closeRequest), and let the monitor to be closed.
	// once monitor receives the closeReady request, it will exit the loop
	// and send closeReady signal back to Close()
	l.closeRequest <- struct{}{}
	<-l.closeReady
	close(l.closeRequest)
	close(l.closeReady)
	l.initialized = false
}

// Ready when waiting for the task to start
func (l *Limiter) Ready() {
	l.count <- 1
	l.limiter <- struct{}{}
}

// Done when task finished
func (l *Limiter) Done() {
	<-l.limiter
	l.count <- -1
}
