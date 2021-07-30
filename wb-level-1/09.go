package main

import (
	"context"
	"fmt"
	"time"
)

//	Написать конвейер чисел. Даны 2 канала - в первый пишутся числа из массива, во второй
//	пишется результат операции 2*x, после чего данные выводятся в stdout.

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	c, cancel := context.WithCancel(context.Background())

	// printer func
	go func(ctx context.Context) {
		for {
			select {
			case v := <-chOut:
				fmt.Println(v)
			case <-ctx.Done():
				fmt.Println("printer stopped")
				return
			}
		}
	}(c)

	// doubler func
	go func(ctx context.Context) {
		for {
			select {
			case v := <-chIn:
				chOut <- 2 * v
			case <-ctx.Done():
				fmt.Println("doubler stopped")
				return
			}
		}
	}(c)


	for i := 0; i < 100; i++ {
		chIn <- i
	}
	cancel()
	time.Sleep(200*time.Millisecond)
}
