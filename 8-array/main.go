package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [3]int
	arr1[0] = 1001
	arr1[1] = 1002
	arr1[2] = 1003
	fmt.Println("before calling changeOfArray", arr1, "type", reflect.TypeOf(arr1))
	changeValOfArray(arr1, 0, 10)
	fmt.Println("before calling changeOfArray", arr1, "type", reflect.TypeOf(arr1))

	// arr1[3] = 1004 // not ok
	var arr2 [4]int

	fmt.Println("array", arr2, "type", reflect.TypeOf(arr2))

	arr3 := [...]int{100, 200, 300} // shorthand declaration
	fmt.Println("array", arr3, "type", reflect.TypeOf(arr3))

	arr4 := [3]int{100, 200} // shorthand declaration
	fmt.Println("array", arr4, "type", reflect.TypeOf(arr4))

	for i, v := range arr4 {
		fmt.Print(i, " ", v, " ")
	}
	println()
	for i := range arr4 {
		fmt.Print(i, " ")
	}
	println()
	for i := 0; i < len(arr4); i++ {
		print(arr4[i], " ")
	}
	println()

	arr5 := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Println("Length and Cap--->", len(arr5), cap(arr5))

	for i := 0; i < len(arr5); i++ {
		for j := 0; j < len(arr5[i]); j++ {
			for k := 0; k < len(arr5[i][j]); k++ {
				print(arr5[i][j][k], " ")
			}
		}
	}
}

func changeValOfArray(array [3]int, index uint, value int) error {
	if index >= uint(len(array)) {
		//return errors.New("slice index out of bounds")
		return fmt.Errorf("slice index out of bounds")
	}
	array[index] = value
	return nil
}
