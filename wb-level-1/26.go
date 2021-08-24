package main

import (
	"fmt"
)

// Написать программу, которая переворачивает строку. Символы могут быть unicode.

func main() {
	s := "йцукенгшщзqwertyuiop12345!№;%:?"
	fmt.Println(reverseString(s))
	fmt.Println(s)
}

func reverseString(s string) string {
	res := []rune(s)

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)
}
