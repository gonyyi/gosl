// Copyright 2021 Gon Y. Yi <https://gonyyi.com/copyright>
// Last Updated: 10/18/2021

package gosl

func NewRunner(ID uint, OnRun func(), OnReject func()) Runner {
	return Runner{
		ID:       ID,
		OnRun:    OnRun,
		OnReject: OnReject,
	}
}

type Runner struct {
	ID    uint   // ID is optional
	OnRun func() // OnStart will be main job running
	OnReject func() // OnReject is for job that didn't even run
}

type RunnerBox struct {
	lDebug  Logger
	queue   chan Runner
	active  chan struct{} // just to limit concurrency
	stop    chan struct{}
	closed  bool // if closed=true, it will not take new Runner from Add(Runner)
	stopped bool // if closed AND no jobs are running, this will be true

	ExitFn       func() // runs when `run` finishes
	StopFn       func() // when stop request received -- such as time.Sleep can be used here
	WaitAllQueue bool   // if WaitAllQueue is true, it will wait until everything in the queue finishes.
	asyncQueue   bool   // if true, Add(Runner) may wait when queue is full
}

// NewRunnerBox will create RunnerBox
func NewRunnerBox(queueSize int) *RunnerBox {
	rb := &RunnerBox{
		stop: make(chan struct{}, 1),
	}
	// If AsyncQueue is false, the go routine will keep try to squeeze in,
	// as a result, although queue size is small, there can be a lot of
	// jobs waiting as a goroutine.
	// I think it's better just have them into the unbuffered channel.
	if queueSize > 0 {
		rb.queue = make(chan Runner, queueSize)
		rb.asyncQueue = true
	} else {
		rb.queue = make(chan Runner)
		rb.asyncQueue = false
	}
	return rb
}

func (RunnerBox) doNothingOnReject() {}

func (b *RunnerBox) SetLoggerOutput(debug Writer) {
	b.lDebug = b.lDebug.SetOutput(debug)
}

func (b *RunnerBox) Stopped() bool {
	return b.stopped
}

// Add will add to queue
func (b *RunnerBox) Add(f Runner) (ok bool) {
	if b.closed {
		defer func() {
			if recover() != nil {
				b.lDebug.KeyInt("panic.OnReject().id", int(f.ID))
			}
		}()
		f.OnReject()
		return false
	}
	start := func() {
		defer func() {
			if recover() != nil {
				b.lDebug.KeyInt("panic.runner.id", int(f.ID))
			}
		}()
		if f.OnRun == nil {
			f.OnRun = DoNothing
		}
		if f.OnReject == nil {
			f.OnReject = b.doNothingOnReject
		}
		b.queue <- f
	}

	if b.asyncQueue {
		start()
	} else {
		go start()
	}

	return true
}

func (b *RunnerBox) Run(concurrency int) {
	b.active = make(chan struct{}, concurrency)
	b.lDebug.KeyInt("conf.Concurrency", concurrency)
	b.lDebug.KeyBool("conf.WaitAllQueue", b.WaitAllQueue)
	b.lDebug.KeyBool("conf.queue.async", b.asyncQueue)
	if b.asyncQueue {
		b.lDebug.KeyInt("conf.queue.capacity", cap(b.queue))
	}
	b.lDebug.KeyBool("conf.ExitFn", b.ExitFn != nil)
	b.lDebug.KeyBool("conf.StopFn", b.StopFn != nil)
	b.lDebug.String("started")

	go b.run()
}

func (b *RunnerBox) Stop() {
	b.stop <- struct{}{}
}

func (b *RunnerBox) Queue() int {
	return len(b.queue)
}

func (b *RunnerBox) Running() int {
	return len(b.active)
}

func (b *RunnerBox) run() {
	for {
		select {
		case runner := <-b.queue:
			if b.closed == false || b.WaitAllQueue {
				b.active <- struct{}{}

				go func() {
					defer func() {
						<-b.active
						if r := recover(); r != nil {
							var msg string
							switch x := r.(type) {
							case string:
								msg = x
							case error:
								msg = x.Error()
							default:
								msg = "unknown"
							}

							b.lDebug.KeyInt("panic.runner.id", int(runner.ID))
							b.lDebug.KeyString("panic.runner.err", msg)
						}
					}()
					runner.OnRun()
				}()

			} else { // After closed(), clear the queue
				b.active <- struct{}{}
				go func() {
					defer func() {
						if recover() != nil {
							b.lDebug.KeyInt("panic.onReject().id", int(runner.ID))
						}
						<-b.active // in case OnReject failed, it will still go on
					}()
					b.lDebug.KeyInt("job.queue.removed", int(runner.ID))
					runner.OnReject()
				}()
			}
		case <-b.stop:
			b.closed = true // RunnerBox.Add() will not take new Runner.

			b.lDebug.String("stop()")
			b.lDebug.KeyInt("running.count", b.Running())
			b.lDebug.KeyInt("queue.count", b.Queue())
			b.lDebug.KeyBool("queue.closed", b.closed)

			if b.StopFn != nil {
				b.lDebug.String("start(StopFn)")
				b.StopFn()
			}

			b.lDebug.String("stop.wait(queue, running)")
			b.lDebug.KeyInt("job.queue.count", b.Queue())
			b.lDebug.KeyInt("job.running.count", b.Running())

			go func() {
				for {
					if b.Queue() == 0 && b.Running() == 0 {
						b.lDebug.KeyBool("job.queue.empty", true)
						b.lDebug.KeyBool("job.running.empty", true)
						break
					}
				}
				b.stopped = true
			}()
		}
		if b.stopped {
			close(b.stop)
			close(b.active)
			close(b.queue)
			break
		}
	}

	if b.ExitFn != nil {
		b.lDebug.String("start(ExitFn)")
		b.ExitFn()
	}

	b.lDebug.KeyBool("stopped", b.stopped)
}
