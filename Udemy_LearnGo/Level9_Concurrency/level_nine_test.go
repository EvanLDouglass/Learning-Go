// This file is used to benchmark the two approaches to incrementing
// in a multi-threaded function (i.e. Mutex locks vs. atomic package).
// It is recommended to comment out the Printlns at the bottom
// of each of the benchmarked functions before running.

package main

import "testing"

func BenchmarkEx4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ex4()
	}
}

func BenchmarkEx5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ex5()
	}
}
