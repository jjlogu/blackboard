package model

import (
	"fmt"
)

func init() {
	fmt.Println("In init config.go")
	BuildNumber = "set at config.go"
}
