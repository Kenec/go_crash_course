package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	msg := "Kene Nnamani"

	ping(pings, msg)
	pong(pings, pongs)

	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(ping <-chan string, pongs chan<- string) {
	msg := <-ping
	pongs <- msg
}
