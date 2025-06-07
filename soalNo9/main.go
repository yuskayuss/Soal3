package main

import "fmt"


type Employee interface {
	SetSalary(salary int)
	GetSalary() int
	SetGrade(grade string)
	GetGrade() string
	Label()
}


type baseEmployee struct{}

func (b baseEmployee) Label() {
	fmt.Println("Employee's data:")
}


type Engineer struct {
	baseEmployee       
	salary     int
	grade      string
}

func (e *Engineer) SetSalary(salary int) {
	e.salary = salary
}

func (e *Engineer) GetSalary() int {
	return e.salary
}

func (e *Engineer) SetGrade(grade string) {
	e.grade = grade
}

func (e *Engineer) GetGrade() string {
	return e.grade
}


type Manager struct {
	baseEmployee       
	salary     int
	grade      string
}

func (m *Manager) SetSalary(salary int) {
	m.salary = salary
}

func (m *Manager) GetSalary() int {
	return m.salary
}

func (m *Manager) SetGrade(grade string) {
	m.grade = grade
}

func (m *Manager) GetGrade() string {
	return m.grade
}

func main() {
	var emp Employee

	
	emp = &Engineer{}
	emp.SetSalary(7000)
	emp.SetGrade("Senior")
	emp.Label()
	fmt.Println("Salary:", emp.GetSalary())
	fmt.Println("Grade:", emp.GetGrade())

	fmt.Println()


	emp = &Manager{}
	emp.SetSalary(10000)
	emp.SetGrade("Lead")
	emp.Label()
	fmt.Println("Salary:", emp.GetSalary())
	fmt.Println("Grade:", emp.GetGrade())
}
