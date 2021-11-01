// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 11/1/2021

package gosl_test

import (
	"fmt"
	"github.com/gonyyi/gosl"
	"time"
)

func Example_RunnerBox_NewRunnerBox() {
	// func TestRunnerBox(t *testing.T) {
	// 3 can run at a time
	rb := gosl.NewRunners(10, 50, true, false)
	rb.Run()
	rb.Run()

	addRunner := func(id int) {
		tmp := gosl.NewJob(gosl.Itoa(id), func() {
			// Create a fake job that takes 1 second to run
			fmt.Printf("\t[%03d] %s (%d)\n", id, "START", 3)
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("\t[%03d] %s (%d)\n", id, "FINISHED", 4)

		})
		tmp.FnAccept = func() {
			println("ACCEPTED")
		}
		tmp.FnReject = func() {
			println("REJECTED")
		}
		tmp.FnCancel = func() {
			println("CANCELLED!!")
		}

		res := rb.Add(tmp)
		if res != true {
			fmt.Printf("\t[%03d] %s (%d)\n", id, "REJECTED", 2)
		} else {
			fmt.Printf("\t[%03d] %s (%d)\n", id, "ACCEPTED", 1)
		}
	}

	go func() {
		time.Sleep(time.Second)
		println("\t** STOP REQUESTED **")
		rb.Stop()
	}()

	for i := 0; i < 200; i++ {
		addRunner(i + 1)
	}

	// rb.Stop()
	rb.Wait()

	r, a, c, f := rb.Stats()
	fmt.Printf("Rejected:  %d\nAccepted:  %d\nCancelled: %d\nFinished:  %d\nTotal:     %d (R+C+F)\n", r, a, c, f, r+c+f)
}

