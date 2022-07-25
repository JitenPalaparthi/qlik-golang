package main

import "fmt"

// defer is a keyword.
// defer defers the execution to the end of the caller
// all defer calles are stored in FILO form

func main() {
	greet("Qlik")
	defer fmt.Println("Hello World-1")
	defer fmt.Println("Hello World-2")
	fmt.Println("Welcome to the session")

	v := new(int)
	*v = 100

	num := 100

	ptr := &num
	defer fmt.Println("Value of V and num  with defer:", *v+100, *ptr+100, num+100)

	*v = *v + 300
	num = num + 300 // *ptr = *ptr+300

	fmt.Println("Value of V, ptr and num:", *v, *ptr, num)

	str := "Hello World!"
	for _, v := range str {
		defer print(string(v)) // !dlroW olleH
	}
	defer println()

}

func greet(name string) {
	defer fmt.Println("Greet ending")
	fmt.Println("Greet starting")
	fmt.Println("Hello " + name)
}

// defer
