package main

import "fmt"

func main() {
	var myInt int = 8
	accept(myInt)
	var myI interface{} = 8
	accept(myI)
	myI = "ssss"
	accept(myI)
}

//空的接口，interface{}，由于任何类型都至少实现了0个方法，所以空接口可以承接任意类型。
func accept(input interface{}) {

	fmt.Printf("type:%T \n", input) //印出interface type的用法
	fmt.Printf("value:%v  \n", input)

	//類型断言 Type Assertion（断言）是用于interface value的一种操作，语法是x.(T)，x是interface type的表达式，而T是asserted type
	// number := input.(int) 可以這樣寫 但如果斷言失敗會引發panic,panic: interface conversion: interface {} is string, not int
	number, ok := input.(int)
	if ok {
		fmt.Printf("number:%d \n", number)
		return
	}

	//類型判定 input.(type)
	//錯誤XXfmt.Println(input.(type)) ERR:use of .(type) outside type switch
	switch v := input.(type) { //神奇的是要用於switch才有辦法對到case
	case float64:
		fmt.Printf("input.(type) %T \n", v)
	case int:
		fmt.Print("input.(type) v=", v, "\n") //如果印v 會是值
		fmt.Printf("input.(type) %T \n", v)
	case string:
		fmt.Printf("input.(type) %T \n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}
