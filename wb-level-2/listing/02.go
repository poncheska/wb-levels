package main

import (
	"fmt"
)

// (https://blog.golang.org/defer-panic-and-recover)
// defer может считывать и изменять именованные возвращаемые значения
func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

// defer вызывается после return, в данном случае значение уже возвращено до срабатывания defer
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())		   // 2
	fmt.Println(anotherTest()) // 1
}
