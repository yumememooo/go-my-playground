package main

import (
	"fmt"
)

func main() {
	s := "23"
	fmt.Printf("Type: %T Value: %v\n", s, s)

	//byte 和 rune 資料型別本質上都是整數
	var firstLetter = 'A'
	fmt.Printf("Type: %T Value: %v\n", firstLetter, firstLetter)

	var myByte byte = 'a' //是 uint8 資料型別的別名。byte 資料型別用於表示 ASCII 字元
	var myRune rune = '♥' //是 int32 資料型別的別名。用於表示以 UTF-8 格式編碼的一組更廣泛的 Unicode 字元。
	fmt.Printf("myByte Type: %T Value: %v\n", myByte, myByte)
	//myByte Type: uint8 Value: 97
	fmt.Printf("myRune Type: %T Value: %v\n", myRune, myRune)
	//myRune Type: int32 Value: 9829
	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
}
