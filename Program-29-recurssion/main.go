package main

import "fmt"

func main() {

	temp := 10
	fact := factorial(temp)
	fmt.Println(fact)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * (factorial(n - 1))
}
