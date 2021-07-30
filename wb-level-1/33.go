package main

import (
	"context"
	"fmt"
	"time"
)

// Даны 2 канала - в первый пишутся рандомные числа после чего они проверяются на четность и
// отправляются во второй канал. Результаты работы из второго канала пишутся в stdout.

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

	// even num filter func
	go func(ctx context.Context) {
		for {
			select {
			case v := <-chIn:
				if v%2 == 0 {
					chOut <- v
				}
			case <-ctx.Done():
				fmt.Println("filter stopped")
				return
			}
		}
	}(c)

	for i := 0; i < 100; i++ {
		chIn <- i
	}
	cancel()
	time.Sleep(200 * time.Millisecond)
}
