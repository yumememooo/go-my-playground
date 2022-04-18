package main

import (
	"fmt"
	"strconv"
	"testing"
)

//go test main_test.go main.go
//go test .  XX???
func Test_square(t *testing.T) {
	t.Log("Test successful")
}

//go test -v -bench=. main_test.go main.go -benchmem
// go test -v -bench=. -run=none . XXX
func BenchmarkSprintf(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", n)
		//	Sprintf(n)
	}
}
func BenchmarkItoa(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(n)
	}
}
func BenchmarkFormatInt(b *testing.B) {
	n := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(n, 10)
	}
}

//resulte BenchmarkSprint最差，但後兩者看起來不一定@@
// goos: windows
// goarch: amd64
// cpu: Intel(R) Core(TM) M-5Y10c CPU @ 0.80GHz
// BenchmarkSprintf
// BenchmarkSprintf-4       8744379               140.1 ns/op             2 B/op          1 allocs/op
// BenchmarkItoa
// BenchmarkItoa-4         222285202                5.142 ns/op           0 B/op          0 allocs/op
// BenchmarkFormatInt
// BenchmarkFormatInt-4    232986952                5.359 ns/op           0 B/op          0 allocs/op
