package main

import "fmt"

// Channel is a type.
// make to be used to create a channel
// can check whether the channel is nil
// sender is blocked until the value is received.
// receiver is blocked until it is received

// <- -> not ok
// ch <- 100 to send
// <- ch to receive
func main() {
	var ch chan int        // 8 bytes
	ch = make(chan int, 2) // unbuffered
	/*go func() {
		time.Sleep(time.Second * 1)
		println("Hello World-1")
		ch <- 101
	}()
	go func() {
		//time.Sleep(time.Second * 5)
		println("Hello World-2")
		ch <- 102
	}()
	val1 := <-ch
	println(val1)
	val2 := <-ch
	println(val2)*/
	done := make(chan struct{}, 4)
	go sender(101, ch)
	go sender(102, ch)
	go sender(103, ch)
	go receiveR(ch, done)
	go receiveR(ch, done)
	go receiveR(ch, done)
	go func() {
		for i := 1; i <= 1000; i++ {
			fmt.Print(i, " ")
		}
		done <- struct{}{}
	}()

	//go receiveR(ch, done)
	<-done
	<-done
	<-done
	<-done
}

func sender(val int, ch chan<- int) {
	fmt.Println("sender:", val)
	ch <- val
}

func receiveR(ch <-chan int, done chan<- struct{}) {
	fmt.Println("received:", <-ch)
	done <- struct{}{}
}

// func receive(ch <-chan int) {
// 	fmt.Println("received:", <-ch)
// }
