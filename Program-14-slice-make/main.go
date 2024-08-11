package main

import "fmt"

func main() {

	xi := make([]int, 0, 5)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("------------------")
	xi = append(xi, 1, 2, 3, 4, 5)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-------------------")
	xi = append(xi, 6, 6)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

}
