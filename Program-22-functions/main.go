package main

import "fmt"

func main() {
	foo()
	bar("Hello")
	s1 := foo_bar("Hello Gophers!!!")
	fmt.Println(s1)
	s2, s3 := aloha("Ajay", "Kumar")
	fmt.Println(s2, s3)
}

func foo() {
	fmt.Println("I am having no parameter and no return statement")
}

func bar(s string) {
	fmt.Println("I am having one parameter but no return statement")
}

func foo_bar(s string) string {
	fmt.Println("I am having one parameter and one return statement")
	return fmt.Sprintln("the return statement is ", s)
}

func aloha(i string, j string) (string, string) {
	fmt.Println("I am having two parameters and 2 return statements")
	return i, j
}
