package main

import "fmt"

type user struct {
	name string
	age int
}


func main() {
	p:=user{
		name: "karan",
		age: 15,
	}
	fmt.Println(p)
	share(&p)
	fmt.Println(p)
}

func share(p *user)  {
	p.age=25
}