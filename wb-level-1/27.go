package main

import (
	"fmt"
	"strings"
)

// Написать программу, которая переворачивает слова в строке (snow dog sun - sun dog snow).

func main() {
	str := "bibi bobo baba bubu bebe"
	fmt.Println(reverseWordsOrder(str))
}

func reverseWordsOrder(str string) string {
	s := strings.Split(str, " ")
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return strings.Join(s," ")
}
