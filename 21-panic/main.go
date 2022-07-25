package main

import "fmt"

func main() {
	// fmt.Println("Divid by zero", 10/0)
	//v := 0
	defer fmt.Println("Hello World!")
	fn := new(string)
	*fn = "Jiten"

	func() {
		fullName(fn, nil)
		//defer fmt.Println("Divid by zero end")
		//fmt.Println("Divid by zero start")
		//fmt.Println("Divid by zero", 10/v)
	}()
}

func fullName(fn, ln *string) {
	defer println("end calling fullName")
	if fn == nil {
		panic("firstname is nil")
	}
	if ln == nil {
		panic("lastname is nil")
	}

	println(*fn + " " + *ln)
}
