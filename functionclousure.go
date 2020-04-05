package main

import "fmt"

//Function closure refers to variable which are outside its body.

func test() func() int {
	var i = 0 // i is declared outside the function.
	return func() int {
		i = i + 1 //But i is  accessible inside the function.
		return i
	}
}

// main func
func main() {
	testfunc := test()
	fmt.Println(testfunc())
	fmt.Println(testfunc())
	fmt.Println(testfunc())

	testfunc1 := test()
	fmt.Println(testfunc1())
	fmt.Println(testfunc1())
}
