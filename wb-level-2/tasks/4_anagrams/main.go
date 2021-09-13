package main

import (
	"sort"
	"strings"
)

func main() {

}

type anagram struct {
	first string
	core  map[rune]int
	elem  map[string]struct{}
}

func newAnagram(s string) *anagram {
	s = strings.ToLower(s)

	m := make(map[rune]int)
	e := make(map[string]struct{})
	rs := []rune(s)

	for _, v := range rs {
		m[v] += 1
	}

	e[s] = struct{}{}

	return &anagram{
		s,
		m,
		e,
	}
}

func (a *anagram) joinIfOk(s string) bool {
	s = strings.ToLower(s)

	m := copyMap(a.core)
	rs := []rune(s)

	for _, v := range rs {
		if m[v] == 0 {
			return false
		} else {
			m[v] -= 1
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	a.elem[s] = struct{}{}

	return true
}

func (a *anagram) GetElem() (string, []string) {
	res := make([]string, 0, len(a.elem))
	for k, _ := range a.elem {
		res = append(res, k)
	}
	sort.Strings(res)
	return a.first, res
}

func copyMap(m map[rune]int) map[rune]int {
	cp := make(map[rune]int)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func Anagrams(s []string) map[string][]string {
	var as []*anagram
	for _, v := range s {
		joined := false
		for _, w := range as {
			if w.joinIfOk(v) {
				joined = true
				break
			}
		}
		if !joined {
			as = append(as, newAnagram(v))
		}
	}

	res := make(map[string][]string)
	for _, v := range as {
		f, e := v.GetElem()
		if len(e) != 1 {
			res[f] = e
		}
	}
	return res
}
