package main

import "fmt"

type company struct {
	name string
}

func changename(c1 company, s string) company {
	c1.name = s
	return c1
}

func changenameptr(c2 *company, s1 string) {
	c2.name = s1
}

func changename_noreturn(c3 company, s string) {
	c3.name = s
	fmt.Println("this is from changename_noreturn inside print statement") //since it is getting stored in stack once the process comess out of this funtion the memory will get vanish
	fmt.Println(c3)
}

func main() {

	str := "Xoriant"
	c1 := company{
		name: "maplelabs",
	}

	c2 := company{
		name: "maple",
	}

	c3 := company{
		name: "maplelabs_solution",
	}

	fmt.Println(c1)
	c1 = changename(c1, str)
	fmt.Println(c1)
	fmt.Println("__________________________________")
	fmt.Println(c2)
	changenameptr(&c2, str)
	fmt.Println(c2)
	fmt.Println("__________________________________")
	fmt.Println(c3)
	changename_noreturn(c3, str)
	fmt.Println(c3)

}
