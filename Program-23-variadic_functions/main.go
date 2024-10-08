package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6}
	sum := variadic("Ajay", xi...)
	sum1 := variadic("kumar", 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println("The Total is", sum)
	fmt.Println(sum1)

}

//variadic parameter should be the last parameter

func variadic(s string, ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
