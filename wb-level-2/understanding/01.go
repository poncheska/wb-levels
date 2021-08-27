package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	// Возвращается слайс с len=3 cap=4, до расширения слайса его массивом будет a[1:]
	//b[0] = 0
	//b = append(b,0)
	//fmt.Println(a)
	fmt.Println(b)
}
