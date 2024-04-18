/////////////////////////////////
// Defer Function Calls
// Go Playground: https://play.golang.org/p/KGulBq828NT
/////////////////////////////////

package main

import (
	"fmt"
)

func foo() {
	fmt.Println("This is foo()!")
}

func bar() {
	fmt.Println("This is bar()!")
}

func foobar() {
	fmt.Println("This is foobar()!")
}

func incr(i *int) int {
	(*i) = (*i) + 1
	return *i
}

func main() {

	// a defer statement defers or postpones the execution of a function until the surrounding function returns.

	// by deferring foo() it will execute it just before exiting the surrounding function which is main()
	defer foo()
	bar()

	fmt.Println("Just a string after deferring foo() and calling bar()")

	// if there are more functions deferred, Go will execute them in the reverse order they were deferred
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	defer foobar()

	/*
	   When executing the program the fallowing output is printed out:

	   This is bar()!
	   Just a string after deferring foo() and calling bar()
	   This is foobar()!
	   This is foo()!
	*/

	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	i := 1
	defer fmt.Println("world", incr(&i))

	fmt.Println("hello", incr(&i))

	/*
		hello 3
		world 2
	*/

}
