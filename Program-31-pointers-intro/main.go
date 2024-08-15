package main

import "fmt"

func main() {
	a := 20
	fmt.Println(a)

	// b is holding the address of a
	b := &a
	fmt.Println(b)

	// * operator will print the value inside which is present inside the address
	fmt.Println(*b)

	// here we are changing the value of b it will change the value of a also because both the variable are sharing the same memory location
	*b = 30
	fmt.Println(&b)

	//now value of a will also be change to 30
	fmt.Println(a)
}
