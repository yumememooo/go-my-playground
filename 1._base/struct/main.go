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
	fmt.Printf("%v name=%s\n", user2, user2.Name) // &{}，user2.Name 會是 ""
	fmt.Printf("%v name=%s\n", user3, user3.Name) // &{}，user3.Name 會是 ""
	fmt.Printf("%v name=%s\n", user4, user4.Name) // {} user4.Name 會是 ""
	fmt.Printf("%v name=%s\n", user5, user5.Name) // {} user5.Name 會是 ""

	//匿名的 struct（anonymous struct）
	Person1 := struct {
		NickName string
	}{
		NickName: "NANA",
	}
	fmt.Printf("%v NickName=%s\n", Person1, Person1.NickName) //{NANA} NickName=NANA

	//匿名欄位會把型別用作變數名
	var emp1 employee
	emp1.firstName = "Andy"
	emp1.lastName = "wu"
	emp1.age = 20
	p := MyPerson{"sort1", 1, emp1}
	p.string = "team1"
	fmt.Println(p, p.string) //{team1 1 {Andy wu 20}} team1
	fmt.Println(p.lastName)  //等價於p.employee.lastName//輸出：wu

	//巢狀結構體 nested struct
	P1 := PersonInfo{Name: "Jannet", AddrInfo: AddrInfo{Code: 723, Location: "Taipai"}}
	fmt.Printf("%v Location=%s\n", P1, P1.Location)

	em2 := employee{
		firstName: "Andy",
		lastName:  "Li",
		GetFullName: func(firstName string, lastName string) string {
			return firstName + " " + lastName //TODO TBD why??
		},
	}
	
	fmt.Printf("%v GetFullName=%s\n", em2, em2.GetFullName(em2.firstName, em2.lastName))
}

//匿名欄位宣告法
type MyPerson struct {
	string
	int
	employee
}
type employee struct {
	firstName   string
	lastName    string
	age         int
	GetFullName GetFullName
}

//nested struct
type PersonInfo struct {
	Name     string
	AddrInfo //匿名欄位
	// AddrInfo AddrInfo

}

type AddrInfo struct {
	Code     int
	Location string
}

//   function type
type GetFullName func(string, string) string
