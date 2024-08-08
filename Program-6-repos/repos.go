package main

import (
	"fmt"

	"github.com/Ajaykumar-N-GITHUB/fish"
)

func main() {
	fmt.Println("This is for accessing code from different github repositories")
	s1 := fish.Fishing("Ajay")
	s2 := fish.NotFishing("Kumar")

	fmt.Println(s1)
	fmt.Println(s2)
}
