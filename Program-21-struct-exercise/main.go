package main

import "fmt"

/*
type engine struct {
	electric bool
}

type vechicle struct {
	engine
	make   string
	model  string
	door   int
	colour string
}

func main() {

	e1 := engine{
		electric: false,
	}

	v1 := vechicle{
		engine: e1,
		make:   "Germany",
		model:  "polo",
		door:   4,
		colour: "brown",
	}

	v2 := vechicle{
		engine: engine{
			electric: true,
		},
		make:   "India",
		model:  "nexon",
		door:   4,
		colour: "blue",
	}

	fmt.Println(v1)
	fmt.Println(v1.model, v1.make, v1.door, v1.engine.electric)

	fmt.Println(v2)
	fmt.Println(v2.model, v2.make, v2.door, v2.engine.electric)

}
*/

//====================================================================================

/*
type person struct {
	f_name      string
	l_name      string
	ice_flavour []string
}

func main() {
	p1 := person{
		f_name:      "Ajay",
		l_name:      "kumar",
		ice_flavour: []string{"black_current", "oreo"},
	}

	fmt.Println(p1.f_name)
	fmt.Println(p1.l_name)

	for _, v := range p1.ice_flavour {
		fmt.Println(v)
	}

	m := map[string]person{
		p1.l_name: p1,
	}
	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.ice_flavour {
			fmt.Println(v.f_name, v.l_name, v2)
		}
	}

}
*/

//====================================================================================

func main() {
	as := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Albert",
		friends: map[string]int{
			"David": 34,
			"Bruse": 41,
		},
		favDrinks: []string{"coke", "redbull", "sprite"},
	}

	fmt.Println(as)
	fmt.Println(as.first)
	fmt.Println(as.friends)
	fmt.Println(as.favDrinks)

}
