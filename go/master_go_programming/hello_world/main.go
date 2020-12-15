package main

import (
	"fmt"
)

func main() {

	var x = 5

	ptr := &x
	var p *int = ptr
	fmt.Printf("%T %v %v\n", p, p, *p)
	*p = 6
	pptr := &ptr
	ppptr := &pptr
	pppptr := &ppptr
	fmt.Printf("%T %v %v\n", ptr, ptr, *ptr)
	fmt.Printf("%T %v %v\n", pptr, pptr, **pptr)
	fmt.Printf("%T %v %v\n", ppptr, ppptr, ***ppptr)
	fmt.Printf("%T %v %v\n", pppptr, pppptr, ****pppptr)
}
