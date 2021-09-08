package main

import "fmt"

func main() {
	s1 := NewServer(HandlerV1)
	s1.HandlerFunc(NewRequest("aaaaa"))

	s2 := NewServer(HandlerV2)
	s2.HandlerFunc(NewRequest("bbbbb"))

	s3 := NewServer(HandlerV3)
	s3.HandlerFunc(NewRequest("ccccc"))

	s := NewServer(NewHandlerStruct("my handler").Handle)
	s.HandlerFunc(NewRequest("ddddd"))
}

var (
	HandlerV1 = func(r *Request) {
		fmt.Printf("handler v1 got %v\n", r.Text)
	}

	HandlerV2 = func(r *Request) {
		fmt.Printf("handler v2 got %v\n", r.Text)
	}

	HandlerV3 = func(r *Request) {
		fmt.Printf("handler v3 got %v\n", r.Text)
	}
)

type Request struct {
	Text string
}

func NewRequest(text string) *Request {
	return &Request{Text: text}
}

type Server struct {
	HandlerFunc func(r *Request)
}

func NewServer(handlerFunc func(r *Request)) *Server {
	return &Server{HandlerFunc: handlerFunc}
}

type HandlerStruct struct {
	info string
}

func NewHandlerStruct(info string) *HandlerStruct {
	return &HandlerStruct{info: info}
}

func (h *HandlerStruct) Handle(r *Request) {
	fmt.Printf("handler(%v) got %v\n", h.info, r.Text)
}
