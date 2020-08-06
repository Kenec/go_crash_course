package main

import "fmt"

func main() {
	// Arrays [Arrays have fixed number of items]
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	foodArr := [2]string{"Rice", "Beans"}

	fmt.Println(fruitArr)
	fmt.Println(foodArr)

	// Slices [Slices do not have a fixed number of items]
	fruitSlices := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlices)
}
