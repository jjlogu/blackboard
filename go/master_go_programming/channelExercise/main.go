package main

import "fmt"

func main() {
	/*
		Coding Exercise #1
		1. Using the var keyword, declare a bidirectional unbuffered channel called c1 that works with values of type float64
		2. Using the make() built-in function declare and initialize a receive-only channel called c2 and a send-only channel called c3. Both work with data of type rune.
		3. Declare a bidirectional buffered channel  called c4 with a capacity of 10 ints.
		4. Print out the type of all the channels declared.
	*/

	var c1 chan float64

	c2 := make(chan<- rune)
	c3 := make(<-chan rune)

	c4 := make(chan [10]int)

	fmt.Printf("c1: %T c2: %T c3: %T c4: %T ", c1, c2, c3, c4)

	/*
		Coding Exercise #2
		Create a function literal (a.k.a. anonymous function) that sends the string value if receives as argument to main func using a channel.
	*/
	c5 := make(chan string)
	go func() {
		c5 <- "Hello Channel!"
	}()

	fmt.Printf("\nfrom c5: %v\n", <-c5)

	/*
		Coding Exercise #4

		Create a goroutine named power() that has one parameter of type int, calculates the square value of its parameter and then sends  the result into a channel.
		In the main function launch 50 goroutines that calculate the square values of all numbers between 1 and 50 included.
		Print out the square values.
		A square(or raising to power 2) is the result of multiplying a number by itself. e.g., 25 is the square of 5.
	*/
	c6 := make(chan int)
	for i := 1; i <= 50; i++ {
		go power(i, c6)
		fmt.Printf("%d = %d\n", i, <-c6)
	}

}
func power(i int, c chan<- int) {
	c <- i * i
}
