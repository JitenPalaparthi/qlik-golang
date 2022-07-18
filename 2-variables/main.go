package main

import (
	"fmt"
	"reflect"
)

// Type inference .. based on the type of the variable the compiler infers a value
// For numerics --> int, float etc.. it is 0
// For boolean --> false
// For string --> ""
// for byte --> 0 ; byte is equal to int8
// for rune --> 0 ; rune is equal to int32
func main() {

	var num int
	var ok bool

	var num1 int8  // -127 to +126
	var num2 uint8 //0-255

	var (
		num3 int16
		num4 int32
		num5 int64
		num6 uint16
		num7 uint32
		num8 uint64
	)

	var (
		float1 float32
		float2 float64
	)

	var (
		rune1 rune
		byte1 byte
	)

	var str string

	fmt.Printf("Value and Type of num %v - %T\n", num, num)
	fmt.Printf("Value and Type of num %v - %T\n", ok, ok)
	fmt.Println("num1:", num1, "type:", reflect.TypeOf(num1))
	fmt.Println("num2:", num2, "type:", reflect.TypeOf(num2))
	fmt.Println("num3:", num3, "type:", reflect.TypeOf(num3))
	fmt.Println("num4:", num4, "type:", reflect.TypeOf(num4))
	fmt.Println("num5:", num5, "type:", reflect.TypeOf(num5))
	fmt.Println("num6:", num6, "type:", reflect.TypeOf(num6))
	fmt.Println("num7:", num7, "type:", reflect.TypeOf(num7))
	fmt.Println("num8:", num8, "type:", reflect.TypeOf(num8))
	fmt.Println("float1:", float1, "type:", reflect.TypeOf(float1))
	fmt.Println("float2:", float2, "type:", reflect.TypeOf(float2))
	fmt.Println("rune1:", rune1, "type:", reflect.TypeOf(rune1))
	fmt.Println("byte1:", byte1, "type:", reflect.TypeOf(byte1))
	fmt.Println("str:", str, "type:", reflect.TypeOf(str))

	num1 = 100
	rune1 = 65
	byte1 = 65
	float1 = 12.25
	float2 = 25.25
	str = "Hello World!"

	// short-hand declaration

	age := 100 // int
	// var age uint8
	float3 := 12.4567 // float64

	fmt.Println("age", age, "type", reflect.TypeOf(age))
	fmt.Println("float3", float3, "type", reflect.TypeOf(float3))
	fmt.Printf("float3 with 2 decimals %0.2f\n", float3)

	// complex is a builtin function
	// there are twi variations of complex numbers
	// 1- complex64
	// 2- complex128
	c1 := complex(12, 45.56)
	fmt.Println("c1:", c1, "type:", reflect.TypeOf(c1))

	var r1, i1 float32 = 12, 45.56

	c2 := complex(r1, i1)
	fmt.Println("c2:", c2, "type:", reflect.TypeOf(c2))

	c3 := 16 + 45.56i // instead of using complex builtin function
	fmt.Println("c3:", c3, "type:", reflect.TypeOf(c3))

	// arthemetical operations

	fmt.Println("Addition of two complex numbers c1+c3", c1+c3)
	fmt.Println("Substract of two complex numbers c1-c3", c1-c3)
	fmt.Println("Multiplication of two complex numbers c1*c3", c1*c3)
	fmt.Println("Division of two complex numbers c1/c3", c1/c3)
	fmt.Println("Real part of c1", real(c1))
	fmt.Println("imaginary part of c1", imag(c1))
}
