package main

import (
	"errors"
	"fmt"
)

type MyInt int // user defined type

func (mi MyInt) ToString() string {
	return fmt.Sprint(mi)
}

var (
	ErrNilMap = errors.New("nil map")
)

type MyMap map[string]interface{}

func (mmp MyMap) GetKeys() (keys []string, err error) {

	if mmp == nil {
		//return nil, errors.New("nil map")
		// return nil, fmt.Errorf("nil map")
		return nil, ErrNilMap
	}

	keys = make([]string, len(mmp))

	index := 0
	for key, _ := range mmp {
		keys[index] = key
		index++
	}

	return keys, nil
}

func (mmp MyMap) GetValues() (values []interface{}, err error) {

	if mmp == nil {
		//return nil, errors.New("nil map")
		// return nil, fmt.Errorf("nil map")
		return nil, ErrNilMap
	}

	values = make([]interface{}, len(mmp))

	index := 0
	for _, value := range mmp {
		values[index] = value
		index++
	}

	return values, nil
}
func main() {
	var num1 MyInt = 100
	fmt.Println("Calling ToString method:", num1.ToString())

	var num2 int = 300

	fmt.Println("Calling ToString Method", MyInt(num2).ToString())

	var mmp MyMap

	mmp = make(map[string]interface{})
	mmp["name"] = "Jiten"
	mmp["age"] = 37
	mmp["isMarried"] = true
	if keys, err := mmp.GetKeys(); err == nil {
		fmt.Println("Keys:", keys)
	} else {
		fmt.Println(err)
	}

	if values, err := mmp.GetValues(); err == nil {
		fmt.Println("Values:", values)
	} else {
		fmt.Println(err)
	}

	var mp map[string]interface{}

	mp = make(map[string]interface{})
	mp["name"] = "Jiten"
	mp["age"] = 37
	mp["isMarried"] = true

	if keys, err := MyMap(mp).GetKeys(); err == nil {
		fmt.Println("Keys:", keys)
	} else {
		fmt.Println(err)
	}

	if values, err := MyMap(mp).GetValues(); err == nil {
		fmt.Println("Values:", values)
	} else {
		fmt.Println(err)
	}

	mp1 := make(map[string]interface{})
	mp1["name"] = "Jiten"
	mp1["age"] = "37"
	mp1["isMarried"] = "true"

	fmt.Println(MyMap(mp1).GetKeys())

}
