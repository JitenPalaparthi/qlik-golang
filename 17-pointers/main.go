package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {

	var num1 int = 100
	var ptr1 *int

	if ptr1 == nil {
		fmt.Println("ptr1 is nil")
	} else {
		*ptr1 = 301 // can never assign a value to a nil pointer
	}
	ptr1 = &num1

	fmt.Println("Address of num1 Value of num1", &num1, num1)
	fmt.Println("Address at ptr1 Value at ptr1", ptr1, *ptr1)
	*ptr1 = 200
	fmt.Println("Address of num1 Value of num1", &num1, num1)
	fmt.Println("Address at ptr1 Value at ptr1", ptr1, *ptr1)
	changeVal(&num1, 300)
	fmt.Println("Address of num1 Value of num1", &num1, num1)
	fmt.Println("Address at ptr1 Value at ptr1", ptr1, *ptr1)
	changeVal(ptr1, 400)
	fmt.Println("Address of num1 Value of num1", &num1, num1)
	fmt.Println("Address at ptr1 Value at ptr1", ptr1, *ptr1)

	//var ptr2 *int

	ptr3 := new(int) // new is a builtin function. Returns a pointer and does type inference
	fmt.Println("Pointer and value", ptr3, *ptr3)

	var ptr4 **int
	ptr4 = &ptr1

	fmt.Println("Value of num1 and address of num1", **ptr4, *ptr4)

	arr := [3]int{10, 3, 16}
	fmt.Println("Address of array elements", &arr[0], &arr[1], &arr[2])
	slice := []int{10, 3, 16}
	fmt.Println("Address of slice elements", &slice[0], &slice[1], &slice[2])

	fmt.Println("Size of T1 and T2", unsafe.Sizeof(T1{}), unsafe.Sizeof(T2{}))
	arrptr := uintptr(unsafe.Pointer(&arr[0]))

	arrptr = arrptr + unsafe.Sizeof(arr[0]) // Arthemetic operation on uintptr

	val := (*int)(unsafe.Pointer(arrptr)) //
	fmt.Println("Array second element is ", *val)
	//824633819592 + 8 = 824633819600

	arr1 := [3]*int{&arr[0], new(int), &arr[2]}
	fmt.Println(arr1)

	*arr1[0] = 101
	fmt.Println("Pointer array", arr1)
	fmt.Println("normal array", arr)

	iarr := []interface{}{"Jiten", &arr[0], T1{}, true}
	for _, v := range iarr {
		switch v.(type) {
		case string:
			fmt.Println(v.(string))
		case *int:
			fmt.Println(v.(*int))
		case T1:
			fmt.Println(v.(T1))
		default:
			fmt.Println(v)
		}
	}
}

func changeVal(num *int, val int) error {
	if num != nil {
		*num = val
	}
	return errors.New("nil pointer")
}

type T1 struct {
	ok  bool   // 1 -->8 -->7
	no  int    // 8 -->8 -->0
	c   rune   // 4 -->8 -->4
	str string // 16-->16-->0
	ok1 bool   // 1 -->8--> 7
}

type T2 struct {
	str string // 16 - 0
	num int    // 8  - 0
	c   rune   // 4  - 4
	ok  bool   // 1 -  3
	ok1 bool   // 1 -  3
}

// defer , panic and recover
// go routines, waitGroup, Mutex, Go Channels, unbuffered, buffered, send only, receive only Close, Range, deadlocks
// escape analysis , profiler
// context --> sending values, parent context, cancel context
// restful, grpc
// gorm, mongodb
// testings
// nats | nsq | kafka | rabbitmq
// dlv
