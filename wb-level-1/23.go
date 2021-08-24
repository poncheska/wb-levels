package main

import (
	"fmt"
	"sort"
)

// Написать бинарный поиск встроенными методами языка.

func main() {
	s := []int{1, 5, 8, 12, 23, 34, 64, 85, 111, 126, 158, 161, 266, 355, 677, 888, 964}
	fmt.Println(binarySearch(s, 314))    // -1
	fmt.Println(binarySearch(s, 23))     // 4
	fmt.Println(sort.SearchInts(s, 314)) // -1
	fmt.Println(sort.SearchInts(s, 23))  // 4
}

func binarySearch(s []int, v int) int {
	l := 0
	r := len(s) - 1

	for l <= r {
		c := (l + r) / 2
		switch {
		case s[c] == v:
			return c
		case s[c] < v:
			l = c + 1
		case s[c] > v:
			r = c - 1
		}
	}

	return -1
}
