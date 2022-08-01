package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Addition", addition())
	fmt.Println("Addition", addition(10))
	fmt.Println("Addition", addition(10, 20))
	fmt.Println("Addition", addition(10, 20, 1, 3, 3, 43, 54, 98))
	slice := make([]int, 10)
	var arr [10]int
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(1000)
		arr[i] = rand.Intn(1000)
	}
	fmt.Println("Addition", addition(slice...))
	fmt.Println("Addition", addition(arr[:]...))
	fmt.Println("Addition", additionI(10, 20, 30, 40))
	fmt.Printf("Addition %0.2f", additionI(10.45, 20.54, 30.65, 40.87))
	fmt.Printf("Addition %0.2f\n", additionG(10.45, 20.54, 30.65, 40.87))
	fmt.Println("Addition", additionG(10, 20, 30, 40))

}
