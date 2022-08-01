package main

func main() {
	// in switch break is automatically taken
	switch day := 5; day {
	case 1:
		println("Sunday")
	case 2:
		println("monday")
	case 3:
		println("tuesday")
	case 4:
		println("wednesday")
	case 5:
		println("thursday")
	case 6:
		println("friday")
	case 7:
		println("saturday")
	default:
		println("no day")
	}

	char := 'A'

	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		println(string(char), " is vovel")
	case 'A', 'E', 'I', 'O', 'U':
		println(string(char), " is vovel")
	default:
		println(string(char), "is either cosonent or any unicode char")
	}

	switch {
	case char >= 65 && char <= 90:
		println(string(char), " is upper case")
	case char >= 96 && char <= 122:
		println(string(char), " is lower case")
	default:
		println(string(char), "special char")
	}

	value := 80
	switch {
	case value%8 == 0:
		println(value, "is divisible by 8")
		fallthrough
	case value%4 == 0:
		println(value, "is divisible by 4")
		fallthrough
	case value%2 == 0:
		println(value, "is divisible by 2")
	}
}
