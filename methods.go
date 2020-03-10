package main

import "fmt"

type employeeDetails struct {
	name            string
	Id, age, salary int
}

func main() {
	e1 := employeeDetails{} //Create the object of the structure but dont initialize it.

	e1.setEmployeeDetails(10, "Akshay", 25, 1000) // with the created object call the functions.
	e1.showEmployeeDetails()
	incrementSalaryOfTheEmployee(&e1)
	e1.showEmployeeDetails()

	pointer := &e1                //Create a pointer to object.
	pointer.showEmployeeDetails() //We can directly call the function. It is similar to (*pounter).showEmployeeDetails()

}

// This function will receive pointer receivers.
//setEmployeeDetails : It fills (*employeeDetails)  with the values from the parameters.
// *employeeDetails needed to store the context of the data outside the functions.
// setEmployeeDetails can access the employee object inside the function.
func (employee *employeeDetails) setEmployeeDetails(ID int, empName string, age int, salary int) {
	employee.name = empName
	employee.Id = ID
	employee.age = age
	employee.salary = salary
}

// showEmployeeDetails displays the employee information.
func (employee employeeDetails) showEmployeeDetails() {
	fmt.Println("Employee ID ", employee.Id)
	fmt.Println("Employee Name ", employee.name)
	fmt.Println("Employee age ", employee.age)
	fmt.Println("Employee Salary  ", employee.salary)
}

//Need to pass address of the employee object structure to reflect the salary changes.
func incrementSalaryOfTheEmployee(employee *employeeDetails) {
	employee.salary = employee.salary + 100
}
