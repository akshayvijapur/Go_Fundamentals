package main

import "fmt"

func main() {

	// Lets create a slice

	var testSlice = []int{10, 20, 30, 40, 50}

	for i, v := range testSlice {
		fmt.Printf("%v %v \n", i, v)
	}

}
