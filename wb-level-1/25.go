package main

import (
	"fmt"
	"sync"
)

// Написать свою структуру счетчик, которая будет инкрементировать и выводить значения в конкурентной среде.

func main() {
	c := NewCounter()
	k := 100
	wg := sync.WaitGroup{}
	wg.Add(k)
	for i := 0; i < k; i++ {
		go func(n int) {
			c.Add()
			fmt.Printf("Goroutine %v counted: %v\n", n, c.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type Counter struct {
	val int
	mx  sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		val: 0,
		mx:  sync.RWMutex{},
	}
}

func (c *Counter) Add() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.val++
}

func (c *Counter) Get() int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.val
}
