package main

import (
	"fmt"
	"tryAndLearn/model"
)

type person struct {
	name   string   `person_name`
	colors []string `fav_color`
}

type person1 struct {
	name   string   `person_name`
	colors []string `like_color`
}

func main() {
	p := person{name: "Logu"}
	p1 := person1(p)

	fmt.Printf("%#v %#v\n", p, p1)

	fmt.Printf("BNum: %q \n", model.BuildNumber)
}
