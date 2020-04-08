//stringer  interface
// override string method

package main

import "fmt"

type employee struct {
	name string
	age  int
}

func main() {

	akshay := employee{
		name: "akshay",
		age:  285,
	}

	ram := employee{
		name: "ram",
		age:  25,
	}

	fmt.Println(akshay)
	fmt.Println(ram)
}

//This method will opverride string method.
//String() string
func (e employee) String() string {
	return fmt.Sprintf("name: %v, age: %v ", e.name, e.age)
}
