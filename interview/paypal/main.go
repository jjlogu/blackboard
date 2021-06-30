/*
temp = 10

def func():
      temp = 20
      print(temp)

print(temp)
func()
print(temp)

"""
10
20
10
"""



----------------------

package main

import (
    "fmt"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    // v{30,40}
    v.Abs(10)
    fmt.Println("%f, %f", v.X, v.Y)
    // 30, 40
}

-------
*/
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("indirect")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

/*

 direct 0
 direct 1
 direct 2
 indirect 0
 indirect 1
 indirect 2
 going
 done




1->3->2->4->5
1->2->5

# head =?

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

    def deleteEvenPos(head):
        previous = head
        pos = 1
        current = head
        while current.next:
            if pos%2 ==0:
               previous.next = current.next
            else:
               previous=current
            pos+=1
            current=current.next

*/
