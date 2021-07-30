package main

import "fmt"

//Создать слайс с предварительно выделенными 100 элементами.

func main() {
	s100 := make([]int, 0, 100) // len = 0, cap = 100
	fmt.Println(len(s100), cap(s100))
}
