package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewStringSet()
	set.AddSlice(s)
	fmt.Println(set.GetSlice(), len(set.GetSlice()))
}

type stringSet map[string]struct{}

func NewStringSet() *stringSet {
	m := stringSet(make(map[string]struct{}))
	return &m
}

func (s *stringSet) Add(str string) {
	(*s)[str] = struct{}{}
}

func (s *stringSet) AddSlice(str []string) {
	for _, v := range str {
		s.Add(v)
	}
}

func (s *stringSet) Delete(str string) {
	delete(*s, str)
}

func (s *stringSet) GetSlice() []string {
	res := make([]string, 0, len(*s))
	for k := range *s {
		res = append(res, k)
	}
	return res
}
