// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"github.com/gonyyi/gosl"
	"time"
)

func Example_Runner_Runner() {
	// func TestRunner(t *testing.T) {
	// 5 can run at a time, and can hold 10 in the queue
	rb := gosl.NewRunner(50, 10, true, true).Run()
	// rb.SetLoggerOutput(os.Stdout)

	newJob := func(id int) gosl.Job {
		// Create a fake job that takes 1 second to run
		tmp := gosl.NewJob(gosl.Itoa(id), func() {
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
	for i := 0; i < 2000; i++ {
		if rb.Add(newJob(i + 1)) {
			/* Runner rejected the job */
		} else {
			/* Runner accepted the job */
		}
	}

	// Wait until runner is closed
	rb.WaitClose()

	// Print stats
	println(rb.String())
}

