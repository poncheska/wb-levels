package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	i := 3
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s)
}
