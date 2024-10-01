package main

//hands-on-1
/*
func foo() int {
	return 2
}
func bar() (string, int) {
	return "Ajay", 25
}
func main() {

	a := foo()
	fmt.Println(a)

	name, age := bar()

	fmt.Println(name, age)
}
*/

//hands-on_2
//variadic
/*
func foo(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func bar(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func main() {

	ans1 := foo(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(ans1)
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans3 := foo(xi...)
	ans2 := bar(xi)
	fmt.Println(ans2)
	fmt.Println(ans3)
}
*/
//hands-on_3
//defer function
// It will run on LIFO order
/*
func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println("Ajay")
	defer fmt.Println(1)
	fmt.Println("kumar")
}
*/

//hands-on_4
//method
/*
type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Name -> %v and age -> %v ", p.first, p.age)
}

func main() {

	p1 := person{
		first: "Ajay",
		age:   25,
	}
	p1.speak()
}
*/

//hands-on_5
//interface
/*
type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(sh shape) float64 {
	return sh.area()
}

func main() {
	s := square{
		length: 10,
		width:  2,
	}

	c := circle{
		radius: 4,
	}

	fmt.Println(info(s))
	fmt.Println(info(c))
}
*/
