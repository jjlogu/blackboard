/////////////////////////////////
// Slices in Go
// Go Playground: https://play.golang.org/p/UvLFUBdoL9v
/////////////////////////////////

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	// declaring a slice using a slice literal
	numbers := [12]int{2, 3, 4, 5, 6, 7, 8} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers)                    // => 2 3 4 5 6 7 8 0 0 0 0 0]
	fmt.Printf("%T\n", numbers)             //=> [12]int

	n1 := numbers[1:5]
	n2 := n1[:2]

	fmt.Println(n1, n2) // => [3 4 5 6] [3 4]
	n2[0] = 123
	n := append(n2, 100)
	fmt.Println(numbers)   // => [2 123 4 100 6 7 8 0 0 0 0 0]
	fmt.Println(n1, n2, n) // => [123 4 100 6] [123 4] [123 4 100]

	// Print slice Header
	fmt.Printf("N1: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&n1))) // => N1: &reflect.SliceHeader{Data:0xc000094068, Len:4, Cap:11}
	fmt.Printf("N2: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&n2))) // => N2: &reflect.SliceHeader{Data:0xc000094068, Len:2, Cap:11}
	fmt.Printf("N: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&n)))   // => N: &reflect.SliceHeader{Data:0xc000094068, Len:3, Cap:11}

	fmt.Printf("&numbers: %p, &numbers[1]: %p\n", &numbers, &numbers[1])          // => &numbers: 0xc00009c060, &numbers[1]: 0xc00009c068
	fmt.Printf("size of %T is %d bytes\n", numbers[1], unsafe.Sizeof(numbers[1])) // => size of int is 8 bytes

}
