package main

import (
	"fmt"
	"math"
)

//%b(二進制表示)
func main() {
	fmt.Printf(" MinInt: %v\n MaxInt: %v\n", math.MinInt, math.MaxInt)
	fmt.Printf(" MinInt8: %v\n MaxInt8: %v\n", math.MinInt8, math.MaxInt8)
	fmt.Printf(" MinInt16: %v\n MaxInt16: %v\n", math.MinInt16, math.MaxInt16)
	fmt.Printf(" MinInt32: %v\n MaxInt32: %v\n", math.MinInt32, math.MaxInt32)
	fmt.Printf(" MinInt64: %v\n MaxInt64: %v\n", math.MinInt64, math.MaxInt64)

	fmt.Printf(" MinUint: %v MaxUint: %v\n", 0, uint(math.MaxUint))  //不能直接印https://www.796t.com/post/MTVxNG8=.html
	fmt.Printf(" MaxUint8: %v MaxUint8: %v\n", 0, math.MaxUint8)
	fmt.Printf(" MaxUint16: %v MaxUint16: %v\n",0, math.MaxUint16)
	fmt.Printf(" MaxUint32: %v MaxUint32: %v\n", 0, math.MaxUint32)
	fmt.Printf(" MaxUint64: %v MaxUint64: %v\n", 0, uint64(math.MaxUint64))



	fmt.Printf("Type: %T MaxFloat32: %v\n", math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("Type: %T MaxFloat64: %v\n", math.MaxFloat64, math.MaxFloat64)
}

//【GO】為什麼打印出最大整數會導致golang編譯錯誤？
//不能直接印https://www.796t.com/post/MTVxNG8=.html