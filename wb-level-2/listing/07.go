package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			// канал закрывается и перестает быть блокирующим и начинает возвращать дефолтные значения,
			// а так как селект в бесконечном цикле то в канал c будут бесконечно приходить нули.
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
		// если заменить цикл на такой, то все будет работать
		//for {
		//	if a==nil && b==nil{
		//		close(c)
		//		break
		//	}
		//	select {
		//	case v,ok := <-a:
		//		if !ok{
		//			a = nil
		//			continue
		//		}
		//		c <- v
		//	case v,ok := <-b:
		//		if !ok{
		//			b = nil
		//			continue
		//		}
		//		c <- v
		//	}
		//}
	}()
	return c
}

// лучше использовать такой мерж
func normMerge(a, b <-chan int) <-chan int {
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for v := range a{
			c <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range b{
			c <- v
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
