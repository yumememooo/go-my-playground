package main

import (
	"fmt"
)

func main() {
	var myInt8 int8 = 88
	myInt64 := int64(myInt8)
	fmt.Printf("Type: %T myInt64: %v\n", myInt64, myInt64)

	myInt8_2 := int8(myInt64)
	fmt.Printf("Type: %T myInt8_2: %v\n", myInt8_2, myInt8_2)

	var myInt8_3 int8 = -88
	myInt8_4 := uint8(myInt8_3)
	fmt.Printf("Type: %T myInt8_2: %v\n", myInt8_4, myInt8_4)
}
