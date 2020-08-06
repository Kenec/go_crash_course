package main

import "fmt"

// Annoymous function implementing closure
func adder() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}

func main()  {
	sum := adder()

	fmt.Println("Starting the sum")
	fmt.Println(sum(1))
	fmt.Println(sum(1))
	fmt.Println(sum(1))

	by_two := adder()

	fmt.Println("Starting the by two count")
	fmt.Println(by_two(2))
	fmt.Println(by_two(2))
}
