package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read the value
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with pointer
	*b = 10
	fmt.Println(a)
}
