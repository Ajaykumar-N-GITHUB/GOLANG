package main

import "fmt"

func addI(a int, b int) int {
	return a + b
}

func addF(c float64, d float64) float64 {
	return c + d
}

//we cann do like this as well
/*
type common interface{
	~int | ~float64
}

func mathT[X common](a X, b X) X {
	return a + b
}
*/

func mathT[X int | float64](a X, b X) X {
	return a + b
}

//type Alias int

func main() {

	// var n Alias = 2
	// fmt.Println(addI(n, 2))

	fmt.Println("Normal call")
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 3.4))

	fmt.Println("Generics call")
	fmt.Println(mathT(1, 2))
	fmt.Println(mathT(1.2, 3.4))

}
