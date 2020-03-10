package main

import "fmt"

// Sample Structure.
type employee struct {
	name  string
	empID int
	panID int
}

func main() {

	// Initilize the structure.
	akshayVijapur := employee{
		name:  "akshay",
		empID: 10,
		panID: 30,
	}

	// Sample Array.
	var testArray [10]int

	for i := 0; i < 10; i++ {
		testArray[i] = i + 10
	}

	for i := 0; i < len(testArray); i++ {
		fmt.Println(testArray[i])
	}

	//Sample Slice.
	testSlice := testArray[1:5]

	for i := 0; i < len(testSlice); i++ {
		fmt.Println(testSlice[i])
		testSlice[i] = i + 50
	}

	for i := 0; i < len(testArray); i++ {
		fmt.Println(testArray[i])
	}

	testPointer := &akshayVijapur

	fmt.Println(testPointer.empID)

	testArraySlice := []string{"hello", "World", "Akshay"}

	for i := 0; i < len(testArraySlice); i++ {
		fmt.Println(testArraySlice[i])
	}

	testArrayStruct := []struct {
		EmpID     int
		EmpName   string
		EmpSalary int
	}{
		{10, "AKshay", 1000},
		{15, "Ram", 1000},
	}

	for i := 0; i < len(testArrayStruct); i++ {
		fmt.Println(testArrayStruct[i])
	}

}
