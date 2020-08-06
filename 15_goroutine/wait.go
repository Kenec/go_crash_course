package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func responseSizeAgain(url string) {
	defer wg.Done()

	fmt.Println("Step 1: ", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 2: ", url)
	defer response.Body.Close()

	fmt.Println("Step 3: ", url)
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 4: ", len(body))
}

func main() {
	wg.Add(3)
	fmt.Println("Start Goroutine")

	go responseSizeAgain("https://www.golangprograms.com")
	go responseSizeAgain("https://stackoverflow.com")
	go responseSizeAgain("https://coderwall.com")

	wg.Wait()
	fmt.Println("Terminating program")
}
