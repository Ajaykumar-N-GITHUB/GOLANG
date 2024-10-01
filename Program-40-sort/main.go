package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{11, 67, 107, 23, 45, 6, 10}

	xs := []string{"james", "Q", "Moneypenny", "Dr. No", "M"}

	fmt.Println("Before sortings")
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("After Sorting")

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
