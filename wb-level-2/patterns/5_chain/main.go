package main

import "fmt"

func main() {
	mh := NewMainHandler()

	am := NewAuthMiddleware()
	am.SetNext(mh)

	lm := NewLogsMiddleware()
	lm.SetNext(am)

	fmt.Println("--------- ok")
	lm.Handle(&Request{false, false})
	fmt.Println("--------- logs error")
	lm.Handle(&Request{true, false})
	fmt.Println("--------- auth error")
	lm.Handle(&Request{false, true})
}

type Request struct {
	loggerError bool
	authError   bool
}

type MainHandler struct{}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

func (h *MainHandler) Handle(r *Request) {
	fmt.Println("main handler done")
}

type Handler interface {
	Handle(r *Request)
}

type LogsMiddleware struct {
	next Handler
	// ...
}

func NewLogsMiddleware() *LogsMiddleware {
	return &LogsMiddleware{}
}

func (m *LogsMiddleware) SetNext(next Handler) {
	m.next = next
}

func (m *LogsMiddleware) Handle(r *Request) {
	if r.loggerError {
		fmt.Println("logger error")
		return
	}
	fmt.Println("logs middleware done")

	if m.next != nil {
		m.next.Handle(r)
	}
}

type AuthMiddleware struct {
	next Handler
	// ...
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) SetNext(next Handler) {
	m.next = next
}

func (m *AuthMiddleware) Handle(r *Request) {
	if r.authError {
		fmt.Println("auth error")
		return
	}
	fmt.Println("auth middleware done")

	if m.next != nil {
		m.next.Handle(r)
	}
}
