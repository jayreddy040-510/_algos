package search

import (
	"math/rand"
	"testing"
)

const size = 100000

var (
	testSlice []int
	target    int
)

func init() {
	for i := 0; i < size; i++ {
		testSlice = append(testSlice, i)
	}
	target = rand.Intn(size)
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linearSearch(testSlice, target)
	}
}

func BenchmarkBSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BSearch(testSlice, target)
	}
}
