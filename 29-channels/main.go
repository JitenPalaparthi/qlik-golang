package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i <= 1000; i++ {
		go sender(i, ch)
	}
	count := 0
	for {
		if count == 1000 {
			close(ch)
			break
		}
		select {
		case val, ok := <-ch:
			fmt.Println(val, ok)
			go func() {
				done <- struct{}{}
			}()
		case <-done:
			count++
		}
	}
}

func sender(val int, ch chan<- int) {
	fmt.Println("Send:", val)
	ch <- val
}
