package util

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//Settings ...
type Settings struct {
	FlagAft int  // -A
	FlagBef int  // -B
	FlagCtx int  // -C
	FlagCnt bool // -c
	FlagIgn bool // -i
	FlagInv bool // -v
	FlagFix bool // -F
	FlagNum bool // -n
}

//NewSettings ...
func NewSettings(flagAft int, flagBef int, flagCtx int, flagCnt bool, flagIgn bool, flagInv bool, flagFix bool, flagNum bool) *Settings {
	return &Settings{
		FlagAft: flagAft,
		FlagBef: flagBef,
		FlagCtx: flagCtx,
		FlagCnt: flagCnt,
		FlagIgn: flagIgn,
		FlagInv: flagInv,
		FlagFix: flagFix,
		FlagNum: flagNum,
	}
}

//Checker ...
type Checker interface {
	check(s string) bool
}

//RegexChecker ...
type RegexChecker struct {
	R *regexp.Regexp
}

//NewRegexChecker ...
func NewRegexChecker(r *regexp.Regexp) *RegexChecker {
	return &RegexChecker{R: r}
}

func (c *RegexChecker) check(s string) bool {
	return c.R.MatchString(s)
}

//EqualChecker ...
type EqualChecker struct {
	pat string
}

//NewEqualChecker ...
func NewEqualChecker(pat string) *EqualChecker {
	return &EqualChecker{pat: pat}
}

func (c *EqualChecker) check(s string) bool {
	return strings.Contains(s, c.pat)
}

//Grep ...
func Grep(r io.Reader, pattern string, ss *Settings) (string, error) {
	var chr Checker

	if ss.FlagIgn {
		pattern = strings.ToLower(pattern)
	}

	if ss.FlagFix {
		chr = NewEqualChecker(pattern)
	} else {
		rx, err := regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
		chr = NewRegexChecker(rx)
	}

	lt := ss.FlagCtx
	rt := ss.FlagCtx
	if ss.FlagBef > 0 {
		lt = ss.FlagBef
	}
	if ss.FlagAft > 0 {
		rt = ss.FlagAft
	}

	mm := make(map[int]bool)
	var resIds []int
	var allStr []string
	counter := 0
	i := 0
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		t := sc.Text()
		allStr = append(allStr, t)
		if ss.FlagIgn {
			t = strings.ToLower(t)
		}

		if chr.check(t) {
			counter++
			if lt > 0 {
				for j := i - lt; j < i; j++ {
					if _, ok := mm[j]; !ok {
						resIds = append(resIds, j)
						mm[j] = false
					}
				}
			}
			if rt > 0 {
				for j := i + rt; j > i; j-- {
					if _, ok := mm[j]; !ok {
						resIds = append(resIds, j)
						mm[j] = false
					}
				}
			}
			if _, ok := mm[i]; !ok {
				resIds = append(resIds, i)
			}
			mm[i] = true
		}
		i++
	}

	if ss.FlagCnt {
		return fmt.Sprintf("%v matches", counter), nil
	}

	sort.Ints(resIds)

	var res []string
	lIndex := ""
	if !ss.FlagInv {
		match := "  "
		for _, k := range resIds {
			if k > -1 && k < len(allStr) {
				if mm[k] {
					match = "* "
				} else {
					match = "  "
				}
				if ss.FlagNum {
					lIndex = strconv.Itoa(k) + " "
				}
				res = append(res, fmt.Sprintf("%v%v%v", lIndex, match, allStr[k]))
			}
		}
	} else {
		for i, v := range allStr {
			if _, ok := mm[i]; !ok {
				if ss.FlagNum {
					lIndex = strconv.Itoa(i) + " "
				}
				res = append(res, fmt.Sprintf("%v%v", lIndex, v))
			}
		}
	}

	return strings.Join(res, "\n"), nil
}
