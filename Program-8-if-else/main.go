package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("Value is %v and the type is %T\n", x, x)

	if x >= 0 && x <= 100 {
		fmt.Println("value is inbetween 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Printf("Value is inbetween 101 and 200")
	} else if x >= 201 && x <= 300 {
		fmt.Printf("Value is inbetween 201 and 250")
	} else {
		fmt.Println("Value is more than 250")
	}

	// rand.intn(n)  will generate a random output from 0 to n-1
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

	y := rand.Intn(10)
	z := rand.Intn(10)

	fmt.Printf("Value is %v and Type is %T\n", y, y)
	fmt.Printf("Value is %v and type is %T\n", z, z)

	if y < 4 && z < 4 {
		fmt.Println("Both are Less then 4")
	} else if y > 6 && z > 6 {
		fmt.Println("Both are greater than 6")
	} else if y >= 4 && z <= 6 {
		fmt.Println("greater than or equal to 4 and less then eqal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("None of the previous case were met")
	}

}
