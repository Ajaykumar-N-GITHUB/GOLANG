package main

import "fmt"

type call interface {
	calling()
	called()
}

type person struct {
	f_name string
}

type agent struct {
	person
	ltk bool
}

func (r person) calling() {
	fmt.Println(r.f_name)
}

func (r1 agent) called() {
	fmt.Println(r1.person.f_name)
	fmt.Println(r1.ltk)
}

func temp(c call) {
	fmt.Println(c.calling())
	fmt.Println(c.called())

}

func main() {
	p1 := person{
		f_name: "Ajay",
	}

	a1 := agent{
		person: person{
			f_name: "kumar",
		},
		ltk: true,
	}

	temp(p1)
	temp(a1)

}
