package main

func main() {
	if age := 17; age >= 18 {
		println("eligible for vote.Because age is ", age)
		//return
	} else {
		println("not eligible for vote. Because age is", age)
	}

	num := 55

	if num <= 50 {
		println("num is  less than 50")
		//} else if (num > 50 && num <= 100) && (num%10 == 0 || num%5 == 0) {
	} else if num > 50 && num <= 100 {
		if num%10 == 0 || num%5 == 0 {
			println("num is between 51-100 and divisible by 5 or 10")
		} else {
			println("num is between 51-100")
		}
	} else if num > 100 && num <= 200 {
		println("num is between 101-200")
	} else {
		println("num is greater than 200")
	}

	// && || !=
}
