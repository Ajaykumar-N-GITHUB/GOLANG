package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Age   int
	Name  string
	Place string
}

func main() {

	p1 := person{
		Age:   25,
		Name:  "Ajay",
		Place: "Banglore",
	}

	p2 := person{
		Age:   22,
		Name:  "Jothika",
		Place: "Coimbatore",
	}

	people := []person{p1, p2}

	// for _, v := range people {
	// 	fmt.Println(v)
	// }

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	// fmt.Println(bs)

}
