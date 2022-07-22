package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// make is a builtin function
// make is used to create slice, map and channel
// nil -> can check whether a variable is nil or not on slice, map, channel, any pointer type
// 	and interfaces
func main() {

	var slice1 []int // declaring a slice

	slice1 = make([]int, 5) // type inference takes place
	fmt.Println("Length of the slice and capacity of the slice:", len(slice1), cap(slice1))

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(1000)
	}

	slice1 = append(slice1, rand.Intn(1000))
	slice1 = append(slice1, rand.Intn(1000))
	slice1 = append(slice1, rand.Intn(1000))
	// slice1 = append(slice1, rand.Intn(1000))
	// slice1 = append(slice1, rand.Intn(1000))
	// slice1 = append(slice1, rand.Intn(1000))
	fmt.Println("Length of the slice and capacity of the slice:", len(slice1), cap(slice1))
	fmt.Println("values", slice1, "Type of the slice", reflect.TypeOf(slice1))
	//slice1 := []int{10, 20}

	var slice2 []int
	//slice2 = make([]int, 0)
	if slice2 == nil {
		println("slice2 is nil")
	}
	// shorthand declaration of slice
	slice3 := []int{10, 11, 12, 13, 14} // this is not an array ; it is a slice
	//slice3 := []int{}
	//slice3:=[...]int{10, 11, 12, 13, 14} // this is an array
	//slice3:=[5]int{10, 11, 12, 13, 14} // this is an array
	fmt.Println("Values:Length:capacity :", slice3, len(slice3), cap(slice3))
	slice3 = append(slice3, 15)
	fmt.Println("Values:Length:capacity :", slice3, len(slice3), cap(slice3))

	slice4 := slice3 // this is a reference copy
	fmt.Println("Calling changeOfSlice function-1")
	//var err error
	if err := changeValOfSlice(slice4, 0, 1000); err != nil {
		println(err.Error())
	}
	fmt.Println("Slice4:--> Values:Length:capacity :", slice4, len(slice4), cap(slice4))
	fmt.Println("Slice3:--> Values:Length:capacity :", slice3, len(slice3), cap(slice3))
	fmt.Println("Calling changeOfSlice function-2")

	if err := changeValOfSlice(slice4, 10, 1110); err != nil {
		println(err.Error())
	}

	slice5 := make([]int, len(slice4))
	// var slice5 []int // not ok; unless instantiated cannot be used in copy function
	copy(slice5, slice4)
	slice5[0] = 1000

	fmt.Println("Slice5:--> Values:Length:capacity :", slice5, len(slice5), cap(slice5))
	fmt.Println("Slice4:--> Values:Length:capacity :", slice4, len(slice4), cap(slice4))
	fmt.Println("Slice3:--> Values:Length:capacity :", slice3, len(slice3), cap(slice3))

	fmt.Println("Sub slice")

	slice6 := make([]int, 10) // type inference takes place
	fmt.Println("Length of the slice and capacity of the slice:", len(slice1), cap(slice1))

	for i := 0; i < len(slice6); i++ {
		slice6[i] = rand.Intn(1000)
	}

	fmt.Println("Slice6", slice6)
	fmt.Println("Slice6 all elements", slice6[:])
	fmt.Println("Slice6 from 0th to 5th", slice6[:5])
	fmt.Println("Slice6 from 5th to 8th", slice6[5:8])
	fmt.Println("Slice6 from 5th to the end ", slice6[5:])

	arr := [5]int{1, 2, 3, 4, 5}
	slice7 := arr[:]                 // array to slice
	changeValOfSlice(arr[:], 0, 100) // passing arr as an argument to the slice
	fmt.Println("Slice7", slice7)
	fmt.Println("Array", arr)

}

// 1- make([]int,5,10)
// 2- append(slice,10,20,30,40,50,60,70,80,90,100,110,120)

// ptr:0x1130
// length:11
// capacity:20

func changeValOfSlice(slice []int, index uint, value int) error {

	if index >= uint(len(slice)) {
		//return errors.New("slice index out of bounds")
		return fmt.Errorf("slice index out of bounds")
	}
	slice[index] = value
	return nil
}
