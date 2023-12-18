package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Person WHERE DNI = dni
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Employee WHERE Id = id
	return Employee{}, nil
}

func GetFulTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	employee, er := GetEmployeeById(id)
	if er != nil {
		return ftEmployee, er
	}

	ftEmployee.Employee = employee

	person, er := GetPersonByDNI(dni)
	if er != nil {
		return ftEmployee, er
	}

	ftEmployee.Person = person

	return ftEmployee, nil
}
