package main

import (
	"fmt"
)

func main() {
	langs := []string{"Go", "Python", "Ruby", "PHP"}

	for _, e := range langs {
		e = e + "xxx"
		fmt.Println(e)
	}
	
	fmt.Printf("Type: %T ,langs: %v\n", langs, langs)
}
