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

	//HandlerV1 ...
	HandlerV1 = func(r *Request) {
		fmt.Printf("handler v1 got %v\n", r.Text)
	}

	//HandlerV2 ...
	HandlerV2 = func(r *Request) {
		fmt.Printf("handler v2 got %v\n", r.Text)
	}

	//HandlerV3 ...
	HandlerV3 = func(r *Request) {
		fmt.Printf("handler v3 got %v\n", r.Text)
	}
)

//Request ...
type Request struct {
	Text string
}

//NewRequest ...
func NewRequest(text string) *Request {
	return &Request{Text: text}
}

//Server ...
type Server struct {
	HandlerFunc func(r *Request)
}

//NewServer ...
func NewServer(handlerFunc func(r *Request)) *Server {
	return &Server{HandlerFunc: handlerFunc}
}

//HandlerStruct ...
type HandlerStruct struct {
	info string
}

//NewHandlerStruct ...
func NewHandlerStruct(info string) *HandlerStruct {
	return &HandlerStruct{info: info}
}

//Handle ...
func (h *HandlerStruct) Handle(r *Request) {
	fmt.Printf("handler(%v) got %v\n", h.info, r.Text)
}
