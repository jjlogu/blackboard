package main

import (
	"fmt"
	"math"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func sum(m float64, n float64, wg *sync.WaitGroup) {
	fmt.Printf("%.2f\n", m+n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHello("Mr. Wick", &wg)

	/*
		1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.
		Format the result with 2 decimal points.
		2. From main launch 3 goroutines that execute the function you have just created (sum)
		3. Synchronize the goroutines and the main function using WaitGroups
	*/
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sum(2., 4., &wg)
	}

	/*
		1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.
		2. Launch the function as a goroutine and synchronize it with main using WaitGroups
	*/

	wg.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		fmt.Printf("Anomy func: %f\n", math.Sqrt(f))
		wg.Done()
	}(64., &wg)

	/*
		Change the code from Exercise #3 and launch 50 goroutines that calculate concurrently the square root of all the numbers between 100 and 149 (both included).
	*/
	for i := 100; i < 150; i++ {
		wg.Add(1)
		go func(f float64, wg *sync.WaitGroup) {
			fmt.Printf("Sqrt %f: %f\n", f, math.Sqrt(f))
			wg.Done()
		}(float64(i), &wg)
	}

	wg.Wait()
}
