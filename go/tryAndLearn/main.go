package main

import (
	"fmt"
	"net/http"
	"reflect"
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

// type Crequest struct {
// 	Method string
// 	Body   *bytes.Reader
// }
type HttpRequest struct {
	Name string
	Req  *http.Request
}

func main() {
	p := person{name: "Logu"}
	p1 := person1(p)

	fmt.Printf("%#v %#v\n", p, p1)

	fmt.Printf("BNum: %q \n", model.BuildNumber)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost", nil)
	httpReq := &HttpRequest{
		Name: "POST",
		Req:  req,
	}

	req1, _ := http.NewRequest(http.MethodPost, "http://localhost", nil)
	httpReq1 := &HttpRequest{
		Name: "POST",
		Req:  req1,
	}
	fmt.Printf("\n%v %v\n", httpReq, httpReq1)
	fmt.Printf("\n%v\n", reflect.DeepEqual(httpReq.Req, httpReq1.Req))
	fmt.Printf("\n%v\n", reflect.DeepEqual(httpReq, httpReq1))
	fmt.Printf("\n%v\n%v\n", httpReq.Req, httpReq1.Req)
}
