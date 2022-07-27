package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	strs := []string{"Jiten", "Rahim", "Bob", "something", "abcd", "", "Hello World", "Hello Go", "Go lang implementation", "Corban language", "java"}
	//	go func() {
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	//wg.Add(len(strs))
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		for i, _ := range strs {
			if i == 500 {
				panic("Things went wrong")
			}
			wg.Add(1)
			go func(wg *sync.WaitGroup, i int) {
				defer wg.Done()
				fmt.Println(i, "--->", reverse(strs[i]))
			}(wg, i)
		}
	}()
	fmt.Println("End of main")
	wg.Add(1)
	go func() {
		//wg.Add(1)
		defer wg.Done()
		go func() {
			time.Sleep(time.Second * 1)
			//	wg.Done()
		}()
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				count++
			}()
		}
	}()
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		wg.Done()
	}

	fmt.Println("Count", count)
}

func reverse(str string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	rstr := ""
	if str == "" {
		panic("wrong string")
	}
	for _, v := range str {
		rstr = string(v) + rstr
	}
	return rstr
}

// string
// splitit into liness
// lines:=strings.split(str,"\n")
// each go routine takes a line and count number of words
