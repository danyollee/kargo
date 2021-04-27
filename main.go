package main

import (
	"fmt"
	"os"
)

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


	m := make(map[string]string)
	fill_map(m)


	first := true
	for _, v := range(os.Args[1:]) {
		if first {
			first = false
		} else {
			fmt.Print(",")			
		}
		for _, c := range(v) {
			fmt.Print(m[string(c)])
		}
	}
	fmt.Print("\n")

}
