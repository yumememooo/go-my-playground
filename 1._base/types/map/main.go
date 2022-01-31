package main

import "fmt"

func main() {
	var m map[string]int //語法宣告 map ///map 的零值是 nil
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}
	// m["one hundred"] = 100 //panic: assignment to entry in nil map

	var m2 = make(map[string]int) //使用內建的 make() 函數初始化 map

	fmt.Println(m2)

	if m2 == nil {
		fmt.Println("m2 is nil")
	} else {
		fmt.Println("m2 is not nil")
	}

	m2["one hundred"] = 100
	fmt.Println(m2)

	var m3 = map[string]int{} //可以透過將大括號留空，使用 map 定數來建立一個空 map
	if m3 == nil {
		fmt.Println("m3 is nil")
	} else {
		fmt.Println("m3 is not nil")
	}
	m3["one hundred"] = 100
	fmt.Println(m3)
}
