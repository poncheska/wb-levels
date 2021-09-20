package util

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// https://linux.die.net/man/1/sort

type handlerFunc func([]string, *Flags) error

//Settings ...
type Settings struct {
	F *Flags
	I *IncompatibleFlags
}

//NewSettings ...
func NewSettings(f *Flags, i *IncompatibleFlags) *Settings {
	return &Settings{F: f, I: i}
}

//Flags ...
type Flags struct {
	KeyFlag int  // -k
	RevFlag bool // -r
	UnqFlag bool // -u
	BckFlag bool // -b
	ChkFlag bool // -c
}

//NewFlags ...
func NewFlags(keyFlag int, revFlag bool, unqFlag bool, bckFlag bool, chkFlag bool) *Flags {
	return &Flags{KeyFlag: keyFlag, RevFlag: revFlag, UnqFlag: unqFlag, BckFlag: bckFlag, ChkFlag: chkFlag}
}

//IncompatibleFlags ...
type IncompatibleFlags struct {
	Num bool // -n
	Mon bool // -M
	Hum bool // -h
}

//NewIncompatibleFlags ...
func NewIncompatibleFlags(num bool, mon bool, hum bool) *IncompatibleFlags {
	return &IncompatibleFlags{Num: num, Mon: mon, Hum: hum}
}

//GetHandler ...
func (s *Settings) GetHandler(ss []string) (sort.Interface, error) {
	errFlags := make([]byte, 0, 3)
	var h sort.Interface

	if s.I.Hum {
		errFlags = append(errFlags, 'h')
		h = newHumHandler(ss, s.F)
	}
	if s.I.Mon {
		errFlags = append(errFlags, 'M')
		h = newMonHandler(ss, s.F)
	}
	if s.I.Num {
		errFlags = append(errFlags, 'n')
		h = newNumHandler(ss, s.F)
	}

	switch len(errFlags) {
	case 0:
		return newDefaultHandler(ss, s.F), nil
	case 1:
		return h, nil
	default:
		return nil, fmt.Errorf("параметров «-%v» несовместимо", string(errFlags))
	}
}

//Sort ...
func Sort(input io.Reader, ss *Settings) (string, error) {
	var s []string
	unq := make(map[string]bool)

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		if ss.F.UnqFlag && !ss.F.ChkFlag {
			text := sc.Text()
			if !unq[text] {
				unq[text] = true
				s = append(s, text)
			}
		} else {
			s = append(s, sc.Text())
		}
	}
	h, err := ss.GetHandler(s)
	if err != nil {
		return "", err
	}

	if ss.F.RevFlag {
		h = sort.Reverse(h)
	}

	if ss.F.ChkFlag {
		if sort.IsSorted(h) {
			return "строки отсортированы", nil
		}
		return "строки не отсортированы", nil
	}
	sort.Sort(h)
	return strings.Join(s, "\n"), nil
}
