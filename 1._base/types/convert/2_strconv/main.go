package main

import (
	"fmt"
	"strconv"
)

func Sprintf(n int) {

}
func square(n int) int {
	return n * n
}
func main() {
	n := 97
	s := strconv.Itoa(n) //整數轉字串
	fmt.Printf("%T, %v \n", s, s)

	i := int64(97)
	sss := strconv.FormatInt(i, 10) //使用FormatInt整數轉字串，需要輸入型態為int64
	fmt.Printf("%T, %v \n", sss, sss)

	ss := string(n) //強轉 string函数的参数若是一个整型数字，它将该整型数字转换成ASCII码值等于该整形数字的字符
	fmt.Printf("%T, %v \n", ss, ss)
}
