package main

import (
	"fmt"
	"udt/employee"
)

func main() {

	// ---- another example

	e1 := employee.Employee{ID: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: employee.Address{Line1: "TempleBells , Flat No 101", Street: "Mahalakshmi", City: "Bangalore", State: "Karnataka", PinCode: "560086", Country: "India"}}
	e1.ShowDetails()

	fmt.Println("Address line1:", e1.Line1) // can be caled directly even Line1 is a filed of embedded struct

	p1 := employee.NewPerson(100, "Jiten", true)
	p1.ShowDetails()

	var e employee.Empty
	e.Greet()

	var e2 employee.Employee // type inference gives default values
	fmt.Println(e2)

	p2 := Person{int: 101, string: "jiten palaparthi", bool: false}
	fmt.Println(p2)
}

type Person struct {
	int
	string
	bool
}

func (p *Person) ShowDetails() {
	fmt.Println("ID:", p.int)
	fmt.Println("Name:", p.string)
	fmt.Println("IsMale:", p.bool)
}
