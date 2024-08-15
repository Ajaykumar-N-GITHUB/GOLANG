package main

import "fmt"

func intDelta(v int) int {
	return v + 1
}

func pointDelta(k *int) int {
	*k = 21
	return *k

}

func main() {

	//value semantics
	a := 10
	fmt.Println(a)           //10
	fmt.Println(intDelta(a)) //11
	fmt.Println(a)           //10

	fmt.Println("___________________________")

	//pointer semantics
	b := 20
	fmt.Println(b)              //20
	fmt.Println(pointDelta(&b)) //21
	fmt.Println(b)              //21
}
