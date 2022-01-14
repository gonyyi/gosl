package limiter

// (c) Gon Y. Yi 2021-2022 <https://gonyyi.com/copyright>
// Last Update: 01/13/2022

// CONCURRENCY LIMITER v1.0.0
// --------------------------
// Usage:
//     l := limiter.NewLimiter(10, 80) // Create Limiter with 10 concurrent workers and 80 queues
//     for someCondition {
//         l.Run(func(){  // Add a job using *Limiter.Run(fn)
//             someCode() // Code that needs to run concurrently
//         })
//     }
//     l.Stop(false) // Stop the limiter. Anything after this will be skipped.
//     for l.IsActive() == false { // Wait until it's completely stopped
//         time.Sleep(time.Second))
//     }
//     l.Close() // Close the Limiter
//

// NewLimiter will return a *Limiter
func NewLimiter(worker, queue uint16) *Limiter {
	l := &Limiter{}
	if worker == 0 {
		worker = 1
	}
	if queue == 0 {
		queue = 1
	}
	l, _ = l.Init(worker, queue)
	return l
}

// Limiter is a queue based concurrent runner.
type Limiter struct {
	worker chan struct{} // worker limits how many concurrent
	queue  chan func()   // queue for jobs
	mu     chan struct{} // mutex
	status bool          // only when status is true, new job can be added to queue.
}

// Init initialize Limiter
// If queue size, job may hold until enough queue is available.
func (l *Limiter) Init(workers, queue uint16) (lim *Limiter, ok bool) {
	// do not allow unbuffered channel.
	if workers == 0 || queue == 0 {
		return nil, false
	}

	// l.mu will be nil if the Limiter (1) is a fresh object, or (2) has called Closed()
	if l.mu != nil {
		return l, false // when Closed() is not called, *Limiter.monitor() maybe still running.
	}

	l.worker = make(chan struct{}, workers)
	l.queue = make(chan func(), queue)
	l.mu = make(chan struct{}, 1) // mutex, let only 1 at a time
	l.status = true               // false -> true
	go l.monitor()                // start monitoring in background. this will be cancelled only when Close() is called.
	return l, true
}

// ifStatusIs will compare status with *Limiter.status. If same, this will return true.
// Also, if status and new values are differ, it will switch status value to new value.
func (l *Limiter) ifStatusIs(status, new bool) (match bool) {
	l.mu <- struct{}{}         // lock
	match = l.status == status // match (output) will be true if a func argument status is same as l.status.
	if match && status != new {
		l.status = new // when given status and new are different, update status AFTER calculate the match (above)
	}
	<-l.mu // unlock
	return
}

// monitor monitors queue, and when new queue arrives, it will hand it to worker if available.
// otherwise, it will wait a worker become available.
// monitor will stop when *Limiter.queue is closed.
func (l *Limiter) monitor() {
loop:
	for {
		select {
		case f, ok := <-l.queue: // as soon as one gets out, new one can be there. and until ifStatusIs, one can be there as well.
			if ok {
				l.worker <- struct{}{} // WAIT until worker is available ** THIS CAN HIDE 1 JOB THAT GOES TO WORKER
				go func() {
					f()
					<-l.worker // NO ISSUE even when empty and closed EXCEPT WHEN cap(worker)=0
				}()
			} else { // if queue is closed, it will exit the monitor
				break loop
			}
		}
	}
}

// Run adds a job func() to queue. If limiter is no longer accepting, it will return false.
func (l *Limiter) Run(f func()) (ok bool) {
	// don't let nil to get in as a func
	if f == nil {
		return false
	}

	// Add new jobs to the queue if status says its available
	if l.ifStatusIs(true, true) {
		l.queue <- f
		return true
	}

	return ok
}

// Stop will make limiter stop taking jobs.
// - allow == false: All jobs in the queue will be cancelled and return how many were cancelled.
// - allow == true:  this will let all jobs in the queue to be finished.
func (l *Limiter) Stop(allow bool) (cancelled int) {
	if l.ifStatusIs(true, false) { // change accept status to false, so worker can't take a job
		if allow { // allow queue to be finished
			return 0
		}
		// drain all jobs in the queue
		cancelled = 0
		for len(l.queue) > 0 {
			<-l.queue
			cancelled += 1
		}

		return cancelled // return how many jobs in queue has been cancelled (drained)
	}
	return 0
}

// Status shows current limiter status. (state: currently taking a job)
func (l *Limiter) Status() (state bool, activeWorkers, activeQueue int) {
	if l.mu != nil {
		return l.ifStatusIs(true, true), len(l.worker), len(l.queue)
	}
	return false, 0, 0
}

// IsActive will return true if there's something running.
// This can be used as a part of wait-all-jobs.
func (l *Limiter) IsActive() bool {
	if s, w, q := l.Status(); s == false && w+q == 0 {
		return false
	}
	return true
}

// Close will attempt to close the limiter. If a job is currently running, it will return false.
func (l *Limiter) Close() (ok bool) {
	l.Stop(false)     // if already stopped, this will ignore
	if l.IsActive() { // if the limiter is still active, return false
		return false
	}

	close(l.queue)  // this will also close the monitor()
	close(l.worker) // DO NOT DRAIN THE WORKER: when all jobs are done, its size should be 0 anyway.
	close(l.mu)
	l.mu = nil // only after closing all channels, mu will set nil in case Init() is called later.
	return true
}