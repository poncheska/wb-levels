package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов значений взятых
//из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	s := []int{2,4,6,8,10}
	printQuads(s)
}

func printQuads(s []int){
	wg := &sync.WaitGroup{}
	wg.Add(len(s))
	for _, v := range s{
		go func(n int) {
			fmt.Println(n*n)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
