package main

import "fmt"

//轉字串轉字串
func main() {
	type Person struct {
		Name string
	}
	var people = Person{Name: "mark"}

	//1.普通占位符
	//%v(相應值的默認格式)
	fmt.Printf("%v", people) //{mark}

	//%+v(打印結構體時，會添加字段名)
	fmt.Printf("%+v", people) //{Name:mark}

	//%#v(相應值的Go語法表示)
	fmt.Printf("%#v", people) //main.Person{Name:"mark"}

	//%T(相應值的類型的Go語法表示)
	fmt.Printf("%T", people) //main.Person

	//%%(字面上的百分號，並非值的占位符)
	fmt.Printf("%%") //%

	//2.布爾占位符
	//%t(true 或 false)
	fmt.Printf("%t", true) //true

	//3.整數占位符
	//%b(二進制表示)
	fmt.Printf("%b", 5) //101

	//%c(相應Unicode碼點所表示的字符)
	fmt.Printf("%c", 0x4E2D) //中

	//%d(十進制表示)
	fmt.Printf("%d", 0x12) //18

	//%o(八進制表示)
	fmt.Printf("%o", 10) //12

	//%q(單引號圍繞的字符字面值，由Go語法安全地轉義)
	fmt.Printf("%q", 0x4E2D) //'中'

	//%x(十六進制表示，字母形式為小寫a-f)
	fmt.Printf("%x", 13) //d

	//%X(十六進制表示，字母形式為小寫A-F)
	fmt.Printf("%X", 13) //D

	//%U(Unicode格式：U+1234，等同於 "U+%04X")
	fmt.Printf("%U", 0x4E2D) //U+4E2D

	//4.浮點數和覆數的組成部分
	//%b(無小數部分的，指數為二的冪的科學計數法)
	fmt.Printf("%b", 10.2) //5742089524897382p-49

	//%e(科學計數法,例如 -1234.456e+78)
	fmt.Printf("%e", 10.2) //1.020000e+01

	//%E(科學計數法,例如 -1234.456E+78)
	fmt.Printf("%E", 10.2) //1.020000E+01

	//%f(有小數點而無指數，例如123.456)
	fmt.Printf("%f", 10.2) //10.200000

	//%g(根據情況選擇%e或%f以產生更緊湊的(無末尾的0))
	fmt.Printf("%g", 10.20) //10.2

	//%G(根據情況選擇%E或%f以產生更緊湊的(無末尾的0))
	fmt.Printf("%G", 10.20+2i) //(10.2+2i)

	//5.字符串與字節切片
	//%s(輸出字符串表示(string類型或[]byte))
	fmt.Printf("%s", []byte("Go語言")) //Go語言

	//%q(雙引號圍繞的字符串，由Go語法安全地轉義)
	fmt.Printf("%q", "Go語言") //"Go語言"

	//%x(十六進制，小寫字母，每字節兩個字符)
	fmt.Printf("%x", "golang") //676f6c616e67

	//%X(十六進制，大寫字母，每字節兩個字符)
	fmt.Printf("%X", "golang") //676F6C616E67

	//6.指針
	//%p(十六進制表示，前綴0x)
	fmt.Printf("%p", &people) //0xc0420421d0
}
