package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1            //??一般來講不知道要這樣寫，math.MaxUint64有定義
	z      complex128 = cmplx.Sqrt(-5 + 12i) //複數
)

//%b(二進制表示)
func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("Type: %T MaxInt: %v\n", math.MaxInt, math.MaxInt)
	fmt.Printf("Type: %T MinInt: %v\n", math.MinInt, math.MinInt)
	//int8可接受負整數，並用最左邊的位元來代表正負號，所以說int8的最大值會比uint8小(因為少一個bit)。
	fmt.Printf("Type: %T MaxInt8 : %v ,binary:%b\n", math.MaxInt8, math.MaxInt8, math.MaxInt8)
	fmt.Printf("Type: %T MinInt8 : %v\n", math.MinInt8, math.MinInt8)
	fmt.Printf("Type: %T MaxInt16: %v\n", math.MaxInt16, math.MaxInt16)
	fmt.Printf("Type: %T MinInt16: %v\n", math.MinInt16, math.MinInt16)
	fmt.Printf("Type: %T MaxInt32: %v\n", math.MaxInt32, math.MaxInt32)
	fmt.Printf("Type: %T MinInt32: %v\n", math.MinInt32, math.MinInt32)
	fmt.Printf("Type: %T MaxInt64: %v\n", math.MaxInt64, math.MaxInt64)
	fmt.Printf("Type: %T MinInt64: %v\n", math.MinInt64, math.MinInt64)

	var s uint = math.MaxUint //如果不指定直接印會out int
	fmt.Printf("Type: %T MaxUint  : %v\n", s, s)
	fmt.Printf("Type: %T MaxUint8 : %v, binary:%b\n", math.MaxUint8, math.MaxUint8, math.MaxUint8)
	fmt.Printf("Type: %T MaxUint16: %v, binary: %b\n", math.MaxUint16, math.MaxUint16, math.MaxUint16)
	fmt.Printf("Type: %T MaxUint32: %v, binary: %b\n", math.MaxUint32, math.MaxUint32, math.MaxUint32)
	var uu uint64 = math.MaxUint
	fmt.Printf("Type: %T MaxUint64: %v\n", uu, uu)

	fmt.Printf("Type: %T MaxFloat32: %v\n", math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("Type: %T MaxFloat64: %v\n", math.MaxFloat64, math.MaxFloat64)

	bb := 2 ^ 8
	fmt.Printf("Type: %T XOR(^): %d\n", bb, bb)

	bb2 := math.Pow(2, 8)
	fmt.Printf("Type: %T 2的八次方: %v\n", bb2, bb2) //uint8可以存256種可能(數字0 ~ 255)

	v := []byte{127, 255, math.MaxUint8}
	fmt.Printf("Type: %T v: %v\n", v, v)
	for _, n := range v {
		fmt.Printf("%v, % 08b [ 08b]\n", n, n) // prints 00000000 11111101
		fmt.Printf("%v, %08b [08b]\n", n, n)   // prints 00000000 11111101
		fmt.Printf("%v, %b [b]\n", n, n)       // prints 00000000 11111101
		fmt.Printf("%v, %#b [#b]\n", n, n)     // prints 00000000 11111101
	}
}
