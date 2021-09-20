package util

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

//Settings ...
type Settings struct {
	FlagF string
	FlagD string
	FlagS bool
}

//NewSettings ...
func NewSettings(flagF string, flagD string, flagS bool) *Settings {
	return &Settings{FlagF: flagF, FlagD: flagD, FlagS: flagS}
}

//ParseF ...
func ParseF(s string) ([]int, error) {
	var res []int
	ss := strings.Split(s, ",")
	for _, v := range ss {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res = append(res, n)
	}
	return res, nil
}

//Cut ...
func Cut(r io.Reader, ss *Settings) (string, error) {
	f, err := ParseF(ss.FlagF)
	if err != nil {
		return "", err
	}

	var res []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		var lRes []string
		fields := strings.Split(sc.Text(), ss.FlagD)
		if !ss.FlagS || len(fields) > 1 {
			for _, v := range f {
				if v > -1 && v < len(fields) {
					lRes = append(lRes, fields[v])
				} else {
					lRes = append(lRes, "")
				}
			}
			res = append(res, strings.Join(lRes, ss.FlagD))
		}
	}

	return strings.Join(res, "\n"), nil
}
