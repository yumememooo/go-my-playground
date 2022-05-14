package main

import "fmt"

func main() {

	var myInt1 int = 123
	var myInt2 int = 123.0
	//var myInt3 int = 123.1//這樣是不行的 cannot use 123.1 (untyped float constant) as int value in variable declaration (truncated) compiler TruncatedFloat
	var myFloat1 float64 = 1

	fmt.Printf("%T %v \n", myInt1, myInt1)//int 123 
	fmt.Printf("%T %v \n", myInt2, myInt2)//int 123 
	fmt.Printf("%T %v \n", myFloat1, myFloat1)//float64 1 
}
