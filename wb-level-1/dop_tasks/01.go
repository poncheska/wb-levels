package main

import "fmt"

func main() {
	s1 := []int{5, 6, 8, 3, 2, 7, 5, 6}
	s2 := []int{7, 4, 7, 3, 5, 3, 6, 6, 5, 3}
	fmt.Println(union(s2, s1))
}

func union(s1, s2 []int) []int {
	var res []int

	counter := make(map[int]struct{})
	for _, v := range s1 {
		counter[v] = struct{}{}
	}

	for _, v := range s2 {
		counter[v] = struct{}{}
	}

	for k,_ := range counter{
		res = append(res, k)
	}

	return res
}
