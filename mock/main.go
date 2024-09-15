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
	Employee
	Person
}

// GetPersonByDNI simulación de consulta a DB
var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(time.Second * 5)
	// SELECT * FROM people WHERE ...
	return Person{
		DNI:  "1234",
		Name: "santi",
		Age:  24,
	}, nil
}

// GetEmployeeById simulación de consulta a DB
var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(time.Second * 5)
	// SELECT * FROM employees WHERE ...
	return Employee{
		Id:       23,
		Position: "Developer",
	}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = p
	ftEmployee.Employee = e

	return ftEmployee, nil
}
