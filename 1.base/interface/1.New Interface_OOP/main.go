package main

import (
	"fmt"
	"strconv"
)


//定義str類型與其方法
type str string

func ( str) String() string {
	return string(s)
}

//定義Stringer interface
type Stringer interface {
	String() string
}

//person 與其方法
type person struct {
	Name string
}

func (s person) String() string {
	return s.Name
}
func (s person) MyName() string {
	return s.Name
}

// New person對象
func NewPerson(name string) person {
	return person{Name: name} //返回person對象
}

//宣告 person介面
func NewPersonI(name string) Person { //回傳是一個介面
	return person{Name: name} //會檢查對象person是否具有Person(interface) 的方法
}

type Person interface {
	String() string
	MyName() string
}

func showTypeAndDoTheirFunc(any interface{}) string {
	// if v, ok := any.(Person); ok {  //可以這樣寫 或是寫進switch case
	// 	fmt.Printf("It is Person.(type) %T \n", v)
	// 	return v.MyName()
	// }
	// if v, ok := any.(Stringer); ok {
	// 	fmt.Printf("It is Stringer.(type) %T \n", v)
	// 	return v.String()
	// }

	switch v := any.(type) {

	case SuperMan:
		fmt.Printf("It is SuperMan.(type) %T \n", v)
		return v.MyName()
	case Person:
		fmt.Printf("It is Person.(type) %T \n", v)
		return v.MyName()
	case Stringer:
		fmt.Printf("It is Stringer.(type) %T \n", v) //Stringer擁有Person其中一樣方法，不能寫在Person前
		return v.String()
	case int:
		fmt.Printf("It is int.(type) %T \n", v)
		return strconv.Itoa(v)
	case float64:
		fmt.Printf("It is float64.(type) %T \n", v)
		return strconv.FormatFloat(v, 'g', -1, 64)
	}
	return "???"
}

func main() {
	var ex2 float64 = 0.1
	fmt.Println(showTypeAndDoTheirFunc(ex2))
	fmt.Printf("-------\n")

	var ex3 Stringer = str("1")
	fmt.Println(showTypeAndDoTheirFunc(ex3))
	fmt.Printf("-------\n")

	var p1 Person = person{Name: "YU_I"} //使用介面
	fmt.Println(showTypeAndDoTheirFunc(p1))
	var p2 = NewPersonI("ALICE_I") //利用func New來指定介面寫法
	fmt.Println(showTypeAndDoTheirFunc(p2))
	fmt.Printf("-------\n")

	var p4 = person{Name: "YU"} //不使用介面
	fmt.Println(showTypeAndDoTheirFunc(p4))
	var p3 = NewPerson("ALICE") //利用func New來指定寫法
	fmt.Println(showTypeAndDoTheirFunc(p3))

	fmt.Printf("-------\n")

	var sm1 = NewSuperManI("xMan") //利用func New來指定介面寫法
	fmt.Println(showTypeAndDoTheirFunc(sm1))
	fmt.Printf("-------\n")

}

func NewSuperManI(name string) SuperMan { //回傳是一個介面
	return superMan{Name: name} //會檢查對象person是否具有Person(interface) 的方法
}

type SuperMan interface {
	Fly() string
	MyName() string
}

type superMan struct {
	Name string
}

func (s superMan) Fly() string {
	return s.Name
}
func (s superMan) MyName() string {
	return s.Name
}
