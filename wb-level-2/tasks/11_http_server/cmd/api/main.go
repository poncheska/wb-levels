package main

import (
	"git.wildberries.ru/Kovgar.Aleksey/wb-levels/wb-level-2/tasks/11_http_server/internal/handler"
	"net/http"
)

func main() {
	s := handler.NewServer()

	http.Serve()
}
