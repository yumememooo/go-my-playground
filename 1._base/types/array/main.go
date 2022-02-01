package main

import (
	"fmt"
	"log"
)

func main() {
	var langs [4]string

	langs[0] = "Go"
	langs[1] = "Python"
	langs[2] = "Ruby"
	langs[3] = "PHP"
	fmt.Printf("Type: %T ,langs: %v\n", langs, langs)
}
