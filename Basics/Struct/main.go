package main

import "fmt"
type contactInfo struct{
	email string
	zipCode int

}

type person struct{
	Firstname string
	Lastname string
	age int8
	contact contactInfo

}

func main()  {
	var Karan person
	Karan.age=18
	Karan.Firstname="Karan"
	Karan.Lastname="Narayanan"
	Karan.contact.email="karanvknarayanan@gmail.com"
	Karan.contact.zipCode=600064
	Karan.print()
	fmt.Println()
	Karan.updateName("Jay")
	// Karan.print()
	
	
}
func (p person)updateName(new string){
	p.Firstname=new
	p.print()
}

func (p person)print()  {
	
	fmt.Printf("%+v 😊",p)
	
}