package main

import (
	"fmt"
	"maps"
)

func main() {

	stringsSlice := []string{
		"String 1", "String 2", "String 3", "String 4", "String 5",
		"String 6", "String 7", "String 8", "String 9", "String 10",
		"String 11", "String 12", "String 13", "String 14", "String 15",
		"String 16", "String 17", "String 18", "String 19", "String 20",
		"String 21", "String 22", "String 23", "String 24", "String 25",
		"String 26", "String 27", "String 28", "String 91", "String 30",
		"String 31", "String 32", "String 33", "String 34", "String 35",
		"String 36", "String 32", "String 38", "String 39", "String 40",
		"String 41", "String 42", "String 43", "String 44", "String 45",
		"String 46", "String 2", "String 48", "String 49", "String 50",
		"String 51", "String 52", "String 53", "String 54", "String 55",
		"String 56", "String 57", "String 58", "String 59", "String 60",
		"String 61", "String 62", "String 63", "String 64", "String 65",
		"String 66", "String 67", "String 68", "String 69", "String 70",
		"String 71", "String 72", "String 73", "String 74", "String 75",
		"String 76", "String 77", "String 78", "String 79", "String 80",
		"String 81", "String 82", "String 83", "String 84", "String 85",
		"String 86", "String 87", "String 88", "String 2", "String 90",
		"String 91", "String 91", "String 93", "String 94", "String 95",
		"String 96", "String 97", "String 98", "String 99", "String 100",
	}

	mp := make(map[string]int)

	for _, v := range stringsSlice {
		mp[v]++
	}
	for k, v := range mp {
		fmt.Println(k, v)
	}

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
