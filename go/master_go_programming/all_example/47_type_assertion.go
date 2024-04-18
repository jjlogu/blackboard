/////////////////////////////////
// Interfaces - Type Assertions and Type Switches
// Go Playground: https://play.golang.org/p/0A07wpGtXYa
/////////////////////////////////

// A type assertion provides access to an interface’s concrete value.

package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
type shape interface {
	area() float64
	perimeter() float64
}

// declaring a struct type
type rectangle struct {
	width, height float64
}

// declaring a struct type
type circle struct {
	radius float64
}

// declaring a method for circle type
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// declaring a method for circle type
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// declaring a method for circle type
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// declaring a method for rectangle type
func (r rectangle) area() float64 {
	return r.height * r.width
}

// declaring a method for rectangle type
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// declaring a function that takes an interface value
func print(s shape) {

	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {

	// declaring an interface value that holds a circle type value
	var s shape = circle{radius: 2.5}

	fmt.Printf("%T\n", s) //interface dynamic type is circle

	// no direct access to interface's dynamic values
	// s.volume() -> error

	// there is access only to the methods that are defined inside the interface
	fmt.Printf("Circle Area:%v\n", s.area())

	// an interface value hides its dynamic value.
	// use type assertion to extract and return the dynamic value of the interface value.
	fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// checking if the assertion succeded or not
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//

	// it permits several type assertions in series
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}

	/**

	# Refer: https://go.dev/tour/methods/15
	A type assertion provides access to an interface value's underlying concrete value.

	t := i.(T)
	This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

	If i does not hold a T, the statement will trigger a panic.

	To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

	t, ok := i.(T)
	If i holds a T, then t will be the underlying value and ok will be true.

	If not, ok will be false and t will be the zero value of type T, and no panic occurs.

	Note the similarity between this syntax and that of reading from a map.


	Refer: https://go.dev/tour/methods/16
	Type switches
	A type switch is a construct that permits several type assertions in series.

	A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

	switch v := i.(type) {
	case T:
		// here v has type T
	case S:
		// here v has type S
	default:
		// no match; here v has the same type as i
	}
	The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

	This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.

	**/

}
