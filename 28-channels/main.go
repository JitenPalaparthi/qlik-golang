package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for i := 0; i <= 1000; i++ {
			wg.Add(1)
			go sender(i, wg, ch)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	//}()
	for val := range ch {
		fmt.Println(val)
	}

	// for {
	// 	val, ok := <-ch
	// 	if !ok {
	// 		break
	// 	} else {
	// 		fmt.Println(val)
	// 	}
	// }
	//wg.Done()
}

func sender(val int, wg *sync.WaitGroup, ch chan<- int) {
	fmt.Println("Send:", val)
	ch <- val
	wg.Done()
}

// func receiver(ch <-chan int) {
// 	fmt.Println("Receive:", <-ch)
// }
