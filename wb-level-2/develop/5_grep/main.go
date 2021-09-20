package main

import (
	"flag"
	"fmt"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/5_grep/util"
	"os"
)

func main() {
	A := flag.Int("A", 0, "")
	B := flag.Int("B", 0, "")
	C := flag.Int("C", 0, "")
	c := flag.Bool("c", false, "")
	i := flag.Bool("i", false, "")
	v := flag.Bool("v", false, "")
	F := flag.Bool("F", false, "")
	n := flag.Bool("n", false, "")
	flag.Parse()

	sets := util.NewSettings(*A, *B, *C, *c, *i, *v, *F, *n)

	if len(os.Args) < 3 {
		fmt.Println("grep: not enough arguments")
	}

	reader, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("grep: ", err.Error())
		return
	}

	pattern := os.Args[len(os.Args)-2]

	res, err := util.Grep(reader, pattern, sets)

	if err != nil {
		fmt.Println("grep: ", err.Error())
	} else {
		fmt.Println(res)
	}
}
