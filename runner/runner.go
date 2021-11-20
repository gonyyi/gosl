// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package runner

import "github.com/gonyyi/gosl"

// WARNING:
//     Code below will be slower than using standard libraries.
//     This is just a test to see if it can be built without
//     using any libraries (not even built-in), but only using
//     builtin functions of the language.

// NewRunner will create Runner. Params are as below.
// - queue:   size of job channel (how many jobs this should hold)
// - workers: how many max number of jobs should run at a given time
// - async:   if adding a job should wait until accepted. (safe way)
// - finish:  if stopping is requested, should jobs in queue to be finished.
//
//  | queue | workers | async | finish |  rej | accpt | cancel | compl |
//  |------:|--------:|:------|:-------|-----:|------:|-------:|------:|
//  |    10 |      50 | true  | true   | 1888 |   112 |      0 |   112 |
//  |    10 |      50 | true  | false  | 1888 |   112 |     11 |   101 |
//  |    10 |      50 | false | true   |    0 |  2000 |      0 |  2000 |
//  |    10 |      50 | false | false  |    0 |  2000 |   1899 |   101 |
func NewRunner(queue uint, workers uint, async bool, finish bool) *Runner {
	rb := &Runner{
		doAsync:      async,
		doFinish:     finish,
		chJobs:       make(chan Job, queue),
		chStopSignal: make(chan struct{}, 1),
		chWorker:     make(chan struct{}, workers),
		mu:           make(gosl.Mutex, 1),
		available:    gosl.MuBool{}.Init(),
		closed:       gosl.MuBool{}.Init(),
		stats:        newRunnerStats(),
	}
	return rb
}

// Runner enables Job to be run concurrently.
type Runner struct {
	Logger gosl.Logger

	// Channels
	chJobs       chan Job
	chWorker     chan struct{} // channel for limiting how many chJobs can run at a time
	chStopSignal chan struct{}

	// Status
	mu        gosl.Mutex
	available gosl.MuBool
	closed    gosl.MuBool // runner is closed

	// ExitFn runs when all the job finished
	ExitFn func() // runs when Runner finishes

	// Configuration
	doFinish bool // if doFinish is true, it will wait until everything in the chJobs finishes.
	doAsync  bool // if true, Add(Job) may wait when chJobs is full, but it will give a status ok.

	stats runnerStats
}

// Stats will show job counts
func (b *Runner) Stats() (rejected, accepted, cancelled, completed int) {
	rejected, accepted, cancelled, completed, _ = b.stats.summary()
	return
}

// String, for now will return summary of job.
func (b *Runner) String() string {
	return b.stats.String()
}

// SetLoggerOutput will enable Runner's builtin debugger
func (b *Runner) SetLoggerOutput(debug gosl.Writer) {
	b.Logger = b.Logger.SetOutput(debug)
}

// Closed will check if the box is completely stopped.
func (b *Runner) Closed() bool {
	return b.closed.Get()
}

// Stop will stop taking new runners, and depend on setting,
// it will cancel jobs in the current queue. (see "finish" param)
func (b *Runner) Stop() {
	b.chStopSignal <- struct{}{}
}

// WaitClose will wait until the box is completely stopped
func (b *Runner) WaitClose() {
	for {
		if b.closed.Get() {
			break
		}
	}
}

// Add will add runner to the queue
func (b *Runner) Add(f Job) (ok bool) {
	if b.doAsync {
		return b.add(f)
	} else {
		go b.add(f)
	}
	return true
}

// Run will start the Runner. If it was already started, it won't do anything.
func (b *Runner) Run() *Runner {
	if b.available.Get() == false {

		b.available.Set(true)
		b.Logger.KeyBool("SET Runner.Available", true)

		b.Logger.KeyInt("Conf.Concurrency", cap(b.chWorker))
		b.Logger.KeyBool("Conf.FinishQueue", b.doFinish)
		if b.doAsync {
			b.Logger.KeyBool("Conf.WaitAdd", b.doAsync)
		}
		b.Logger.KeyInt("Conf.QueueSize", cap(b.chJobs))
		b.Logger.KeyBool("Conf.ExitFn", b.ExitFn != nil)
		b.Logger.String("Runner.Run()")
		go b.run()
	}
	return b
}

// Queue will return currently runners in the queue
func (b *Runner) Queue() int {
	return len(b.chJobs)
}

// Running will return currently running runner tasks
func (b *Runner) Running() int {
	return len(b.chWorker)
}

