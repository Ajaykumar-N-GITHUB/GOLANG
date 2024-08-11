package main

import "fmt"

func main() {

	var x [3]int
	x[0] = 1
	x[1] = 2
	x[2] = 3

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	xi := [...]int{7, 6, 5}

	for _, v := range xi {
		fmt.Println(v)
	}

	xs := [3]string{"Ajay", "kumar", "N"}

	for i, v := range xs {
		fmt.Printf("Index - %v \t Value - %v\n", i, v)
	}

}
