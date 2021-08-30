package main

import (
	"fmt"
	"reflect"
)

// Написать программу, которая в рантайме способна определить тип
// переменной — int, string, bool, channel из переменной типа interface{}.

func main() {
	ch := make(chan int)
	printType(ch)
	printTypeReflect(ch)
	printTypeSimple(ch)
}

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("unknown")
	}
}

func printTypeReflect(v interface{}) {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("chan")
	default:
		fmt.Println("unknown")
	}
}

func printTypeSimple(v interface{}) {
	//fmt.Println(reflect.TypeOf(v).String())
	fmt.Printf("%T\n", v)
}
