package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
