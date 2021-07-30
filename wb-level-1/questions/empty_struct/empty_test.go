package empty_struct

import "testing"

func benchmarkEmptyStructSlice(k int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var a []struct{}
		for i := 0; i < k; i++ {
			a = append(a, struct{}{})
		}
	}
}

func BenchmarkEmptyStructSlice100(b *testing.B) {
	benchmarkEmptyStructSlice(100, b)
}

func BenchmarkEmptyStructSlice1000(b *testing.B) {
	benchmarkEmptyStructSlice(1000, b)
}

func BenchmarkEmptyStructSlice10000(b *testing.B) {
	benchmarkEmptyStructSlice(10000, b)
}
