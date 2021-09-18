package main

import (
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/handler"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errChan := make(chan error)

	go func() {
		svc := service.NewEventService()
		h := handler.LoggerMiddleware(handler.New(svc))
		errChan <- http.ListenAndServe(":8080", h)
	}()
	log.Println("server started")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("error: %v", err)
	case <-sigChan:
		log.Println("terminated")
	}
}
