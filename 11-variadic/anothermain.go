package main

func addition(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v // sum =sum +v
	}
	return sum
}

// To get it work call with all int arguments or all float64 argumanets
func additionI(vals ...interface{}) interface{} {
	var intsum int
	var floatsum float64
	for _, v := range vals {
		switch v.(type) {
		case int:
			intsum = intsum + v.(int)
		case float64:
			floatsum = floatsum + v.(float64)
		}
	}
	if intsum != 0 {
		return intsum
	}
	return floatsum
}

//func additionG[numeric int | float64](nums ...numeric) numeric {
func additionG[numeric Numeric](nums ...numeric) numeric {
	var sum numeric
	for _, v := range nums {
		sum += v
	}
	return sum
}

type Numeric interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | int | float32 | float64
}
