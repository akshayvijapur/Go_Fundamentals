package main

import "fmt"

type employeeDetails struct {
	name    string
	Id, age int
}

func main() {
	e1 := employeeDetails{} //Create the object of the structure but dont initialize it.

	e1.setEmployeeDetails(10, "Akshay", 25) // with the created object call the functions.
	e1.showEmployeeDetails()
}

// setEmployeeDetails : It fills (*employeeDetails) with the values from the parameters.
// *employeeDetails needed to store the context of the data outside the functions.
// setEmployeeDetails can access the employee object inside the function.
func (employee *employeeDetails) setEmployeeDetails(ID int, empName string, age int) {
	employee.name = empName
	employee.Id = ID
	employee.age = age
}

// showEmployeeDetails displays the employee information.
func (employee employeeDetails) showEmployeeDetails() {
	fmt.Println("Employee ID ", employee.Id)
	fmt.Println("Employee Name ", employee.name)
	fmt.Println("Employee age ", employee.age)
}
