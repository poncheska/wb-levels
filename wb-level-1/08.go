package main

import "fmt"

// Дана переменная int64. Написать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	fmt.Println(setBit(0, 1, true))  // 0 = 0 -> 10 = 2
	fmt.Println(setBit(4, 1, true))  // 4 = 100 -> 110 = 6
	fmt.Println(setBit(6, 1, true))  // 6 = 110 -> 110 = 6
	fmt.Println(setBit(6, 1, false)) // 6 = 110 -> 100 = 4
	fmt.Println(setBit(6, 2, false)) // 6 = 110 -> 010 = 2
}

func setBit(v int64, pos int, bit bool) int64 {
	if bit {
		return v | (1 << pos)
	} else {
		return v &^ (1 << pos)
	}
}
