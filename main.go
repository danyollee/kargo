package main

import "fmt"

func fill_map(m map[int]string) {

	m[0] = "Zero"
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"
	m[4] = "Four"
	m[5] = "Five"
	m[6] = "Six"
	m[7] = "Seven"
	m[8] = "Eight"
	m[9] = "Nine"

}

func main() {
	fmt.Println("Hello wrld")

	// create a test array
	// array := [10]int{0,1,2,3,4,5,6,7,8,9}
	array := [5]int{1,3,5,7,9}
	fmt.Println(array)

	// loop thru array

	for i, val := range array {
		fmt.Println(i, val)
	}

	// Maps are similar to a python dictionary
	// specify key/val pair 
	m:= make(map[int]string)

	fill_map(m)

	fmt.Println(m)
	fmt.Println(m[5])
	fmt.Println(m[0])
	fmt.Println(m[3])
}
