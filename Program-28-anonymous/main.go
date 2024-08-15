package main

import "fmt"

func foo() {
	fmt.Println("Function Foo!!!")
}

func main() {

	foo()

	//func expression
	x := func() {
		fmt.Println("Function Expression print sttement")
	}
	x()

	func() {
		fmt.Println("Anonymous call")
	}()
	func(s string) {
		fmt.Println("Anonymous functions calling a name ", s)
	}("Ajay")

}
