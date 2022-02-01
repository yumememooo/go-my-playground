package main

import (
	"fmt"
)

func modifySlice(s []string) {
	s[3] = "PHP_M"
}

func main() {
	langs := []string{"Go", "Python", "Ruby", "PHP"}

	s2 := langs
	s2[1] = "NA"
	fmt.Printf("angs: %v\n", langs) //[Go NA Ruby PHP]
	fmt.Printf("s2: %v\n", s2)      //按引用傳遞的 改變s2也會改變
	fmt.Print("-----modifySlice------\n")

	modifySlice(langs)
	fmt.Printf("angs: %v\n", langs)
	fmt.Printf("s2: %v\n", s2)

	fmt.Print("-----range------\n")
	for _, e := range langs {
		e = e + "xxx" //這樣改無效
		fmt.Println(e)
	}
	fmt.Printf("angs: %v\n", langs)
	for i, e := range langs {
		langs[i] = e + "iiii" //要用index改
		fmt.Println(e)
	}
	fmt.Printf("angs: %v\n", langs)
	for i := 0; i < len(langs); i++ {
		langs[i] = langs[i] + "xxxx"
	}
	fmt.Printf("angs: %v\n", langs)

}
