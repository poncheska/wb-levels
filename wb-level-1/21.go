package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая в конкурентном виде читает элементы из массива в stdout.

func main() {
	slice := []string{"a", "c", "d", "e", "f", "g", "h"}
	PrintSlice(slice)
}

func PrintSlice(s []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(s))
	for _, v := range s {
		go func(str string) {
			fmt.Println(str)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
