package main

import "fmt"

type User struct {
	Name string
}

type UserStru2 struct {
	Name string
}

func main() {
	user1 := User{Name: "test"}
	fmt.Printf("%T %v \n", user1, user1) //main.User {test}

	user2 := User(UserStru2{Name: "test"}) //T(v)
	fmt.Printf("%T %v \n", user2, user2)//main.User {test}
	
}
