package main

import (
	"fmt"
	"math"
)

//%b(二進制表示)
func main() {
	//int8可接受負整數，並用最左邊的位元來代表正負號，所以說int8的最大值會比uint8小(因為少一個bit)。
	fmt.Printf("Type: %T MaxInt8 : %v ,binary:%b\n", math.MaxInt8, math.MaxInt8, math.MaxInt8)
	fmt.Printf("Type: %T MinInt8 : %v,binary:%b\n", math.MinInt8, math.MinInt8, math.MinInt8)
	fmt.Printf("Type: %T MaxUInt8 : %v ,binary:%b\n", math.MaxUint8, math.MaxUint8, math.MaxUint8)
	fmt.Printf("Type: %T MinUint8 : %v,binary:%b\n", 0, 0, 0)
	bb := 2 ^ 8
	fmt.Printf("Type: %T XOR(^): %d\n", bb, bb)

	bb2 := math.Pow(2, 8)
	fmt.Printf("Type: %T 2的八次方: %v\n", bb2, bb2) //uint8可以存256種可能(數字0 ~ 255)
}

// RESULT
// Type: int MaxInt8 : 127 ,binary:1111111
// Type: int MinInt8 : -128,binary:-10000000
// Type: int MaxUInt8 : 255 ,binary:11111111
// Type: int MinUint8 : 0,binary:0
// Type: int XOR(^): 10
// Type: float64 2的八次方: 256
