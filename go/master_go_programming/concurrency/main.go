/////////////////////////////////
// Getting Information
// Go Playground: https://play.golang.org/p/OHyIck2IwkS
/////////////////////////////////

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Declaring two functions: f1 and f2
func f1(wg *sync.WaitGroup) { // wg is passed as a pointer
	fmt.Println("f1(goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		// sleep for a second to simulate an expensive task.
		time.Sleep(time.Second)

	}
	fmt.Println("f1 execution finished")

	//3.
	// Before exiting, call wg.Done() in each goroutine
	// to indicate to the WaitGroup that the goroutine has finished executing.
	wg.Done()
	//or:
	// (*wg).Done()
}

func f2() {
	fmt.Println("f2 execution started")
	time.Sleep(time.Second)

	for i := 5; i < 8; i++ {
		fmt.Println("f2(), i=", i)

	}
	fmt.Println("f2 execution finished")

}

func main() {
	// NumCPU returns the number of logical CPUs or cores usable by the current process.
	fmt.Println("No. of CPUs:", runtime.NumCPU()) // => No. of CPUs: 4

	// NumGoroutine returns the number of goroutines that currently exists.
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => No. of Goroutines: 1

	// GOOS and GOARCH are environment variables
	fmt.Println("OS:", runtime.GOOS)     // => OS: linux
	fmt.Println("Arch:", runtime.GOARCH) // => Arch: amd64

	//  GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns
	//  the previous setting.
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // => GOMAXPROCS: 4
	// If n < 1, it does not change the current setting.

	fmt.Println("main execution started")

	// 1.
	// Create a new instance of sync.WaitGroup (we’ll call it symply wg)
	// This WaitGroup is used to wait for all the goroutines that have been launched to finish.
	var wg sync.WaitGroup

	// 2.
	// Call wg.Add(n) method before attempting to
	// launch the go routine.
	wg.Add(1) //  n which is 1 is the number of goroutines to wait for

	// Launching a goroutine
	go f1(&wg) // it takes in a pointer to sync.WaitGroup

	// No. of running goroutines
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => 2

	// calling other functions:
	f2()

	// 4.
	// Finally, we call wg.Wait()to block the execution of main() until the goroutines
	// in the WaitGroup have successfully completed.
	wg.Wait()

	fmt.Println("main execution stopped")

	const gr = 100
	//var wg sync.WaitGroup
	wg.Add(gr * 2)

	// declaring a shared value
	var n int = 0

	// 1.
	// Declaring a mutex. It's available in sync package
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)

			// 2.
			// Lock the access to the shared value
			m.Lock()
			n++

			// 3.
			// Unlock the variable after it's incremented
			m.Unlock()

			wg.Done()
		}()

		// Doing the same for the 2nd goroutine
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()

	}

	wg.Wait()

	// printing the final value of n
	fmt.Println(n) // the final final of n will be always 0

}
