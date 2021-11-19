// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package runner

// Job is an interface that will be triggered
type Job interface {
	ID() string // shows ID
	Accept()    // This will be hit when accepted
	Reject()    // When failed or rejected
	Cancel()    // Originally accepted but later removed from the queue
	Run()       // Runs
}

// NewJob creates new simple job
func NewJob(ID string, fnRun func()) *simpleJob {
	return &simpleJob{
		id:  ID,
		run: fnRun,
	}
}

// ******************************************************************************************
// SIMPLE JOB for RUNNER
// ******************************************************************************************

type simpleJob struct {
	id     string // ID holds the job identifier
	accept func() // When accepted
	reject func() // When rejected
	cancel func() // When accepted but later cancelled from QUEUE
	run    func() // When started running
}

// SetID takes a job ID and set
func (r *simpleJob) SetID(id string) *simpleJob {
	r.id = id
	return r
}

// SetReject takes a function and set reject function
func (r *simpleJob) SetReject(f func()) *simpleJob {
	r.reject = f
	return r
}

// SetAccept takes a function and set accept function
func (r *simpleJob) SetAccept(f func()) *simpleJob {
	r.accept = f
	return r
}

// SetCancel takes a function and set cancel function
func (r *simpleJob) SetCancel(f func()) *simpleJob {
	r.cancel = f
	return r
}

// SetRun takes a function and set run function
func (r *simpleJob) SetRun(f func()) *simpleJob {
	r.run = f
	return r
}

// ID returns the job ID
func (r *simpleJob) ID() string {
	return r.id
}

// Accept execute accept function if available
func (r *simpleJob) Accept() {
	if r.accept != nil {
		r.accept()
	}
}

// Reject execute reject function if available
func (r *simpleJob) Reject() {
	if r.reject != nil {
		r.reject()
	}
}

// Cancel execute cancel function if available
func (r *simpleJob) Cancel() {
	if r.cancel != nil {
		r.cancel()
	}
}

// Run execute main run function if available
func (r *simpleJob) Run() {
	if r.run != nil {
		r.run()
	}
}

