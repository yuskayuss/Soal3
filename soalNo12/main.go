package main

import (
	"fmt"
)


type Employee interface {
	SetSalary(int)
	GetSalary() int
	SetGrade(string)
	GetGrade() string
	Label()
}


type BaseEmployee struct {
	salary int
	grade  string
}

func (e *BaseEmployee) SetSalary(s int) {
	e.salary = s
}
func (e *BaseEmployee) GetSalary() int {
	return e.salary
}
func (e *BaseEmployee) SetGrade(g string) {
	e.grade = g
}
func (e *BaseEmployee) GetGrade() string {
	return e.grade
}
func (e *BaseEmployee) Label() {
	fmt.Println("Employee's data:\n")
}


type Manager struct {
	BaseEmployee
}


type Supervisor struct {
	BaseEmployee
}


func printEmployeeData(e Employee) {
	e.Label()
	fmt.Println("Grade:", e.GetGrade())
	fmt.Println("Salary:", e.GetSalary())
}

func main() {
	manager := &Manager{}
	manager.SetSalary(8000)
	manager.SetGrade("A")

	supervisor := &Supervisor{}
	supervisor.SetSalary(6000)
	supervisor.SetGrade("B")

	
	printEmployeeData(manager)
	fmt.Println()
	printEmployeeData(supervisor)
}
