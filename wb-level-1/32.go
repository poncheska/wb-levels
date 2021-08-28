package main

import (
	"fmt"
	"sync"
	"time"
)

// Написать собственную функцию Sleep.

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		MySleep(1 * time.Second)
		fmt.Println("MySleep")
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Sleep")
		wg.Done()
	}()
	wg.Wait()
}

func MySleep(t time.Duration) {
	<-time.After(t)
}
