package model

import ("fmt"
)
var BuildNumber string


func init()  {
	fmt.Println("In init version.go")
	fmt.Println(BuildNumber)
}
