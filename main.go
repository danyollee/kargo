package main

import (
	"fmt"
	"strconv"
)

// fills map to convert int to string output
func fill_map(m map[string]string) {

	m["0"] = "Zero"
	m["1"] = "One"
	m["2"] = "Two"
	m["3"] = "Three"
	m["4"] = "Four"
	m["5"] = "Five"
	m["6"] = "Six"
	m["7"] = "Seven"
	m["8"] = "Eight"
	m["9"] = "Nine"

}

func main() {
	fmt.Println("Hello wrld")

	// create a test array
	array := [6]int{1,3,5,7,9, 35}

	// create and fill map of int->string
	m:= make(map[string]string)

	fill_map(m)

	fmt.Println(array)
	for _, v := range(array) {
		s := strconv.Itoa(v)
		for _, c := range(s) {
			fmt.Println(m[string(c)])
		}
	}

}
