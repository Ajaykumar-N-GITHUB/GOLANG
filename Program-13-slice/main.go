package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Println("length of the array is ", len(xs))

	fmt.Println(xs[1:])
	fmt.Println(xs[3:8])
	fmt.Println(xs[:6])
	fmt.Println(xs[:])

	//appending to an slice
	xs = append(xs, 10, 11, 12)
	fmt.Println(xs[:])

	//deleting an element in a slice

	xs = append(xs[:5], xs[6:]...)

	fmt.Println(xs[:])

}
