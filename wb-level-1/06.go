package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Какие существуют способы остановить выполнения горутины? Написать примеры использования.

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch1 := make(chan int)
	ch2 := make(chan int)

	chM := wgMerge(ch1, ch2)

	go ctxWriter(ctx, ch1)
	go ctxWriter(ctx, ch2)
	go rangePrinter(chM)
	go timeoutRoutine(chM, time.Second)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	cancel()
	time.Sleep(2 * time.Second)
}

func ctxWriter(ctx context.Context, ch chan int) {
	defer fmt.Println("ctx routine stopped")
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- rand.Int()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func rangePrinter(ch chan int) {
	defer fmt.Println("range routine stopped")
	for v := range ch {
		fmt.Println("range routine:", v)
	}
}

func timeoutRoutine(ch chan int, timeout time.Duration) {
	defer fmt.Println("timeout routine stopped")
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
				continue
			}
			fmt.Println("timeout routine:", v)
		case <-time.After(timeout):
			return
		}
	}
}

func wgMerge(chs ...chan int) chan int {
	wg := &sync.WaitGroup{}
	wg.Add(len(chs))
	resCh := make(chan int)
	for _, ch := range chs {
		go func(c chan int) {
			for v := range c {
				resCh <- v
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resCh)
		fmt.Println("merge stopped")
	}()

	return resCh
}
