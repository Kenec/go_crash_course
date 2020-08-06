package main

import "fmt"

func main()  {
	// define map
	emails := make(map[string]string)

	// assign kv
	emails["kene"] = "kene@gmail.com"
	emails["amaka"] = "amaka@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "amaka")
	fmt.Println(emails)

	// declare map and add kv
	nationals := map[string]string{"Kene": "Nigerian", "Judith": "German"}
	fmt.Println(nationals)



}
