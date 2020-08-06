package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	//messages <- "Hello World" // This will fail unless the channel is buffered.

	msg := <-messages
	fmt.Println(msg)
}
