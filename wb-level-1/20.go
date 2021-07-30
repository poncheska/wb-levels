package main

import "fmt"

// Какой результат выполнения данного кода и почему?

func main() {
	slice := []string{"a", "a"} // len=cap=2

	func(slice []string) {
		// Происходит расширение массива на который ссылается слайс переданный из main(),
		//то есть новый элемент добавляется только в массив локального слайса, и все дальнейшие
		//изменения слайса не будут отражаться на оригинале.
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice) // [b b a]
	}(slice)
	fmt.Print(slice)     // [a a]
}
