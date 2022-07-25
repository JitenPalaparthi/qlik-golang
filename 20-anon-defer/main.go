package main

import "fmt"

func main() {
	defer func(name string) {
		defer fmt.Println("Greet Ending")
		fmt.Println("Greet Starting")
		fmt.Println("Hello", name)
	}("Qlik!-1")

	grt := func(name string) {
		defer fmt.Println("Greet Ending")
		fmt.Println("Greet Starting")
		fmt.Println("Hello", name)
	}

	defer grt("Qlik!-2")

	defer greet("Qlik!-3", grt)

	fmt.Println("Hello World-1")
}

func greet(name string, fn func(string)) {
	fn(name)
}
