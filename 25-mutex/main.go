package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100000000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer mu.Unlock()
				if !mu.TryLock() {
					fmt.Println("True")
					mu.Lock()
				}
				count++
			}()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer mu.Unlock()
				mu.Lock()
				count++
			}()
		}
	}()
	//time.Sleep(time.Second * 1)
	wg.Wait()
	fmt.Println(count)
}
