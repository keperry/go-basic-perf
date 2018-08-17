package perf

import (
	"testing"
	)

var result int

func benchmarkFact(b *testing.B, num int) {
	var y int
	for x := 0; x < b.N; x++ {
		//run my awesome test method
		y = Factorial(num)
		//fmt.Printf("Y = %d\n", y)
	}
	result = y
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
