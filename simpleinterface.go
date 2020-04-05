package main

import (
	"fmt"
)

type i interface {
	sayHello()
}

func main() {
	var a i // a is interface
	intval := intnum(10)
	stringval := stringnum("akshay")

	a = intval // a intnum implements sayHello
	a.sayHello()
	a = stringval // a stringnum implements Abser
	a.sayHello()
}

type intnum int

func (i intnum) sayHello() {
	fmt.Println("int value is ", i)
}

type stringnum string

func (f stringnum) sayHello() {
	fmt.Println("string  value is ", f)
}
