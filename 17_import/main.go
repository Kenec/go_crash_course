package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Starting main")
	mainWithExitCode()
}

func mainWithExitCode() int {
	parseFlags()
	if stop := initLogrus(); stop != nil {
		defer fmt.Println(stop())
	}
	fmt.Println("Man with exit code")
	return 1
}

func initLogrus() func () string {
	return func () string  {
		return "Hello World in the initLogrus"
	}
}