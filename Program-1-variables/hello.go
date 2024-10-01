package main

import "fmt"

func main() {
	var name = "Ajay"
	var age int = 25

	fmt.Printf("This is %s and i am %d years old. \t The type of parameter used in this Print function are %T and %T\n", name, age, name, age)

	a := 112
	fmt.Println(a / 10)
	b := 845
	fmt.Println(b % 10)

	b, c, _, e := 1, 2.0, 3, "Numbers"
	fmt.Println(b, c, e)
}
