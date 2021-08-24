package main

import "fmt"

// Реализовать паттерн адаптер на любом примере.

func main() {
	p := Person{
		Name:    "John",
		Surname: "Cena",
	}
	PrintString(PersonAdapter{p})
}

func PrintString(g Getter) {
	fmt.Println(g.GetString())
}

type Getter interface {
	GetString() string
}

type PersonAdapter struct {
	p Person
}

func (pa PersonAdapter) GetString() string {
	return pa.p.GetInfo()
}

type Person struct {
	Name    string
	Surname string
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("%v %v", p.Name, p.Surname)
}
