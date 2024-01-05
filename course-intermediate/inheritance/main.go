package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func (ftEmployee FullTimeEmployee) GetMessage() string {
	return "Full time employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) GetMessage() string {
	return "Temporary employee"
}

type PrintInfo interface {
	GetMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e Employee) GetId() int {
	return e.id
}

func GetMessage(p Person) {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 30
	ftEmployee.id = 1
	fmt.Println(ftEmployee)

	temporaryEmployee := TemporaryEmployee{}
	getMessage(temporaryEmployee)
	getMessage(ftEmployee)
}
