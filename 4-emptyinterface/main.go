package main

import "fmt"

func main() {

	var val interface{} // empty interface. It can hold any kind of data

	val = 100
	// type casting does not work when assign value from interface{} to underlining type
	// var num int = val // not ok
	// use type assertion rather than type casting when converting from interface{} to the type
	// v.(T) --> emptyInterface.(Type)
	var num int = val.(int) // ok

	val = 12.45

	var float float32 = float32(val.(float64))

	val = false

	var boolean bool = val.(bool)

	val = "Hello World"

	var str string = val.(string)
	fmt.Println(num, float, boolean, str)

	// swapping

	a, b := 10, 20
	// t := a
	// a = b
	// b = t
	c := 30
	fmt.Println("Before swap--> a,b,c", a, b, c)
	a, b, c = c, a, b
	fmt.Println("After swap--> a,b,c", a, b, c)

}

// java: Object, csharp: Object
