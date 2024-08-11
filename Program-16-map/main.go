package main

import "fmt"

func main() {
	//ranging oncer a Map
	m := make(map[string]int)

	m["Ajay"] = 25
	m["Jothika"] = 22
	m["Naren"] = 20

	for k, v := range m {
		fmt.Println(k, v)
	}
	for _, v := range m {
		fmt.Println(v)
	}
	for k := range m {
		fmt.Println(k)
	}

	//rangiing over a slice

	xi := make([]int, 0, 5)
	xi = append(xi, 11, 12, 13, 14, 15)

	for k, v := range xi {
		fmt.Println(k, v)
	}
	for _, v := range xi {
		fmt.Println(v)
	}
	for k := range xi {
		fmt.Println(k)
	}

	n := make(map[string]int)
	n["Ajay"] = 25
	n["Jothika"] = 22
	n["Naren"] = 20
	n["unknown"] = 12

	for k, v := range n {
		fmt.Println(k, v)
	}

	//delete a element in a map
	delete(n, "unknown")
	for k, v := range n {
		fmt.Println(k, v)
	}

	//acessing a value that doesn't exist
	fmt.Println(m["Hello"])

	csk := make(map[string]int)
	csk["Dhoni"] = 43
	csk["Jadeja"] = 45
	csk["Rutu"] = 28

	if v, ok := csk["Dhoni"]; ok {
		fmt.Println("Dhoni's Age is", v)
	} else {
		fmt.Println("Dhoni's Age is not present")
	}
	//n Go, when you use the range keyword with a string, it iterates over the string and provides the index and the Unicode code point (rune) of each character.

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for i, c := range "go" {
		fmt.Printf("Index: %d, Rune: %d, Character: %c\n", i, c, c)
	}

}
