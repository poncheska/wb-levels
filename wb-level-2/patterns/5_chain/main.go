package main

import "fmt"

func main() {
	mh := NewLastHandler()

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

//Request ...
type Request struct {
	loggerError bool
	authError   bool
}

//LastHandler ...
type LastHandler struct{}

//NewLastHandler ...
func NewLastHandler() *LastHandler {
	return &LastHandler{}
}

//Handle ...
func (h *LastHandler) Handle(r *Request) {
	fmt.Println("main handler done")
}

//Handler ...
type Handler interface {
	Handle(r *Request)
}

//LogsMiddleware ...
type LogsMiddleware struct {
	next Handler
	// ...
}

//NewLogsMiddleware ...
func NewLogsMiddleware() *LogsMiddleware {
	return &LogsMiddleware{}
}

//SetNext ...
func (m *LogsMiddleware) SetNext(next Handler) {
	m.next = next
}

//Handle ...
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

//AuthMiddleware ...
type AuthMiddleware struct {
	next Handler
	// ...
}

//NewAuthMiddleware ...
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

//SetNext ...
func (m *AuthMiddleware) SetNext(next Handler) {
	m.next = next
}

//Handle ...
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
