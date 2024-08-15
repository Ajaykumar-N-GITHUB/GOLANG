package main

import "fmt"

func intDelta(n *int) {
	*n = 30
}

func sliceDelta(ii []int) []int {
	ii[0] = 7
	return ii
}

func mapDelta(mp map[string]int, s string) {
	mp[s] = 26
}

func main() {

	a := 20
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)
	fmt.Println(sliceDelta(xi))
	fmt.Println(xi)

	m := make(map[string]int)
	m["Ajay"] = 25
	fmt.Println(m)
	mapDelta(m, "Ajay")
	fmt.Println(m)

}
