package main

import "fmt"

//Interface are collection of method signature.
// In an organization, we can have two kinds of employees.
// One is manager and other one is reporter.
// If we  want to raise the salary of all employees without calling explicitly manager or worker.
// we can create an interface which has method called as raiseSalary.
// then we can implement it according to the calling object.

type Employees interface {
	raiseSalary()
	showEmployeeDetails()
}

type Manager struct {
	employeeID     int
	employeeName   string
	employeeSalary int
}

type Reportee struct {
	employeeID     int
	employeeName   string
	employeeSalary int
}

func main() {
	m1 := Manager{}
	m1.setEmployeeDetails(1, "akshay", 1000)
	e1 := Reportee{}
	e1.setEmployeeDetails(2, "ram", 100)

	EmployeeManagementSystem(&m1)
	EmployeeManagementSystem(&e1)

	// or you can also create like this

	var emp Employees

	emp = &m1
	emp.raiseSalary()
	emp.showEmployeeDetails()

	emp = &e1
	emp.raiseSalary()
	emp.showEmployeeDetails()

}

func EmployeeManagementSystem(e Employees) {
	e.showEmployeeDetails() // Show the details of the employee before increment.
	e.raiseSalary()         // Perform increment
	e.showEmployeeDetails() // Show the details of the employee after increment.
}

func (mgr *Manager) setEmployeeDetails(id int, name string, salary int) {
	mgr.employeeID = id
	mgr.employeeName = name
	mgr.employeeSalary = salary
}

func (rep *Reportee) setEmployeeDetails(id int, name string, salary int) {
	rep.employeeID = id
	rep.employeeName = name
	rep.employeeSalary = salary
}

func (mgr *Manager) showEmployeeDetails() {
	fmt.Println(mgr.employeeID, mgr.employeeName, mgr.employeeSalary)
}

func (rep *Reportee) showEmployeeDetails() {
	fmt.Println(rep.employeeID, rep.employeeName, rep.employeeSalary)
}

func (mgr *Manager) raiseSalary() {
	mgr.employeeSalary = mgr.employeeSalary + 1000
}

func (rep *Reportee) raiseSalary() {
	rep.employeeSalary = rep.employeeSalary + 10
}
