package main

import "fmt"

//Go is statically typed language and does not permit operation that mix the numeric types

func main() {

	const Pi = 3.14
	const Greeting = "Hello, world!"
	fmt.Println(Pi)
	fmt.Println(Greeting)

	const x = 42
	var y int = x // Untyped constant assigned to a variable of type int
	fmt.Println(y)

	const (
		Zero = 0
		One  = 1
		Two  = 2
	)

	fmt.Println(Zero)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		A = 10
		B = A * 2
		C = B + 5
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	//iota is a special identifier used in Go to simplify the creation of incrementing constants. It’s often used for enumerations or flags:
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday)    // Output: 0
	fmt.Println(Monday)    // Output: 1
	fmt.Println(Tuesday)   // Output: 2
	fmt.Println(Wednesday) // Output: 3
	fmt.Println(Thursday)  // Output: 4
	fmt.Println(Friday)    // Output: 5
	fmt.Println(Saturday)  // Output: 6

	const (
		Base     = 2
		Exponent = 3
		Result   = Base << Exponent // Bitwise shift operation
	)

	fmt.Println(Result)
}
