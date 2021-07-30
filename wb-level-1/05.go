package main

import (
	"fmt"
	"time"
)

//Написать программу, которая будет последовательно писать значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершиться.

func main() {
	ch := make(chan int)
	t := 3 * time.Second

	// пишет
	go func() {
		timer := time.After(t)
		for{
			select {
			case <-timer:
				close(ch)
				return
			default:
				ch <- 1
			}
		}
	}()

	// читает
	for v := range ch{
		fmt.Println(v)
	}
}
