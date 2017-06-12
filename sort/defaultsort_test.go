package sort_test

import (
	"sort"
	"testing"
)

func BenchmarkDefaultSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := randomArray()
		sort.Ints(arr)
	}
}
