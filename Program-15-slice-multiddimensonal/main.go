package main

import "fmt"

func main() {

	xa := []string{"a", "b", "c", "d"}
	xb := []string{"A", "B", "C", "D"}

	fmt.Println(xa)
	fmt.Println(xb)

	xab := [][]string{xa, xb}

	fmt.Println(xab)
	fmt.Println("------------------------------------------------------")

	fmt.Println("Difference between assign and copy")
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a)
	fmt.Println(b)

	a[0] = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("=======================================================")
	//c := []int{4, 5, 6}
	c := make([]int, 0, 3)
	d := make([]int, 3)
	c = append(c, 4, 5, 6)

	fmt.Println(c)

	copy(d, c)
	c[0] = 20

	fmt.Println(c)
	fmt.Println(d)

}
