package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// К каким негативным последствиям может привести данный кусок кода и как это исправить?
var symbols = []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюяabcdefghijklmnopqrstuvwxyz")

var justString string

func someFunc() {
	//v := createHugeString(1 << 5)
	//fmt.Println(v)          // tкnqgdciёxyёejфpтtъвsббiчgуждфyф
	// берет первые 10 байт на последнем символе может оборваться кодировка
	//justString = v[:10]
	//fmt.Println(justString) // tкnqgdci�
	v := createHugeString(1 << 5)
	fmt.Println(v)
	// берет первые 10 символов
	justString = string([]rune(v)[:10])
	fmt.Println(justString)
}

func main() {
	someFunc()
}

func createHugeString(l int) string {
	var b bytes.Buffer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		b.WriteRune(symbols[r.Intn(len(symbols)-1)])
	}
	return b.String()
}
