package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	//pctx := context.Background()
	ch := make(chan int)
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*10))
	//	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Num of CPU,GORoutines running", runtime.NumCPU(), runtime.NumGoroutine())
	//go square(ctx, ch)
	go square(ctx, ch)

	// for n := 1; n <= 20; n++ {
	// 	fmt.Println("Received Squre of n-->", <-ch)
	// }
	go func(context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("I am done receiving values. Returning now")
				return
			default:
				fmt.Println("Received Squre of n-->", <-ch)
			}
		}
	}(ctx)

	fmt.Println("Num of CPU,GORoutines running", runtime.NumCPU(), runtime.NumGoroutine())
	//
	//defer cancel()
	time.Sleep(time.Second * 2)
	//cancel()
	//time.Sleep(time.Millisecond * 2)
	fmt.Println("Num of CPU,GORoutines running", runtime.NumCPU(), runtime.NumGoroutine())
}

func square(ctx context.Context, out chan int) {
	i := 0
	for {
		select {
		case <-ctx.Done(): // even the channel is closed select case receives it
			fmt.Println(ctx.Err())
			return
		case out <- i * i:
			i++
		}
	}
}
