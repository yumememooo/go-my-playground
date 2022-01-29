package main

import (
	"fmt"
	"reflect"
)

type Member struct {
	id   int
	name string
	age  int
}

func main() {
	var myInt int = 230
	fmt.Println(reflect.TypeOf(myInt))  //output int
	fmt.Println(reflect.ValueOf(myInt)) //output 8
	fmt.Println(reflect.TypeOf(myInt).Size())
	member := Member{1, "Adam", 100}
	fmt.Println(reflect.TypeOf(member))  //output main.Member  //取得所有元素
	fmt.Println(reflect.ValueOf(member)) //output {1 Adam 100}/獲得值
	//TODO https://openhome.cc/Gossip/Go/Reflect.html 資料的 Kind NumField()
	//反射（Reflection）是探知資料自身結構的一種能力，不同的語言提供不同的反射機制
}
