package main

import "fmt"

// Что выведет данная программа и почему?

func someAction(v []int8, b int8) {
	// ссылается на тот же массив, что и переданный из main() слайс
	v[0] = 100
	// создается новый массив в который будет добавлен новый элемент,
	// слайс в main() не изменится
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5} // len=cap=5
	someAction(a, 6)
	fmt.Println(a) // [100 2 3 4 5]
}
