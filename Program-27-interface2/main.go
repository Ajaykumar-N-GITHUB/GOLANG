package main

import "fmt"

type inter interface {
	print_first()
	print_last()
	print_address()
}

type temp struct {
	f_name  string
	l_name  string
	address string
}

func (t temp) print_first() {
	fmt.Println(t.f_name)
}

func (t temp) print_last() {
	fmt.Println(t.l_name)
}

func (t temp) print_address() {
	fmt.Println(t.address)
}

func inter_func(i inter) {
	fmt.Println("Common")
	i.print_first()
	i.print_last()
	i.print_address()
}

func main() {

	t1 := temp{
		f_name:  "Ajay",
		l_name:  "Kumar",
		address: "Banglore",
	}

	// t1.print_first()
	// t1.print_last()
	// t1.print_address()

	inter_func(t1)

}
