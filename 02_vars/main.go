package main

import "fmt"


func main() {
	// Using var
	var name string = "Kene"
	var age int32	= 30
	const isCool = true

	// Shorthand [ This type of declaration cannot happen outside of a function ]
	full_name := "Kenechukwu Nnamani"
	sex, party := "Male", "PDP"

	fmt.Printf("Fullname: %s\n", full_name)
	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", age)
	fmt.Printf("Gender: %s\nPolitical party: %s\n", sex, party)
}