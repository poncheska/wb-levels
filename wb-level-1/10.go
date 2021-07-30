package main

import "fmt"

// Дана последовательность температурных колебаний (-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5).
// Объединить данный значения в группы с шагов в 10 градусов. Последовательность в подмножностве не важна.

func main() {
	s := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(groupDec(s))
}

func groupDec(s []float64) map[int][]float64 {
	res := make(map[int][]float64)
	for _, v := range s {
		res[(int(v)/10)*10] = append(res[(int(v)/10)*10], v)
	}
	return res
}
