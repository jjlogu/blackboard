package main

import (
	"fmt"

	mypack "packages/mypackages/arithmetic"
	/* "github.com/jjlogu/go_math/calc"
	"github.com/jjlogu/go_math/hello" */)

func main() {
	/* hello.SayHello("tamil")
	fmt.Printf("Sum: %d, sub: %d", calc.Add(5, 5), calc.Subract(10, 5)) */

	fmt.Println("IsPrime(13): ", mypack.IsPrime(13))
}
