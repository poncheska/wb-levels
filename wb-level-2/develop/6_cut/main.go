package main

import (
	"flag"
	"fmt"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/6_cut/util"
	"os"
)

func main() {
	f := flag.String("f", "", "fields")
	d := flag.String("d", "   ", "delimiter")
	s := flag.Bool("s", false, "separated")
	flag.Parse()

	reader, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("grep: ", err.Error())
		return
	}
	sets := util.NewSettings(*f, *d, *s)

	res, err := util.Cut(reader, sets)

	if err != nil {
		fmt.Println("grep: ", err.Error())
	} else {
		fmt.Println(res)
	}

}
