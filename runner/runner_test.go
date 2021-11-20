// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/5/2021

package runner_test

import (
	"github.com/gonyyi/gosl"
	"github.com/gonyyi/gosl/runner"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	// 5 can run at a time, and can hold 10 in the queue
	rb := runner.NewRunner(50, 10, true, true).Run()
	// rb.SetLoggerOutput(os.Stdout)

	newJob := func(id int) runner.Job {
		// Create a fake job that takes 1 second to run
		tmp := runner.NewJob(gosl.Itoa(id), func() {
			time.Sleep(time.Millisecond * 500)
		})

		// tmp.SetReject(func() { /* Do something */ })
		// tmp.SetAccept(func() { /* Do something */ })
		// tmp.SetCancel(func() { /* Do something */ })
		// tmp.SetRun(func() { /* Do something */ })
		return tmp
	}

	// Cancel job 1 seconds later
	go func() {
		time.Sleep(time.Second * 1)
		rb.Stop()
	}()

	// Create 2,000 jobs and add it to runner
	for i := 0; i < 200; i++ {
		if rb.Add(newJob(i + 1)) {
			/* Runner rejected the job */
		} else {
			/* Runner accepted the job */
		}
	}

	// Wait until runner is closed
	rb.WaitClose()

	{
		r, a, c, f := rb.Stats()
		gosl.Test(t, 200, r+a)
		gosl.Test(t, 200, c+f+r)
	}
}

