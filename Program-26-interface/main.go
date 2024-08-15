package main

import "fmt"

type sample interface {
	calling()
}

type person struct {
	f_name string
}

type agent struct {
	f_name string
}

func (r person) calling() {
	fmt.Println(r.f_name)

}

func (a agent) calling() {
	fmt.Println(a.f_name)
}

func inter_func(s sample) {
	s.calling()
}

func main() {
	p1 := person{
		f_name: "Ajay",
	}

	a1 := agent{
		f_name: "Tina",
	}

	//way of calling
	p1.calling()
	a1.calling()

	// another way of calling
	inter_func(p1)
	inter_func(a1)

}
