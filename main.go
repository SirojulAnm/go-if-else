package main

import (
	"fmt"
)

func main() {
	var x int
	x = 1 //input the number as you wish

	if x == 1 {
		fmt.Println("one")
	} else if x == 2 {
		fmt.Println("two")
	} else if x == 3 {
		fmt.Println("three")
	} else {
		fmt.Println("more than three")
	}
}
