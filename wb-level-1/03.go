package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Println(quadsSum(s))
}

func quadsSum(s []int) int {
	wg := &sync.WaitGroup{}
	wg.Add(len(s))
	resChan := make(chan int, len(s))
	for _, v := range s {
		go func(n int) {
			resChan <- n * n
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	res := 0
	for v := range resChan {
		res += v
	}

	return res
}
