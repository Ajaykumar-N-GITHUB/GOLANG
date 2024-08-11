package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Ths is where the initialization starts!!")
}

func main() {
	for i := 0; i < 100; i++ {
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
	var u int = 0
	for u < 10 {
		fmt.Println(u)
		u++
	}

	for {
		fmt.Println(u)
		if u > 10 {
			break
		}
		u++
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop =  %v and inner loop =  %v\n", i, j)
		}
		fmt.Println()
	}

}
