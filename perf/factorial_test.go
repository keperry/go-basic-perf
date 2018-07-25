package perf

import (
	"fmt"
	"testing"
)

func benchmarkFact (b *testing.B, num int) {
	for x := 0; x < b.N; x++ {
		//run my awesome test method
		x := Factorial(num)
		fmt.Printf("X = %d\n", x)
	}
}

func BenchmarkFact5(b *testing.B) {
	benchmarkFact(b, 5)
}
func BenchmarkFact10(b *testing.B) {
	benchmarkFact(b, 10)
}
func BenchmarkFact15(b *testing.B) {
	benchmarkFact(b, 15)
}
func BenchmarkFact20(b *testing.B) {
	benchmarkFact(b, 20)
}
