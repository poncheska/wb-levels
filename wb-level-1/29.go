package main

import (
	"fmt"
	"math/big"
)

// Написать программу, которая перемножает, делит, складывает,
// вычитает 2 числовых переменных a,b, значение которые > 2^20.

func main() {
	a := big.NewInt(10)
	b := big.NewInt(5)
	fmt.Printf("%v+%v=%v\n", a, b, Sum(a, b))
	fmt.Printf("%v-%v=%v\n", a, b, Sub(a, b))
	fmt.Printf("%v*%v=%v\n", a, b, Mul(a, b))
	fmt.Printf("%v/%v=%v\n", a, b, Div(a, b))
}

func Sum(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}
