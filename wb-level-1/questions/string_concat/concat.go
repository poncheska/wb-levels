package main

import (
	"bytes"
)

func StringSum(s ...string) string {
	res := ""
	for _, v := range s{
		res += v
	}
	return res
}

func StringBufferConcat(s ...string) string {
	var buffer bytes.Buffer

	for _, v := range s{
		buffer.WriteString(v)
	}

	return buffer.String()
}

func StringCopyConcat(resLen int,s ...string) string {
	res := make([]byte, resLen)
	cur := 0

	for _, v := range s{
		cur += copy(res[cur:],v)
	}

	return string(res)
}
