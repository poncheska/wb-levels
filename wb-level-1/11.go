package main

import "fmt"

// Написать пересечение двух неупорядоченных массивов.

func main() {
	s1 := []int{5, 6, 8, 3, 2, 7, 5, 6}
	s2 := []int{7, 4, 7, 3, 5, 3, 6, 6, 5, 3}
	fmt.Println(intersect(s2, s1))
}

func intersect(s1, s2 []int) []int {
	var res []int

	counter := make(map[int]int)
	for _, v := range s1 {
		counter[v]++
	}

	for _, v := range s2 {
		if counter[v] > 0 {
			counter[v]--
			res = append(res, v)
		}
	}

	return res
}
