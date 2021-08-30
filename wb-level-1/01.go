package main

import (
	"fmt"
)

// Реализовать композицию структуры Action от родительской структуры Human.

func main() {
	a1 := ActionEmbedded{NewHuman()}
	a2 := ActionComposed{h: NewHuman()}
	a1.Jump()
	a2.Jump()
	fmt.Println(a1.log)
	fmt.Println(a2.h.log)
}

type Human struct {
	log string
}

func NewHuman() *Human{
	return &Human{}
}

type ActionEmbedded struct {
	*Human
}

func (ae *ActionEmbedded) Jump(){
	ae.log += "jumped\n"
}

type ActionComposed struct {
	h *Human
}

func (ac *ActionComposed) Jump(){
	ac.h.log += "jumped\n"
}