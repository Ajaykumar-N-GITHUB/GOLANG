//In Go (Golang), you have type conversion, but there is not  a concept of type casting as you might find in some other languages
//like C++ or Java. Here is how type conversion works in Go:

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 1
	b := 2.0
	var c string = "123"

	fmt.Println("Before type conversion")

	fmt.Printf("Value is %v and the type is %T\n", a, a)
	fmt.Printf("Value is %v and the type is %T\n", b, b)
	fmt.Printf("Value is %v and the type is %T\n", c, c)

	d := float64(a)
	e := int(b)

	fmt.Println("After type conversion")

	fmt.Printf("Value is %v and the type is %T\n", d, d)
	fmt.Printf("Value is %v and the type is %T\n", e, e)

	f, err := strconv.Atoi(c)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value is %v and the type is %T\n", f, f)
	}

}
