package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var p = fmt.Printf

func main() {

	// stringExercise()
	//fileExercise()
	//structExercise()
	//funcExercise()
	// pointerExercise()
	// oopsExercise()
	interfaceExercise()
}
func stringExercise() {
	/*
			1. Using the var keyword declare a string called name and initialize it with your name.

		2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.

		3. Print the following string on multiple lines like this:

		Your name: `here the value of name variable`

		Country: `here the value of country variable`

		4. Print out the following strings:

		a) He says: "Hello"

		b) C:\Users\a.txt
	*/

	s := "ஆன்லைனில்"

	result := strings.Contains(s, "ன்")

	p("Result %v\n", result)

	p("Count %v\n", strings.Count(s, ""))

	s = ".!?.Logu!?.!"

	p("%v", strings.Trim(s, ".!?"))
	p("%T\n", strings.Split("Go for Go!", ""))

	var name = "Logu"
	country := "India"

	p("Your name: %v\nCountry: %v\n", name, country)
	p("%q %v\n", "Hello!", `C:\Users\a.txt`)

	s1 := "țară means country in Romanian"

	for i := 0; i < len(s1); i++ {
		p("%c-", s1[i])
	}

	for i, v := range s1 {
		p("\n%v %c", i, v)
	}

	s2 := "你好 Go!"

	var s2b = []byte(s2)

	// printing out the byte slice
	p("%#v\n", s2b)

	// iterating over the byte slice and printing out each index and byte in the byte slice
	for i, v := range s2b {
		fmt.Printf("%d -> %d\n", i, v)
	}

}

func fileExercise() {

	/*
		Create a new file in the current working directory called info.txt and then close the file. If the file cannot be created (no permissions, wrong path etc) then print out the error and stop the program (do error handling).
	*/
	file, err := os.OpenFile("info.txt", os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal("Error:", err)
	}

	/*
		Rename the file created at Exercise #1 to data.txt, Check if file exists before renaming it. If it doesn't exist print a message and stop the program.
	*/
	_, err = os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("info.txt not exists.")
		}
		log.Fatal(err)
	} else {
		p("File exists.\n")
	}
	err = os.Rename("info.txt", "data.txt")
	if err != nil {
		log.Fatal("Error renaming:", err)
	} else {
		p("Renamed.\n")
	}

	/*
		Remove the file created at exercise #1.
		Take care that the file is now called data.txt (it has been renamed at exercise #2).
	*/
	err = os.Remove("data.txt")
	if err != nil {
		log.Fatal("Error while delete:", err)
	} else {
		p("Deleted data.txt.\n")
	}

	/*
		Create a Go Program that reads the entire contents of a file called info.txt into a string.
		You can use ioutil.ReadAll() or any other function you want.
		The file exists in the current working directory.
	*/
	file, err = os.Open("info1.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("Error:", err)
	}
	data, err := ioutil.ReadAll(file)

	if err != nil {
		p("%v\n", err)
	} else {
		p("%s\n", data)
	}

	/*
		Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.
		The file exists in the current working directory.
	*/
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		p("%v\n", scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	/*
		1. Using single function creates a file called info.txt in the current directory. If the file already exists it will truncate it to zero size.
		2. Write the string "The Go gopher is an iconic mascot!" to the file.
	*/
	file, err = os.OpenFile("info.txt", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	_, err = io.WriteString(file, "The Go gopher is an iconic mascot!")
	if err != nil {
		log.Fatal(err)
	}

}
func structExercise() {
	/*
		1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.
		2. Declare and initialize two values of type person, one called me and another called you.
		3. Print out the struct values.
	*/
	type grades struct {
		grade  string
		course int
	}
	type person struct {
		name      string
		age       int
		favColors []string
		grade     grades
	}
	me := person{name: "Logu", age: 10, grade: grades{grade: "Golang", course: 98}}
	you := person{name: "You", favColors: []string{"green", "blue"}}

	p("Me: %+v\nYou: %+v\n", me, you)

	/*
		1. Change the name or the struct value called me to "Andrei"
		2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
		3. Add a new favorite color to the second struct value called you.
		4. Print out the struct values.
	*/
	me.name = "Andrei"
	you.name = "Logu"
	p("%T\n", you.favColors)
	you.favColors = append(you.favColors, "yellow")

	p("Me: %+v\nYou: %+v\n", me, you)

	me.favColors = append(me.favColors, "pink", "sky blue", "violet", "white")
	/*
		Iterate and print out the favorite colors of the struct value called me.
	*/

	for _, color := range me.favColors {
		p("%s\n", color)
	}

	/*Change the code from Exercise #1 and:

	1. Create a new struct type called grades with  2 fields: grade and course

	2. Add another field of type grades to person struct type (embedded struct).

	3. Change the initialization of the struct values called me and you to contain information about the grades.

	4. Change the grade and the course of one struct value to "Golang" and 98.

	5. Print out the struct values.*/

}
func cube(x float64) float64 {
	return math.Pow(x, 3)
}
func f1(x uint) (uint, uint) {
	var fact uint = 1
	var sum uint = 0
	var i uint
	for i = 1; i <= x; i++ {
		fact *= i
		sum += i
	}
	return fact, sum

}
func myFunc(s string) int {
	is, _ := strconv.Atoi(s)
	iss, _ := strconv.Atoi(strings.Repeat(s, 2))
	isss, _ := strconv.Atoi(strings.Repeat(s, 3))
	return is + iss + isss
}

func sum(nums ...int) (sum int) {
	for _, i := range nums {
		sum += i
	}
	return
}

func searchItem(items []string, item string) bool {
	for _, v := range items {
		if strings.EqualFold(v, item) { // v == item {
			return true
		}
	}
	return false
}

func funcExercise() {
	/*
	   Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3).
	*/
	p("Cube: %v\n", cube(2.))

	/*
		Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:
			a) the factorial of n
			b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)
		Test the program by calling the function.
	*/
	a, b := f1(3)
	p("Fact: %d Sum: %d\n", a, b)
	a, b = f1(4)
	p("Fact: %d Sum: %d\n", a, b)

	/*
		Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes (this is in fact a string). If the argument is integer 'n', the function should return the result of n + nn + nn
		Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107
	*/
	p("myFunc: %d\n", myFunc("5"))
	/*
		Create a function with the identifier sum that takes in a variadic parameter of type int and returns the sum of all values of type int passed in.
	*/
	p("%v\n", sum(1, 2, 3, 4))
	/*
		Change the function from the previous exercise and use a `naked return`.
	*/
	/*
			Create a function called searchItem() that takes 2 parameters: a) a string slice and b) a string.
		The function should search for the string (the second parameter) in the slice (the first parameter) and returns true if it finds the string in the slice and false otherwise. Do function does an case-sensitive search.
	*/
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false

	/*
		Create a function that takes in an int value and prints out that value.
		Assign the function to a variable, print out the type of the variable and then call that function through the variable name.
	*/
	f := func(a int) { p("anomy: %v\n", a) }

	p("%T\n", f)
	f(1)
}
func swap(x *int, y *int) {
	*x, *y = *y, *x
}
func pointerExercise() {
	/*
		Consider the following variable declaration x := 10.10
		1. Print out the address of x
		2. Declare a pointer called ptr that stores the address of x.
		3. Print out the type and the value of ptr.
		4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
	*/

	x := 1
	ptr := &x
	p("ptr type: %T value: %v\n", ptr, ptr)
	p("addr ptr: %v deref: %v\n", &ptr, *ptr)

	x, y := 10, 2
	ptrx, ptry := &x, &y

	/* Declare a new variable called z and initialize the variable by dividing x by y through the pointers. */
	z := *ptrx / *ptry
	p("z: %v\n", z)

	/*
		Consider the following variable declaration:x, y := 5.5, 8.8
		Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5
	*/
	p("Before x:%v y:%v\n", x, y)
	swap(&x, &y)
	p("After x:%v y:%v\n", x, y)

}

type money float64

func (m money) print() {
	p("%.2f\n", m)
}
func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price -= b.price * 0.1
}
func (b *book) changePrice(p float64) {
	b.price = p
}

