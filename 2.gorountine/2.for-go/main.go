package main

import (
	"fmt"
	"time"
)

func main() {
	//testRoutine()  //踩到ｂｕｇ
	//testRoutineFix1() //修正寫法1
	//testRoutineFix2() //修正寫法2
	testRoutineFix3()
}

func testRoutine() { //Golang中Routine闭包中的一个坑
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i) //值可能是0-100中的任意数字的，有一些可能会重复出现-->引用传递（同一个对象）给了子协程
		}()
	}
	time.Sleep(1 * time.Second)
}

func testRoutineFix1() {
	for i := 0; i < 100; i++ {
		go print(i)
	}
	time.Sleep(1 * time.Second)
}

func print(i int) {
	fmt.Println(i)
}
func testRoutineFix2() {
	for i := 0; i < 100; i++ {
		go func(ii int) {
			fmt.Println(ii)
		}(i) //参数都是通过值传递进行传递的
	}
	time.Sleep(1 * time.Second)
}

func testRoutineFix3() {
	for i:=0; i < 100; i++ {
			ii := i
			go func() {
					fmt.Println(ii)
			}()
	}
	time.Sleep(1 * time.Second)
}
