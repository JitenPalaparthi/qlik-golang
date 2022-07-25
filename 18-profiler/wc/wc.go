package wc

func WordCount(str string) (wc int, lc int) {

	//var si int
	hasChar := false

	for i, v := range str {
		if string(v) == "\n" {
			lc++
		}
		if string(v) != " " && string(v) != "\n" {
			hasChar = true
		}

		if (string(v) == " " || string(v) == "\n") && i != 0 && hasChar {
			wc++
			//fmt.Println(str[si:i])
			//si = i
			hasChar = false
		}
	}
	if hasChar {
		wc++
	}
	return wc, lc
}
