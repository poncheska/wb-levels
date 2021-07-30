package main

import (
	"fmt"
	"sync"
)

// Чем завершится данная программа?

func main() {
	// С указателем все будет работать корректно:
	// wg := &sync.WaitGroup{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// В функцию передается значение sync.WaitGroup{} и каждая горутина будет изменять только локальную копию,
		// то есть в main() счетчик wg будет только увеличиваться, а wg.Wait() никогда не разблокируется.
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
