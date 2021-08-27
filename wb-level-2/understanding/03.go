package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	//fmt.Printf("%#v\t%#v\n", nil, err)   // <nil>   (*main.customError)(nil)
	// (https://medium.com/golangspec/equality-in-golang-ff44da79b7f1)
	// Неравенство будет выполнено, так как для выполнения равенства err == nil необходимо, чтобы сравниваемые
	// величины имели одинаковый тип и значение, в данном случае совпадает только значение.
	fmt.Println(err)
	fmt.Println(err == nil)
}
