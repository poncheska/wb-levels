package main

import (
	"bytes"
	"fmt"
)

// Написать программу, которая переворачивает строку. Символы могут быть unicode.

func main() {
	s := "йцукенгшщзqwertyuiop12345!№;%:?"
	fmt.Println(reverseString(s))
}

func reverseString(s string) string{
	var buffer bytes.Buffer
	sr := []rune(s)

	for i := 0; i < len(sr); i++ {
		buffer.WriteRune(sr[len(sr)-1-i])
	}

	return buffer.String()
}