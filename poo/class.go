package main

import (
	"fmt"
)

// El struct sería como la clase:

type Employee struct {
	id       int
	vacation bool
}

type Person struct {
	name     string
	lastname string
	age      int
}

// Herencia (en go se utiliza la composición):

type FullTimeEmployee struct {
	Employee
	Person
	endDate string
}

type TemporaryEmployee struct {
	Employee
	Person
	taxRate int
}

func (e *Employee) SetId(id int) {
	e.id = id
}

// Methods with Pointer receivers

func (e *Person) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Person) GetName() string {
	return e.name
}

func (fte FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (te TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

// Simulación de constructor (mas recomendada):

func NewEmployee(id, age int, name, lastname string, vacation bool, endDate string) *FullTimeEmployee {
	return &FullTimeEmployee{
		Employee: Employee{
			id:       id,
			vacation: vacation,
		},
		Person: Person{
			name:     name,
			lastname: lastname,
			age:      age,
		},
		endDate: endDate,
	}
}

func NewTemporaryEmployee(
	id,
	taxRate,
	age int,
	name,
	lastname string,
	vocation bool,
) *TemporaryEmployee {
	return &TemporaryEmployee{
		Employee: Employee{
			id:       id,
			vacation: vocation,
		},
		Person: Person{
			name:     name,
			lastname: lastname,
			age:      age,
		},
		taxRate: taxRate,
	}
}

// En go no se puede aplicar polimorfismo como si se puede aplicar en otros programing languages

func printMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

// interfaces

type PrintInfo interface {
	getMessage() string
}

func main() {
	pepe := new(FullTimeEmployee)
	fmt.Println("pepe:", pepe)
	pepe.SetId(111)
	fmt.Printf("pepe ID: %d\n", pepe.GetId())
	santi := FullTimeEmployee{
		Employee: Employee{
			id:       1,
			vacation: true,
		},
		Person: Person{
			name:     "Santiago",
			lastname: "Restrepo",
			age:      24,
		},
		endDate: "2024",
	}
	fmt.Printf("santi: %+v\n", santi)

	santi.SetId(2)
	santi.SetName("SantiSite")
	fmt.Printf("santi: %+v\n", santi)
	fmt.Println("nombre de santi:", santi.GetName())
	fmt.Println("id de santi:", santi.GetId())

	alejo := FullTimeEmployee{}
	alejo.name = "Alejandro"
	alejo.SetId(3)
	fmt.Println("alejo:", alejo)

	// otra manera de simular el constructor (new):
	juan := new(FullTimeEmployee)
	juan.name = "Juan David"
	juan.SetId(4)
	fmt.Println("juan:", *juan)

	cristian := NewEmployee(5, 32, "Cristian", "Castro", true, "2024")
	fmt.Println("cristian:", *cristian)

	// Herencia (en go se utiliza la compisicion):
	fullTimeEmployee := FullTimeEmployee{}
	fullTimeEmployee.id = 6
	fullTimeEmployee.name = "Sebastian"
	fullTimeEmployee.lastname = "Quintero"
	fullTimeEmployee.age = 30
	fullTimeEmployee.vacation = false
	fmt.Println("fullTimeEmployee:", fullTimeEmployee)

	// En Go no se puede aplicar polimorfismo.
	// Para alcanzar este mismo objetivo, Go aplica
	// un principio llamado "Composition Over Inheritance"
	// que utilizando un compo anónimo en un struct puede
	// "heredar" todas las propiedades y métodos.
	//GetMessage(fullTimeEmployee)

	// aplicando interfaces:
	temporaryEmployee := TemporaryEmployee{}
	printMessage(temporaryEmployee)
	printMessage(fullTimeEmployee)
}
