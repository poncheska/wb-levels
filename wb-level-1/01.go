package main

import "fmt"

// Реализовать композицию структуры Action от родительской структуры Human.

func main() {
	a1 := ActionEmbedded{Human{}}
	a2 := ActionComposed{h: Human{}}
	a1.jump()
	a2.jump()
	fmt.Println(a1.log)
	fmt.Println(a2.h.log)
}

type Human struct {
	log string
}

type ActionEmbedded struct {
	Human
}

func (ae *ActionEmbedded) jump(){
	ae.log += "jumped\n"
}

type ActionComposed struct {
	h Human
}

func (ac *ActionComposed) jump(){
	ac.h.log += "jumped\n"
}