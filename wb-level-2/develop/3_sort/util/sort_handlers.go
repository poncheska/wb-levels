package util

import (
	"strconv"
	"strings"
	"unicode"
)

//MinInt ...
const MinInt = -int(^uint(0)>>1) - 1

type defaultHandler struct {
	s  []string
	sf []string
	f  *Flags
}

func newDefaultHandler(s []string, f *Flags) *defaultHandler {
	sf := parseDefSF(s, f)
	return &defaultHandler{s: s, sf: sf, f: f}
}

func parseDefSF(s []string, f *Flags) []string {
	res := make([]string, 0, len(s))
	for _, v := range s {
		if f.KeyFlag > -1 {
			fs := strings.Fields(v)
			if len(fs) > f.KeyFlag {
				res = append(res, fs[f.KeyFlag])
			} else {
				res = append(res, "")
			}
		} else {
			if f.BckFlag {
				res = append(res, strings.TrimSpace(v))
			} else {
				res = append(res, v)
			}
		}
	}
	return res
}

func (h *defaultHandler) Len() int {
	return len(h.s)
}

func (h *defaultHandler) Less(i, j int) bool {
	return h.sf[i] < h.sf[j]
}

func (h *defaultHandler) Swap(i, j int) {
	h.sf[i], h.sf[j] = h.sf[j], h.sf[i]
	h.s[i], h.s[j] = h.s[j], h.s[i]
}

// numHandler используется как numHandler, monHandler и humHandler
type numHandler struct {
	s  []string
	sf []int
	f  *Flags
}

func newNumHandler(s []string, f *Flags) *numHandler {
	sf := parseNumSF(s, f)
	return &numHandler{s: s, sf: sf, f: f}
}

func parseNumSF(s []string, f *Flags) []int {
	res := make([]int, 0, len(s))
	for _, v := range s {
		fs := strings.Fields(v)
		if f.KeyFlag > -1 {
			if len(fs) > f.KeyFlag {
				num, err := strconv.Atoi(fs[f.KeyFlag])
				if err != nil {
					res = append(res, MinInt)
				} else {
					res = append(res, num)
				}
			} else {
				res = append(res, MinInt)
			}
		} else {
			if len(fs) > 0 {
				num, err := strconv.Atoi(fs[0])
				if err != nil {
					res = append(res, MinInt)
				} else {
					res = append(res, num)
				}
			} else {
				res = append(res, MinInt)
			}
		}
	}
	return res
}

func (h *numHandler) Len() int {
	return len(h.s)
}

func (h *numHandler) Less(i, j int) bool {
	return h.sf[i] < h.sf[j]
}

func (h *numHandler) Swap(i, j int) {
	h.sf[i], h.sf[j] = h.sf[j], h.sf[i]
	h.s[i], h.s[j] = h.s[j], h.s[i]
}

func newHumHandler(s []string, f *Flags) *numHandler {
	sf := parseHumSF(s, f)
	return &numHandler{s: s, sf: sf, f: f}
}

func parseHumSF(s []string, f *Flags) []int {
	res := make([]int, 0, len(s))
	for _, v := range s {
		fs := strings.Fields(v)
		if f.KeyFlag > -1 {
			if len(fs) > f.KeyFlag {
				res = append(res, ParseHumNum(fs[f.KeyFlag]))
			} else {
				res = append(res, MinInt)
			}
		} else {
			if len(fs) > 0 {
				res = append(res, ParseHumNum(fs[0]))
			} else {
				res = append(res, MinInt)
			}
		}
	}
	return res
}

func newMonHandler(s []string, f *Flags) *numHandler {
	sf := parseMonSF(s, f)
	return &numHandler{s: s, sf: sf, f: f}
}

func parseMonSF(s []string, f *Flags) []int {
	res := make([]int, 0, len(s))
	for _, v := range s {
		fs := strings.Fields(v)
		if f.KeyFlag > -1 {
			if len(fs) > f.KeyFlag {
				res = append(res, MonToNum(fs[f.KeyFlag]))
			} else {
				res = append(res, 0)
			}
		} else {
			if len(fs) > 0 {
				res = append(res, MonToNum(fs[0]))
			} else {
				res = append(res, 0)
			}
		}
	}
	return res
}

//MonToNum ...
func MonToNum(s string) int {
	switch s {
	case "JAN":
		return 1
	case "FEB":
		return 2
	case "MAR":
		return 3
	case "APR":
		return 4
	case "MAY":
		return 5
	case "JUN":
		return 6
	case "JUL":
		return 7
	case "AUG":
		return 8
	case "SEP":
		return 9
	case "OCT":
		return 10
	case "NOV":
		return 11
	case "DEC":
		return 12
	default:
		return 0
	}
}

//ParseHumNum ...
func ParseHumNum(s string) int {
	num, err := strconv.Atoi(s)
	if err == nil {
		return num
	}

	if len(s) < 2 {
		return MinInt
	}
	num, err = strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return MinInt
	}

	switch byte(unicode.ToUpper(rune(s[len(s)-1]))) {
	case 'K':
		return num * 1000
	case 'M':
		return num * 1000000
	case 'G':
		return num * 1000000000
	default:
		return MinInt
	}
}
