// You are given two arbitrarily large numbers,
// stored one digit at a time in a slice.
// The first must be added to the second,
// and the second must be reversed before addition.
//
// The goal is to calculate the sum of those two sets of values.
//
// IMPORTANT NOTE:
// - The input can be any lengths (i.e: it can be 20+ digits long).
// - num1 and num2 can be different lengths.
//
// Sample Inputs:
// num1 = 123456
// num2 = 123456
//
// Sample Output:
// Result: 777777
//
// We would also like to see a demonstration of appropriate unit tests
// for this functionality.

package main

import (
	"fmt"
)

func Add(num1 []int, num2 []int) string {
	// Implement this
	var result string
	lnum1 := len(num1)
	lnum2 := len(num2)
	max := lnum2
	if lnum1 > lnum2 {
		max = lnum1
	}
	var n1, n2, tens, sum, v int
	for i := 0; i < max; i++ {
		if i < lnum1 {
			// fmt.Printf("lnum1: %v ", num1[lnum1-i-1])
			n1 = num1[lnum1-i-1]
		} else {
			n1 = 0
		}
		if i < lnum2 {
			// fmt.Printf(" lnum2: %v", num2[i])
			n2 = num2[i]
		} else {
			n2 = 0
		}
		sum = n1 + n2 + tens
		tens, v = sum/10, sum%10

		result = fmt.Sprintf("%d", v) + result

		// fmt.Printf(" (tens: %v v:%v sum:%v n1:%v n2:%v)", tens, v, sum, n1, n2)
	}
	if tens != 0 {
		result = fmt.Sprintf("%d", tens) + result
	}
	// fmt.Printf("\n%#v %#v\n", num1, num2)
	return result
}

func main() {
	num1 := []int{}
	num2 := []int{}

	num1length := 6
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i)
	}

	num2length := 6
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i)
	}

	result := Add(num1, num2)

	fmt.Println("Result:", result)
}
