package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"} // cap=3
	modifySlice(s)
	fmt.Println(s) // [3 2 3]
}

// Слайс хранит в себе ссылку на массив, поэтому при передаче его в качестве аргумента функции
// все изменения происходящие со слайсом внутри функции до расширения его массива, будут
// отражены на оригинальном слайсе.
func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4") // Слайс удваивает cap, создавая новый слайс
	i[1] = "5"
	i = append(i, "6")
}
