package main

import (
	"fmt"
	"strconv"
)

// Define struct
type Person struct {
	firstName	string
	lastName	string
	city		string
	gender		string
	age			int
}

// Greeting method (value receiver)
func (p Person) greet() string  {
	return "Hello, I am " + p.firstName + " " + p.lastName + ". I am " + strconv.Itoa(p.age) + " and I live in " + p.city
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday()  {
	p.age++
}

func main()  {
	// Init Person using struct
	person1 := Person{firstName:"Kene", lastName:"Nnamani", city:"Lagos", gender:"Male", age:30}

	person1.hasBirthday()
	fmt.Println(person1.greet())
}
