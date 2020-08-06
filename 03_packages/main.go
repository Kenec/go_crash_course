package main

import (
	"fmt"
	"math"
	"github.com/Kenec/go_crash_course/03_packages/strutli"
)

func main() {
	fmt.Println(math.Floor(2.8))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutli.Reverse("dlroW olleH"))
}
