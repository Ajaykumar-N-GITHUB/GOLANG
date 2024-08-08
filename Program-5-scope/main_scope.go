package main

import (
	"fmt"
	"golang-project/Program-5-scope/explorer"
)

// package level scope

var a int = 10

func main() {

	hello()
	fmt.Println(a)

	explorer.Training()

	fmt.Println(explorer.Top_Spped)

	fmt.Println("This is from another package")

}

func hello() {

	// block level scope
	z := 7
	fmt.Println(z)

	fmt.Println(a)
}
