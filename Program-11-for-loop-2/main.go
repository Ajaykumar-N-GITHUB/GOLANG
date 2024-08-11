package main

import (
	"fmt"
	"math/rand"
)

func main() {

	xi := []int{11, 12, 13, 14, 15}

	for k, v := range xi {
		fmt.Printf("Index is %v \t Value is %v\n", k, v)
	}

	m := map[string]int{
		"Ajay":  111,
		"Kumar": 110,
	}

	for k, v := range m {
		fmt.Printf("key is %v and the Value is %v\n", k, v)
	}

	Ajay_Roll_no := m["Ajay"]

	fmt.Println(Ajay_Roll_no)

	if v, ok := m["Ajay"]; ok {
		fmt.Println(v)
	}

	if v, ok := m["david"]; !ok {
		fmt.Println("There is no Rollno so returning", v)
	}

	for h := 0; h < 100; h++ {
		if g := rand.Intn(5); g == 3 {
			fmt.Printf("Iteration %v number is %v\n", h, g)
		}

	}

}