func oopsExercise() {
	/*
		1. Create a new type called money. Its underlying type is float64.
		2. Create a method called print that prints out the money value with exact 2 decimal points.
	*/
	var rupee money = 5.5566

	rupee.print()

	/*
		Consider the money type declared at exercise #1. Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.
	*/
	p("%T\n", rupee.printStr())

	/*
		1. Create a new struct type called book with 2 fields: price and title of type float64 and string.
		2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.
		3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.
	*/

	var b = book{price: 100., title: "Golang book"}

	p("%#v\n", b)
	p("Book vat %v\n", b.vat())
	p("Before %#v\n", b)
	b.discount()
	p("After discount %#v\n", b)
	// changing the price
	b.changePrice(11.99)
	p("New price %#v\n", b)

}

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}
func (c car) Name() string {
	return c.brand
}

type cube1 struct {
	edge float64
}

func volume(c cube1) float64 {
	return c.edge * c.edge * c.edge
}
func interfaceExercise() {
	/*
		1. Create a Go program where car type implements the vehicle interface.
		2. Create a variable of type vehicle that holds a car struct value.
		3. Call the methods (Licence and Name) on the interface value declared at step 2
		4. Run the program without errors.
	*/
	c := car{licenceNo: "TN33 2338", brand: "Maruti"}
	var v vehicle = c

	p("%+v\n", v)

	/*
		1. Declare a variable called empty of type empty interface. Print out its type.
		2. Assign an int value to the variable called empty. Print out its type.
		3. Assign a float64 value to empty. Print out its type.
		4. Assign an int slice value to empty. Print out its type.
		5. Add a new int value to the slice (empty variable).
		6. Print out the slice (empty variable).
	*/
	var empty interface{}

	p("%T %v\n", empty, empty)

	empty = 1
	p("%T %v\n", empty, empty)
	empty = 1.
	p("%T %v\n", empty, empty)
	empty = []int{1, 2, 3}
	p("%T %v\n", empty, empty)
	empty = append(empty.([]int), 4)
	p("%T %v\n", empty, empty)

	var x interface{}
	x = cube1{edge: 5}
	v1 := volume(x.(cube1))
	fmt.Printf("Cube Volume: %v\n", v1)
}
