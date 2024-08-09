package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Ths is where the initialization starts!!")
}

func main() {

	x := rand.Intn(250)
	fmt.Printf("Value is %v and the type is %T\n", x, x)

	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x <= 200:
		fmt.Println("between 101 and 200")
	case x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("More than 250")
	}

	y := rand.Intn(10)
	z := rand.Intn(10)

	fmt.Printf("Value is %v and Type is %T\n", y, y)
	fmt.Printf("Value is %v and type is %T\n", z, z)

	switch {
	case y < 4 && z < 4:
		fmt.Println("Both are Less then 4")
	case y > 6 && z > 6:
		fmt.Println("Both are greater than 6")
	case y >= 4 && z <= 6:
		fmt.Println("greater than or equal to 4 and less then eqal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("None of the previous case were met")

	}

}
