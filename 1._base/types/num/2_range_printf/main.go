package main

import (
	"fmt"

)


//%b(二進制表示)
func main() {
	v := []byte{127, 255}
	fmt.Printf("Type: %T v: %v\n", v, v)
	for _, n := range v {
		fmt.Printf("%v, % 08b [ 08b]\n", n, n) 
		fmt.Printf("%v, %08b [08b]\n", n, n)   
		fmt.Printf("%v, %b [b]\n", n, n)       
		fmt.Printf("%v, %#b [#b]\n", n, n)     
	}
}
