package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-OrChan(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		//sig(20*time.Millisecond),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}

//OrChan ...
func OrChan(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	isCanc := false
	for _, v := range channels {
		go func(ch <-chan interface{}) {
			select {
			case <-ch:
				if !isCanc {
					isCanc = true
					close(res)
				}
				return
			case <-res:
				return
			}
		}(v)
	}
	return res
}
