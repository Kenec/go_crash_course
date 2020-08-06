package main

import "fmt"

func main(){
	ids := []int{1,4,2,55,3,55}

	// Loop through the ids
	for i, id := range ids {
		fmt.Printf("%d - ID:  %d\n", i, id)
	}

	// Add all the ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)


	// Range with map
	emails := map[string]string{"Kene": "kene@gmail.com", "Amaka": "amaka@gmail.com"}
	for k, v := range emails {
		 fmt.Printf("%s: %s\n", k, v)
	}
}
