package main

import "fmt"

func main() {
	var name = "Ajay"
	var age int = 25

	fmt.Printf("This is %s and i am %d years old. \t The type of parameter used in this Print function are %T and %T", name, age, name, age)

	a := 1
	fmt.Println(a)
	a = 8
	fmt.Println(a)

	b, c, _, e := 1, 2.0, 3, "Numbers"
	fmt.Println(b, c, e)
}
