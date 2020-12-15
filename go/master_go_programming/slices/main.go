package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	numbers := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers = append(numbers, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n := []int{100, 200, 300}
	numbers = append(numbers, n...) // ... is the ellipsis operator
	fmt.Println(numbers)            // -> [2 3 10 20 30 40 100 200 300]

	//** Slice's Length and Capacity **//

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn) // => [10 20 30] [10 20 30] 3

	fmt.Println(strings.Repeat("#", 100))
	letters := []string{"a", "b", "d", "e"}
	fmt.Printf("%#v  len=%d cap=%d\n", letters, len(letters), cap(letters))

	letters = append(letters[1:2], "c")

	fmt.Printf("%#v  len=%d cap=%d\n", letters, len(letters), cap(letters))
	fmt.Printf("%#v  len=%d cap=%d\n", letters[0:3], len(letters), cap(letters))
	fmt.Println(strings.Repeat("#", 100))
	exercise()
}

/*
1. Declare a slice called nums with three float64 numbers.

2. Append the value 10.1 to the slice

3. In one statement append to the slice the values: 4.1,  5.5 and 6.6

4. Print out the slice

5. Declare a slice called n with two float64 values

6. Append n to nums

7. Print out the nums slice
*/
func exercise() {
	var t = []string{"Logu", "Sathya", "Seyon"}

	fmt.Println(strings.Repeat("#", 100))
	for _, v := range t {
		fmt.Println(v)
	}
	fmt.Println(strings.Repeat("#", 100))
	nums := []float64{1, 2, 3}
	fmt.Printf("%#v\n", nums)

	nums = append(nums, 10.1)
	fmt.Printf("%#v\n", nums)

	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Printf("%#v\n", nums)

	n := []float64{4.4, 5.5}

	nums = append(nums, n...)

	fmt.Printf("%#v\n", nums)

	if len(os.Args) == 3 {
		n1, err1 := strconv.Atoi(os.Args[1])
		n2, err2 := strconv.Atoi(os.Args[2])

		if err1 == nil || err2 == nil {

			fmt.Printf("Sum: %v, Product: %v\n", n1+n2, n1*n2)

		} else {
			fmt.Println("Invalid numbers passed!")
		}
	}

	fmt.Println(strings.Repeat("#", 100))

	nums1 := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	sum := 0

	for _, v := range nums1[2 : len(nums1)-2] {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println(strings.Repeat("#", 100))

	friends := []string{"Marry", "John", "Paul", "Diana"}
	var friends1 = make([]string, len(friends))

	fmt.Println(copy(friends1, friends))
	fmt.Printf("%#v\n", friends)

	friends1[1] = "Raju"
	fmt.Printf("%#v\n", friends)
	fmt.Printf("%#v\n", friends1)

	f2 := []string{}

	f2 = append(f2, friends...)
	fmt.Printf("%#v\n", f2)
	f2[0] = "Raju"
	fmt.Printf("%#v\n", friends)
	fmt.Printf("%#v\n", f2)

	fmt.Println(strings.Repeat("#", 100))
	//Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.  newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	var newYears []int
	// newYears = append(newYears, years[:3]...)
	// newYears = append(newYears, years[len(years)-3:]...)
	newYears = append(years[:3], years[len(years)-3:]...)

	fmt.Printf("%#v\n", newYears)

	fmt.Println(strings.Repeat("#", 100))
	s := "ஆன்லைனில்"
	fmt.Println(s[:3])
	fmt.Printf("Length in bytes: %v Length of strings: %v\n", len(s), utf8.RuneCountInString(s))

	for _, v := range s {
		fmt.Printf("%c %d\n", v, unsafe.Sizeof(v))
	}

	s1 := "ஆ"
	fmt.Printf("%v %v", s1, len(s1))

}
