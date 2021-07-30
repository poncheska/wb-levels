package main

import "fmt"

// Написать быструю сортировку встроенными методами языка.

func main() {
	s := []int{75, 34, 32, 353, 3, 747, 21, 65, 9, 454, 35, 75, 235, 3, 65, 65, 65}
	QuickSort(s)
	fmt.Println(s)
}

func QuickSort(s []int) {
	qSort(s, 0, len(s)-1)
}

func qSort(s []int, l, r int) {
	if l < r {
		p := l
		ll := l + 1
		rr := r

		for ll <= rr {
			switch {
			case s[p] >= s[ll]:
				ll++
			case s[p] <= s[rr]:
				rr--
			case s[ll] > s[p] && s[rr] <= s[p]:
				s[ll], s[rr] = s[rr], s[ll]
				ll++
				rr--
			}
		}

		s[p], s[rr] = s[rr], s[p]
		qSort(s, l, rr-1)
		qSort(s, rr+1, r)
	}
}
