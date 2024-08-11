package main

import "fmt"

type person struct {
	f_name string
	l_name string
	age    int
}

func main() {

	s1 := person{
		f_name: "Ajay",
		l_name: "Kumar",
		age:    25,
	}

	fmt.Println(s1)
	fmt.Println(s1.f_name)
	fmt.Println(s1.l_name)
	fmt.Println(s1.age)

	//anonymous struct

	s2 := struct {
		f_name string
		l_name string
		age    int
	}{
		f_name: "colour",
		l_name: "govind",
		age:    22,
	}

	fmt.Println(s2)
	fmt.Println(s2.f_name)
	fmt.Println(s2.l_name)
	fmt.Println(s2.age)

}
