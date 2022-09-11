package main

type Calculate interface {
	annualSalary() int
	getName() string
	getAge() int
	getProfession() Employee
}
type Employee struct {
	Name   string
	Age    int
	Salary int
	Bonus  int
}
type Technician struct {
	details Employee
}
type Engineer struct {
	details Employee
}

// Technician
func (t Technician) annualSalary() int {
	Total := (t.details.Salary * 12) + t.details.Bonus
	return Total
}
func (t Technician) getName() string {
	return t.details.Name
}
func (t Technician) getAge() int {
	return t.details.Age
}
func (t Technician) getProfession() Employee {
	return t.details
}

// Engineer
func (e Engineer) annualSalary() int {
	Total := (e.details.Salary + e.details.Bonus) * 12
	return Total
}
func (e Engineer) getName() string {
	return e.details.Name
}
func (e Engineer) getAge() int {
	return e.details.Age
}
func (e Engineer) getProfession() Employee {
	return e.details
}
