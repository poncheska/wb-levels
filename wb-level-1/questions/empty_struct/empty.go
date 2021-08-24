package main

import (
	"fmt"
	"unsafe"
)

type S1 struct {
	i int64
	s struct{}
}

type S2 struct {
	s struct{}
	i int64
}


func main() {
	fmt.Println(unsafe.Sizeof(S1{}))
	fmt.Println(unsafe.Sizeof(S2{}))
}
