package main

import "encoding/json"

func main() {
	t1 := T1{num1: 10}
	println(t1.num1)

	t2 := &T1{num1: 10}
	println(t2.num1)

	t3 := new(T1)
	println(t3.num1)

	changeVal(t3, 100)
	println(t3.num1)
	buf, _ := json.Marshal(t3)
	println(buf)
	json.Unmarshal(buf, t3)
	slice := make([]int, 1)
	slice = append(slice, 1000000)
	println(slice)
	changeSlice(slice)
	n1 := new(int64)
	n2 := new(int64)
	n3 := new(int64)
	n4 := new(int64)

	*n1 = 100
	*n2 = 200
	*n3 = 300
	*n4 = 400
	//recursive(n1, n2, n3, n4)
}

func changeVal(t *T1, val int) {
	t.num1 = val
}

func changeSlice(s []int) {
	println(s)
}

func recursive(n1 *int64, n2 *int64, n3 *int64, n4 *int64) {
	recursive(n1, n2, n3, n4)
}

type T1 struct {
	num1 int
}
