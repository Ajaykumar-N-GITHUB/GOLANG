package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Age   int    `json:"Age"`
	Name  string `json:"Name"`
	Place string `json:"Place"`
}

func main() {

	s := `[{"Age":25,"Name":"Ajay","Place":"Banglore"},{"Age":22,"Name":"Jothika","Place":"Coimbatore"}]`
	bs := []byte(s)

	//fmt.Println(s)
	//fmt.Println(bs)

	people := []person{}
	//var people []person
	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println("Person Number ------>", i)
		fmt.Println("Details ----->", v.Age, v.Name, v.Place)
	}
}
