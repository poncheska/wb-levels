package main

import "fmt"

// Написать программу, которая проверяет, что все символы в строке уникальные.

func main() {
	s1 := "fdввпвпеик43"
	fmt.Println(IsSymbolsUnique(s1))
	s2 := "qwreцваи123-=*(:"
	fmt.Println(IsSymbolsUnique(s2))
}

func IsSymbolsUnique(str string) bool {
	s := []rune(str)
	m := make(map[rune]struct{})
	for _,v := range s{
		m[v] = struct{}{}
	}
	return len(m)==len(s)
}
