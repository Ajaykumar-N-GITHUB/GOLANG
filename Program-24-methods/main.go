package main

import "fmt"

type person struct {
	first string
}

// if we attach a struct type to a function then we will call it as methods
func (p person) test() string {
	return fmt.Sprintln("I am", p.first)
}

func main() {

	p1 := person{
		first: "Ajay",
	}

	s1 := p1.test()

	fmt.Println(s1)

}
