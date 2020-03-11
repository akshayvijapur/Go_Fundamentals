package main

import (
	"fmt"
	"time"
)

// Go routine can be any function which executes parallely with main function.
// Routine will be dead once the main program terminates as both main and routine runs in same address space.
func main() {
	fmt.Println("Main Program: Hello ")
	go sayHello()
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Main Program: Hello ")
}

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("GO Routine: World")
	}
}
