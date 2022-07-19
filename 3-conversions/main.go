package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// there is no implicit conversion in go
// to convert from one type to other must use type casting syntax
// type casting is Type(value) --> T(v)
func main() {

	num1 := 19000 // This is of type int

	// implicitly cannot convert from int to int8
	var num2 int64 = int64(num1) //int64(num1)-->ok (int64)num1 -->not ok

	fmt.Println("num2", num2)

	var num3 uint8 = uint8(num1)
	fmt.Println("num3", num3)

	// convert to string
	var str1, str2 string
	str1 = string(num1)
	fmt.Println("Convert from rune to str", str1)

	// want to conver  19000 as "19000"

	// 1- option using fmt.Spring
	str2 = fmt.Sprint(num1)
	fmt.Println("Convert from int to as it is string", str2)

	var char rune = '䨸'
	var num4 int = int(char)
	fmt.Println("Rune to int", num4)

	// 2- option using strconv package
	num5, _ := strconv.Atoi("19000") // _ blank identifier
	fmt.Println("num5", num5, "type", reflect.TypeOf(num5))

	num6 := 19000
	str3 := strconv.Itoa(num6)
	fmt.Println("str3", str3, "type", reflect.TypeOf(str3))

}
