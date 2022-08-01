package main

// Buffered channel does not block sender until the buffer is full
func main() {
	ch := make(chan int, 1) // buffered channel
	ch <- 100
	val := <-ch
	ch <- val * 10
	println(<-ch)
}

// bufferes
// pod 128mb , 1/10 processor
// Based on what you give buffer size?
// 1-
// 2-
