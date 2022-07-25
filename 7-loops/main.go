package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	i := 1
	for ; i <= 10; i++ {
		print(i, " ")
	}
	println()
	j := 1
	for j <= 10 {
		print(j, " ")
		j++
	}
	println()
	k := 1

	for {
		if k > 10 {
			break
		}
		print(k, " ")
		k++
	}
	println()

	l := 1
loop:
	print(l, " ")
	l++
	if l >= 10 {
		goto loop
	}

	println("\ngoto loop finished")

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		print(i, "")
	}
	println()

outer:
	for i := 0; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == j {
				break outer
			}
			println(i, j)
		}
	}

	// range loop
	// array, slice, map or channel

	str := "Hello, 世界"

	for _, v := range str {
		fmt.Print(" ", string(v))
	}
	println()
	for i = 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}
}
