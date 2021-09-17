package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(Unpack("saf3saa0\\\\\\3aaa"))
}

func Unpack(s string) (string, error) {
	var res bytes.Buffer
	var prevRune rune
	var escaped bool

	size := len([]rune(s)) + 1
	runes := make([]rune, size, size)
	copy(runes, []rune(s))

	for _, v := range runes {
		switch {
		case unicode.IsDigit(v):
			switch {
			case prevRune == 0:
				return "", errors.New("invalid input")
			case prevRune == '\\' && !escaped:
				prevRune = v
				escaped = true
			default:
				k, err := strconv.Atoi(string(v))
				if err != nil {
					return "", err
				}
				for j := 0; j < k; j++ {
					res.WriteRune(prevRune)
				}
				prevRune = 0
				escaped = false
			}
		case v == '\\':
			if prevRune == '\\' {
				if escaped {
					res.WriteRune(prevRune)
					escaped = false
				} else {
					escaped = true
				}
			} else {
				if prevRune != 0 {
					res.WriteRune(prevRune)
				}
				prevRune = v
				escaped = false
			}
		default:
			if prevRune == '\\' && !escaped {
				return "", errors.New("invalid input")
			}
			if prevRune != 0 {
				res.WriteRune(prevRune)
			}
			prevRune = v
			escaped = false
		}
	}
	return res.String(), nil
}

//func Unpack(s string) (string, error) {
//	res := bytes.NewBuffer([]byte{})
//	numBuf := bytes.NewBuffer([]byte{})
//	var wasEscaped bool
//	var lastNotDigit rune
//
//	size := len([]rune(s)) + 1
//	runes := make([]rune, size, size)
//	copy(runes, []rune(s))
//
//	for _, v := range runes {
//		if unicode.IsDigit(v) {
//			if lastNotDigit == 0 {
//				return "", fmt.Errorf("invalid input")
//			}
//			numBuf.WriteRune(v)
//		} else {
//			if lastNotDigit == '\\' && !wasEscaped {
//				if v == '\\' {
//					wasEscaped = true
//					continue
//				}
//
//				numStr := numBuf.String()
//				if numStr == "" {
//					return "", fmt.Errorf("invalid input")
//				} else {
//					res.WriteString(numStr)
//					numBuf.Reset()
//					lastNotDigit = v
//					continue
//				}
//			}
//
//			num, err := readNum(numBuf)
//			if err != nil {
//				return "", err
//			}
//
//			for j := 0; j < num; j++ {
//				res.WriteRune(lastNotDigit)
//			}
//
//			lastNotDigit = v
//			wasEscaped = false
//		}
//	}
//	return res.String(), nil
//}

//func readNum(buf *bytes.Buffer) (int,error){
//	defer buf.Reset()
//	var num int
//	numStr := buf.String()
//	if len(numStr) == 0 {
//		return 1, nil
//	} else {
//		var err error
//		num, err = strconv.Atoi(numStr)
//		if err != nil {
//			return 0, err
//		}
//		return num, nil
//	}
//}
