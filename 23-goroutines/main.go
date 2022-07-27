package main

import "time"

// scheduling is handled by go runtime..

// Apache server / Nginx / DB server

// Thread Pool ?

// 1- main is also a go routine
// 2- by default,no go routine waits for other goroutine to complete its execution
// 3- order of execution of go routines are concurrent
func main() {
	go func() {
		for i := 1; i <= 10000; i++ {
			go println("Hello Anonymous function-->", i)
		}
	}()
	time.Sleep(time.Second * 10)
	println("Hello World")
}

// 10 --> 10 threads
// 100 --> 10 threads
// 1000000 -> 10 threads

//tinyGo ->llvm --> MicroControllers
