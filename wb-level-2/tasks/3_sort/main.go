package main

import (
	"flag"
	"fmt"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/tasks/3_sort/util"
	"os"
)

func main() {
	k := flag.Int("k", -1, "start a key at POS (default whole line / first key)")
	n := flag.Bool("n", false, "compare according to string numerical value")
	r := flag.Bool("r", false, "reverse the result of comparisons")
	u := flag.Bool("u", false, "only unique strings")
	m := flag.Bool("M", false, "compare (unknown) < 'JAN' < ... < 'DEC'")
	b := flag.Bool("b", false, "ignore leading blanks")
	c := flag.Bool("c", false, "check for sorted input; do not sort")
	h := flag.Bool("h", false, "compare human readable numbers (e.g., 2K 1G)")
	flag.Parse()
	ifs := util.NewIncompatibleFlags(*n, *m, *h)
	fs := util.NewFlags(*k-1, *r, *u, *b, *c)
	ss := util.NewSettings(fs, ifs)
	reader, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("sort: ", err.Error())
		return
	}
	res, err := util.Sort(reader, ss)
	if err != nil {
		fmt.Println("sort: ", err.Error())
	} else {
		fmt.Println(res)
	}
}
