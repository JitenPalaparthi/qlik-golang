package main

func main() {
	num := 100
	println(num)
	c := add1(10, 30)
	println(c)
	str := concat1("Hello", "World")
	println(str)
	concat2("Hello", "World")

	str1 := new(string)
	str2 := new(string)
	println("Initial addresses-->", str1, str2)
	*str1 = "Hello"
	*str2 = "World"
	//concat3(str1, str2)
	//println("After concatination", *str1)
	println("Addressess after assigning values-->", str1, str2)
	concat3(str1, str2)
	println("Addressess after first concat3 call-->", str1, str2)
	concat3(str1, str2)
	println("Addressess after second concat3 call-->", str1, str2)
	a1 := new(int)
	b1 := new(int)
	*a1 = 10
	*b1 = 20
	add2(a1, b1)

}
func add1(a, b int) int {
	return a + b
}

func add2(a, b *int) {
	*a = *a + *b
}

func add3(a, b *int) *int {
	c := new(int) // this escapes to heap
	d := new(int) // this does not
	println(d)
	*c = *a + *b
	return c
}

func concat1(str1, str2 string) string {
	str1 = str1 + str2
	return str1
	//return "Hello World"
}

func concat2(str1, str2 string) {
	str := str1 + str2
	println(str)
}

func concat3(str1, str2 *string) {
	*str1 = *str1 + *str2
}

// go run -gcflags '-m -l' main.go
