package main

import "fmt"

type person struct {
	f_name string
	l_name string
	age    int
}

type agent struct {
	person
	ltk bool
}

func main() {

	p1 := person{
		f_name: "Ajay",
		l_name: "kumar",
		age:    25,
	}
	sa := agent{
		person: p1,
		ltk:    true,
	}

	sa2 := agent{
		person: person{
			f_name: "colour",
			l_name: "govind",
			age:    23,
		},
		ltk: false,
	}

	fmt.Println(sa)
	fmt.Println(sa2)
	fmt.Println(sa.person.age)
	fmt.Println(sa2.person.age)

}
