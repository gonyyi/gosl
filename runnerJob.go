// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl

// Job is an interface that will be triggered
type Job interface {
	ID() string // shows ID
	Accept()    // This will be hit when accepted
	Reject()    // When failed or rejected
	Cancel()    // Originally accepted but later removed from the queue
	Run()       // Runs
}

func NewJob(ID string, OnRun func()) *simpleJob {
	return &simpleJob{
		id:    ID,
		FnRun: OnRun,
	}
}

type simpleJob struct {
	id       string // ID holds the job identifier
	FnAccept func() // When accepted
	FnReject func() // When rejected
	FnCancel func() // When accepted but later cancelled from QUEUE
	FnRun    func() // When started running
}

func (r *simpleJob) ID() string {
	return r.id
}
func (r *simpleJob) Accept() {
	if r.FnAccept != nil {
		r.FnAccept()
	}
}
func (r *simpleJob) Reject() {
	if r.FnReject != nil {
		r.FnReject()
	}
}
func (r *simpleJob) Cancel() {
	if r.FnCancel != nil {
		r.FnCancel()
	}
}
func (r *simpleJob) Run() {
	if r.FnRun != nil {
		r.FnRun()
	}
}

