package main

import "fmt"

func main() {
	var personAgeMap1 = map[string]int{
		"Apple": 25,
		"James": 32,
		"Sarah": 29,
	}
	age, ok := personAgeMap1["Apple"]
	fmt.Println("age:", age, ",ok", ok)

	var m2 = personAgeMap1 //，它們都參考相同的底層資料結構
	fmt.Println("--ori.---")
	fmt.Println("personAgeMap1 = ", personAgeMap1)
	fmt.Println("m2 = ", m2)
	fmt.Println("-modify personAgeMap1=18----")
	for name, age := range personAgeMap1 {
		fmt.Println(name, age)
		personAgeMap1[name] = 18
	}

	fmt.Println("personAgeMap1 = ", personAgeMap1)
	fmt.Println("m2 = ", m2)
	fmt.Println("-modify add m2[Amy] = 10----")
	m2["Amy"] = 10
	fmt.Println("personAge = ", personAgeMap1)
	fmt.Println("m2 = ", m2)
	fmt.Println("--modifymap(personAgeMap1)---")
	modifymap(personAgeMap1)
	fmt.Println("personAge = ", personAgeMap1)
	fmt.Println("m2 = ", m2)
	
}

func modifymap(personAge map[string]int) {
	for name := range personAge {
		personAge[name] = 20
	}
	delete(personAge, "Apple")
}
