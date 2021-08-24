package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать набор из N воркеров, которые читают из канала произвольные
//данные и выводят в stdout. Данные в канал пишутся из главного потока.
//Необходима возможность выбора кол-во воркеров при старте, а также способ
//завершения работы всех воркеров.


func main() {
	nFlag := flag.Int("n", 10, "help message for flag n")
	flag.Parse()

	n := *nFlag

	ctx, cancel := context.WithCancel(context.Background())

	c := startPrinters(ctx, n)

	go spammer(ctx, c)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	cancel()
	time.Sleep(time.Second)
}

func startPrinters(ctx context.Context, n int) chan string {
	ch := make(chan string)
	for i:=0;i<n;i++{
		go printer(ctx, ch)
	}
	return ch
}

func printer(ctx context.Context, ch chan string){
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-ctx.Done():
			fmt.Println("printer stopped")
			return
		}
	}
}

func spammer(ctx context.Context, ch chan string){
	for {
		select {
		case <-ctx.Done():
			fmt.Println("spammer stopped")
			return
		default:
			ch <- "spam"
		}
	}
}
