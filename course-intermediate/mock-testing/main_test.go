package main

import (
	"reflect"
	"testing"
)

func TestGetFulTimeEmployeeById(t *testing.T) {

	tableTest := []struct {
		description      string
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			description: "Successful retrieval of full-time employee",
			id:          1,
			dni:         "123456789",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "123456789",
						Name: "John Doe",
						Age:  30,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
				Person: Person{
					DNI:  "123456789",
					Name: "John Doe",
					Age:  30,
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range tableTest {
		t.Run(test.description, func(t *testing.T) {
			test.mockFunc()
			ft, err := GetFulTimeEmployeeById(test.id, test.dni)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(ft, test.expectedEmployee) {
				t.Errorf("Unexpected result. Got: %v, want: %v", ft, test.expectedEmployee)
			}
		})
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
