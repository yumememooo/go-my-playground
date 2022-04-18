package main

import "fmt"

func ModifyArray(arr [5]int) {
	arr[0] = 5
	fmt.Println(arr) // [5,0,0,0,1]
}

func main() {
	a := [...]int{4: 1} //陣列為按值傳遞的，函式內對陣列的值的改變不影響初始陣列
	ModifyArray(a)
	fmt.Println(a) // [0,0,0,0,1]
	b := a
	b[0] = 88
	fmt.Println("b:", b) //修改b也不會影響到b
	fmt.Println("a:", a)

	fmt.Println("mo arr")
	arr := [5]int{1, 2, 3, 4, 5}

	for i, e := range arr {
		fmt.Println(fmt.Sprintf("%d: %d", i+1, e))
		e = e * e //不可以使用range中修改元素
	}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i] //若要修改陣列中的元素，要以索引走訪陣列，再修改陣列的元素的值
	}
	fmt.Println(arr)
	//arr = append(arr, 1)//invalid argument: arr (variable of type [5]int) is not a slice
}
