package main

import "fmt"

func main() {
	//Maps are like key value pairs.

	var students = map[string]int{"akshay": 10, "rohit": 30}
	fmt.Println(students)

	testMap := make(map[string]string)
	testMap["hello"] = "world"
	testMap["GO"] = "lang"
	fmt.Println(testMap)

	for i, v := range testMap {
		fmt.Printf("%v %v\n", i, v)

	}

}
