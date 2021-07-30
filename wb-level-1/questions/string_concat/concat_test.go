package main

import "testing"

// go test -bench=. -benchmem

var benchData = []string{"one", "two", "three", "four", "five", "one", "two", "three", "four", "five",
	"one", "two", "three", "four", "five","one", "two", "three", "four", "five",
	"one", "two", "three", "four", "five","one", "two", "three", "four", "five",
	"one", "two", "three", "four", "five","one", "two", "three", "four", "five",
	"one", "two", "three", "four", "five","one", "two", "three", "four", "five"}

func benchmarkConcat(f func(s ...string) string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f(benchData...)
	}
}

func BenchmarkStringSum(b *testing.B) {
	benchmarkConcat(StringSum, b)
}

func BenchmarkStringBufferConcat(b *testing.B) {
	benchmarkConcat(StringBufferConcat, b)
}

func BenchmarkStringCopyConcat(b *testing.B) {
	l := 0
	for _,v := range benchData{
		l += len(v)
	}

	for n := 0; n < b.N; n++ {
		StringCopyConcat(l, benchData...)
	}
}
