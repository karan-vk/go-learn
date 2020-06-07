package main

import "fmt"

type example struct {
	flag    bool
	counter int8
	pi      float32
}
type bill struct {
	flag    bool
	counter int8
	pi      float32
}

type alice struct {
	flag    bool
	counter int8
	pi      float32
}

func main() {
	var e1 example
	fmt.Printf("%+v\n", e1)
	e2 := example{
		flag:    true,
		counter: 25,
		pi:      3.141,
	}
	e2.counter = 2
	fmt.Printf("%+v \n", e2)
	var b bill
	var a alice
	b = bill(e2)
	fmt.Println(b, a)

}
