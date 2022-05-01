package main

import (
	"fmt"
	"time"
)

//如果使用for{go func()}`for循環+goroutine的坑`	的傳遞注意事項

func main() {
	testRoutine()     //踩到ｂｕｇ
	testRoutineFix1() //修正寫法1
	testRoutineFix2() //修正寫法2
	testRoutineFix3() //修正寫法3
}

func testRoutine() { //Golang中Routine闭包中的一个坑
	for i := 0; i < 10; i++ {
		go func() { //这是因为很有可能当 for-loop 执行完之后 goroutine 才开始执行，这个时候 val 的值指向切片中最后一个元素。
			fmt.Println("testRoutine", i) //值可能是0-100中的任意数字的，有一些可能会重复出现-->引用传递（同一个对象）给了子协程
		}()
	}
	time.Sleep(1 * time.Second)
}

func testRoutineFix1() {
	for i := 0; i < 10; i++ {
		go print(i)
	}
	time.Sleep(1 * time.Second)
}

func print(i int) {
	fmt.Println("testRoutineFix1", i)
}
func testRoutineFix2() {
	for i := 0; i < 10; i++ {
		go func(ii int) {
			fmt.Println("testRoutineFix2", ii)
		}(i) //参数都是通过值传递进行传递的
	}
	time.Sleep(1 * time.Second)
}

func testRoutineFix3() {
	for i := 0; i < 10; i++ {
		ii := i
		go func() {
			fmt.Println("testRoutineFix3", ii)
		}()
	}
	time.Sleep(1 * time.Second)
}
