package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee person

func main() {
	p := new(person)
	p.name = "john"
	p.age = 29
	e := new(employee)
	e = (*employee)(p) // error: cannot convert p (type *person) to type employee
	employeeOnly(e)
	fmt.Println(e.age)
}

func employeeOnly(e *employee) {
	fmt.Printf("%s is a person\n", e.name)
}
