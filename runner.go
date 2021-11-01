// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// NewRunners will create Runners
func NewRunners(queueSize uint, workers uint, waitAdd bool, finishQueue bool) *Runners {
	rb := &Runners{
		waitAdd:      waitAdd,
		finishQueue:  finishQueue,
		chStopSignal: make(chan struct{}, 1),
		chWorker:     make(chan struct{}, workers),
	}
	// If AsyncQueue is false, the go routine will keep try to squeeze in,
	// as a result, although chJobs size is small, there can be a lot of
	// jobs waiting as a goroutine.
	// I think it's better just have them into the unbuffered channel.
	if queueSize > 0 {
		rb.chJobs = make(chan Job, queueSize)
	} else {
		rb.chJobs = make(chan Job)
	}
	return rb
}

// Runners enables Job to be run concurrently.
type Runners struct {
	Logger Logger

	// Channels
	chJobs       chan Job
	chWorker     chan struct{} // channel for limiting how many chJobs can run at a time
	chStopSignal chan struct{}

	// Status
	acceptNewJob     bool // if false, Runners will NOT take new Job from Add(Job)
	runnerBoxStopped bool // if acceptNewJob == false AND no jobs are chWorker, this will be true

	ExitFn func() // runs when Runners finishes

	finishQueue bool // if finishQueue is true, it will wait until everything in the chJobs finishes.
	waitAdd     bool // if true, Add(Job) may wait when chJobs is full, but it will give a status ok.

	stats struct {
		accepted  uint // accepted
		rejected  uint // rejected
		finished  uint // finished
		cancelled uint // Accepted, but later rejected
	}
}

// Stats returns job status
func (b *Runners) Stats() (rejected, accepted, cancelled, finished uint) {
	return b.stats.rejected, b.stats.accepted, b.stats.cancelled, b.stats.finished
}

// SetLoggerOutput will enable Runners's builtin debugger
func (b *Runners) SetLoggerOutput(debug Writer) {
	b.Logger = b.Logger.SetOutput(debug)
}

// Stopped will check if the box is completely stopped.
func (b *Runners) Stopped() bool {
	return b.runnerBoxStopped
}

// Stop will stop taking new runners, and depend on setting, it will cancel
// jobs in the currenet queue.
func (b *Runners) Stop() {
	b.Logger.KeyBool("AcceptNewRunner", false)
	b.acceptNewJob = false // no more accepting
	b.chStopSignal <- struct{}{}
}

// add will add the runner to the queue
func (b *Runners) add(f Job) (ok bool) {
	defer func() {
		if v := recover(); v != nil {
			b.Logger.KeyString("Panic().ID", f.ID())
		}
	}()

	if b.acceptNewJob == false {
		defer func() {
			if v := recover(); v != nil {
				b.Logger.KeyString("Panic(RejectFn).ID", f.ID())
			}
		}()
		f.Reject()
		b.stats.rejected += 1
		return false
	}

	b.chJobs <- f
	f.Accept()
	b.stats.accepted += 1
	return true
}

// Add will add runner to the queue
func (b *Runners) Add(f Job) (ok bool) {
	if b.waitAdd {
		return b.add(f)
	} else {
		go b.add(f)
	}
	return true
}

// Run will start the Runners. If it was already started, it won't do anything.
func (b *Runners) Run() *Runners {
	if b.acceptNewJob == false {
		b.acceptNewJob = true
		b.Logger.KeyInt("Conf.Concurrency", cap(b.chWorker))
		b.Logger.KeyBool("Conf.FinishQueue", b.finishQueue)
		if b.waitAdd {
			b.Logger.KeyBool("Conf.WaitAdd", b.waitAdd)
		}
		b.Logger.KeyInt("Conf.QueueSize", cap(b.chJobs))
		b.Logger.KeyBool("Conf.ExitFn", b.ExitFn != nil)
		b.Logger.String("Started")
		go b.run()
	}
	return b
}

// JobQueue will return currently runners in the queue
func (b *Runners) JobQueue() int {
	return len(b.chJobs)
}

// JobsRunning will return currently running runner tasks
func (b *Runners) JobsRunning() int {
	return len(b.chWorker)
}

// run will receive channels and run the job
func (b *Runners) run() {
wait:
	for {
		select {
		case runner := <-b.chJobs:
			if b.acceptNewJob || b.finishQueue {
				// This blocks how many can run concurrently
				b.chWorker <- struct{}{}
				go func() {
					defer func() {
						<-b.chWorker // when job ended, take one worker out
						if r := recover(); r != nil {
							b.Logger.KeyString("Panic(runFn).RunnerID", runner.ID())
						}
					}()
					b.Logger.KeyString("Run().RunnerID", runner.ID())
					runner.Run()
					b.stats.finished += 1
				}()

			} else {
				// This blocks how many can run concurrently
				b.chWorker <- struct{}{}
				go func() {
					defer func() {
						<-b.chWorker // in case FnReject failed, it will still go on
						if recover() != nil {
							b.Logger.KeyString("Panic(rejectFn).RunnerID", runner.ID())
						}
					}()
					b.Logger.KeyString("Reject().RunnerID", runner.ID())
					runner.Cancel()
					b.stats.cancelled += 1
				}()
			}
		case <-b.chStopSignal:
			b.acceptNewJob = false // Runners.Add() will not take new Job.

			b.Logger.KeyInt("Signal(STOP).Worker.Count()", b.JobsRunning())
			b.Logger.KeyInt("Signal(STOP).JobQueue.Count()", b.JobQueue())
			b.Logger.KeyBool("Signal(STOP).AcceptNewRunner", b.acceptNewJob)

			go b.runWait()
		default:
			if b.runnerBoxStopped {
				b.close()
				break wait
			}
		}
	}
}

// runWait will wait until there's no job and no queue left.
func (b *Runners) runWait() {
	for {
		if b.JobQueue() == 0 && b.JobsRunning() == 0 {
			b.Logger.String("Stop(LISTENING)")
			break
		}
	}
	b.runnerBoxStopped = true
}

// Wait will wait until the box is completely stopped
func (b *Runners) Wait() {
	for {
		if b.runnerBoxStopped {
			break
		}
	}
}

// close will close all the channel if open
func (b *Runners) close() {
	// IF CHANNEL HAS NOT BEEN CLOSED, CLOSE IT.
	select {
	case <-b.chStopSignal:
	default:
		close(b.chStopSignal)
		b.Logger.String("Close(StopSignal)")
	}

	select {
	case <-b.chWorker:
	default:
		close(b.chWorker)
		b.Logger.String("Close(Worker)")
	}

	select {
	case <-b.chJobs:
	default:
		close(b.chJobs)
		b.Logger.String("Close(JobQueue)")
	}

	if b.ExitFn != nil {
		b.Logger.String("Start:ExitFn()")
		b.ExitFn()
	}

	b.runnerBoxStopped = true
	b.Logger.KeyBool("Runners.Stopped", b.runnerBoxStopped)
}

