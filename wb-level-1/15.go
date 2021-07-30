package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 0
	b := 1
	a, b = b, a
	fmt.Printf("a = %v ; b = %v\n", a, b)
}
