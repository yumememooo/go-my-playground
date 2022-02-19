package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	var user1 *Person
	user2 := &Person{}
	user3 := new(Person)
	var user4 Person
	user5 := Person{}

	fmt.Printf("%v \n", user1) // nil
	// fmt.Printf("%v name=%s\n", user1,user1.Name) 
	//panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("%v name=%s\n", user2,user2.Name)  // &{}，user2.Name 會是 ""
	fmt.Printf("%v name=%s\n", user3,user3.Name)  // &{}，user3.Name 會是 ""
	fmt.Printf("%v name=%s\n", user4,user4.Name) // {} user4.Name 會是 ""
	fmt.Printf("%v name=%s\n", user5,user5.Name) // {} user5.Name 會是 ""

}
