package main

import (
	"encoding/json"
	"fmt"
)

type sample struct {
	Fname   string // we need to specify first letter as caps
	Age     int
	Address string
}

func main() {

	s1 := sample{
		Fname:   "Ajay",
		Age:     25,
		Address: "banglore",
	}

	s2 := sample{
		Fname:   "jothika",
		Age:     22,
		Address: "Coimbatore",
	}

	fmt.Println("Before marshelling")
	fmt.Println(s1)
	fmt.Println(s2)

	users := []sample{s1, s2}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
