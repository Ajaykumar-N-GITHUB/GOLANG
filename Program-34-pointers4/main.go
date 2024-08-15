package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Printf("I am %v and i am walking\n", d.first)
}

func (d *dog) run() {
	fmt.Printf("I am %v and i am running\n", d.first)
}

type excercise interface {
	walk()
	run()
}

func test(e excercise) {
	e.run()
}

func main() {

	d := dog{
		first: "Tommy",
	}

	d1 := &dog{
		first: "Tommy",
	}

	d.run()
	d.walk()
	//test(d)  in interface we cannot pass this. if you are passing value then receiver should has only value based

	d1.run()
	d1.walk()
	test(d1) // but if we are sending pointer then receiver can be of both pointer receiver or normal reciver

}
