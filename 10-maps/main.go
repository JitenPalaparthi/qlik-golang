package main

import "fmt"

func main() {

	// == operation; Keys can be int, strings, float,bool
	var mp map[string]string // declaration of the map

	mp = make(map[string]string)

	mp["rajajinagar-1"] = "560086"
	mp["rajajinagar-2"] = "560076"
	mp["yeshvantpur"] = "560096"
	mp["guntur"] = "522001"

	fmt.Println(mp)

	// iterating the map

	for key, value := range mp {
		fmt.Println("Key", key, "value", value)
	}

	// if the key exists or not

	val, ok := mp["trivandrum"]
	if !ok {
		fmt.Println("there is no value with the given key")
	}

	val, ok = mp["rajajinagar"]
	if !ok {
		fmt.Println("there is no value with the given key")
	}
	fmt.Println(val, ok)

	if mp != nil {
		fmt.Println("Deleting an element from the map")
		delete(mp, "rajajinagar")
	}
	fmt.Println("Fetching keys and values from the map")
	for key, value := range mp {
		fmt.Println("Key", key, "value", value)
	}
	// rajajinagar --> hash --> 16bit --> 8bits | 8bits
}
