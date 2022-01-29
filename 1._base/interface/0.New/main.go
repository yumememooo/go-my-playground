package main

import (
	"fmt"
	"strconv"
)

type str string

func (s str) String() string {
	return string(s)
}

type Stringer interface {
	String() string
}

type person struct {
	Name string
}

func (s person) String() string {
	return s.Name
}
func (s person) MyName() string {
	return s.Name
}

func NewPerson(name string) person {
	return person{Name: name} //返回person對象
}

func NewPersonI(name string) Person {
	return person{Name: name} //會檢查對象person是否具有Person(interface) 的方法
}

type Person interface {
	String() string
	MyName() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	if v, ok := any.(Person); ok {
		return v.MyName()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	}
	return "???"
}

func main() {
	var ex2 float64 = 0.1
	fmt.Println(ToString(ex2))

	var ex3 Stringer = str("1")
	fmt.Println(ToString(ex3))

	var p1 Person = person{Name: "YU_I"}
	fmt.Println(ToString(p1))
	var p2 = NewPersonI("ALICE_I")
	fmt.Println(ToString(p2))

	var p3 = NewPerson("ALICE")
	fmt.Println(ToString(p3))
	var p4 = person{Name: "YU"}
	fmt.Println(ToString(p4))

}