// add will add the runner to the queue
func (b *Runner) add(f Job) (ok bool) {
	defer func() {
		if v := recover(); v != nil {
			b.Logger.KeyString("Panic().JobID", f.ID())
		}
	}()

	if b.available.Get() == false {
		defer func() {
			if v := recover(); v != nil {
				b.Logger.KeyString("Panic(RejectFn).JobID", f.ID())
			}
		}()
		f.Reject()
		// b.stats.rejected += 1
		b.stats.Rejected.Add(1)
		return false
	}

	b.chJobs <- f
	f.Accept()
	b.stats.Accepted.Add(1)
	return true
}

// run will receive channels and run the job
func (b *Runner) run() {
	for {
		if b.closed.Get() {
			b.close()
			break
		}

		select {
		case job := <-b.chJobs:
			if b.doFinish || b.available.Get() {
				// if b.getAcceptNewJob() || b.doFinish {
				// This blocks how many can run concurrently
				b.chWorker <- struct{}{}
				go func() {
					{
						buf := gosl.GetBuffer()
						buf.WriteString("RunnerID(").WriteString(job.ID()).WriteString(").Run()")
						b.Logger.String(buf.String())
						buf.Free()
					}

					job.Run()
					b.stats.Completed.Add(1)
					<-b.chWorker // when job ended, take one worker out
				}()

			} else {
				// This blocks how many can run concurrently
				b.chWorker <- struct{}{}
				go func() {
					{
						buf := gosl.GetBuffer()
						buf.WriteString("RunnerID(").WriteString(job.ID()).WriteString(").Reject()")
						b.Logger.String(buf.String())
						buf.Free()
					}

					job.Cancel()
					b.stats.Cancelled.Add(1)
					<-b.chWorker
				}()
			}
		case <-b.chStopSignal:
			b.available.Set(false)
			b.Logger.KeyBool("SET Runner.Available", false)
			{
				buf := gosl.GetBuffer()
				buf.WriteString("Jobs:Queue=").WriteInt(b.Queue()).WriteBytes(',').
					WriteString("Running=").WriteInt(b.Running())
				b.Logger.String(buf.String()) // this can also be done by buf.WriteTo(Writer)
				buf.Free()
			}

			go func() {
				b.waitToFinish()
			}()
		default:
			// having default with nothing here makes it keep checking status
			// don't wait on the signal.
		}
	}
}

// waitToFinish will wait until there's no job and no queue left.
func (b *Runner) waitToFinish() {
	for {
		if b.Queue() == 0 && b.Running() == 0 {
			b.Logger.String("Jobs:Queue=0,Running=0")
			break
		}
	}

	b.closed.Set(true)
	b.Logger.KeyBool("SET Runner.Close", true)
}

// close will close all the channel if open
func (b *Runner) close() {
	// IF CHANNEL HAS NOT BEEN CLOSED, CLOSE IT.
	select {
	case <-b.chStopSignal:
	default:
		close(b.chStopSignal)
		b.Logger.String("CloseChannel:StopSignal")
	}

	select {
	case <-b.chWorker:
	default:
		close(b.chWorker)
		b.Logger.String("CloseChannel:Worker")
	}

	select {
	case <-b.chJobs:
	default:
		close(b.chJobs)
		b.Logger.String("CloseChannel:Queue")
	}

	if b.ExitFn != nil {
		b.Logger.String("Runner.ExitFn()")
		b.ExitFn()
	}

	b.Logger.String("Runner.Status: Closed")
}

// *************************************************************************
// RUNNER STATS
// *************************************************************************

func newRunnerStats() runnerStats {
	return runnerStats{
		Rejected:  gosl.MuInt{}.Init(),
		Accepted:  gosl.MuInt{}.Init(),
		Cancelled: gosl.MuInt{}.Init(),
		Completed: gosl.MuInt{}.Init(),
	}
}

type runnerStats struct {
	Rejected  gosl.MuInt
	Accepted  gosl.MuInt
	Cancelled gosl.MuInt
	Completed gosl.MuInt
}

func (s *runnerStats) String() string {
	_, _, _, _, out := s.summary()
	return out
}

func (s *runnerStats) summary() (rejected, accepted, cancelled, completed int, out string) {
	rejected, accepted, cancelled, completed = s.Rejected.Get(), s.Accepted.Get(), s.Cancelled.Get(), s.Completed.Get()
	buf := gosl.GetBuffer()
	buf.WriteString("Stats.Job.Rejected: ").WriteInt(rejected).WriteBytes('\n').
		WriteString("Stats.Job.Accepted: ").WriteInt(accepted).WriteBytes('\n').
		WriteString("Stats.Job.Accepted=>Cancelled: ").WriteInt(cancelled).WriteBytes('\n').
		WriteString("Stats.Job.Accepted=>Completed: ").WriteInt(completed).WriteBytes('\n').
		WriteString("Stats.Job.ReceivedTotal: ").WriteInt(completed + cancelled + rejected)
	out = buf.String()
	buf.Free()
	return
}

