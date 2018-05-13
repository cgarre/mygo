package main

import "fmt"

func main() {

	fmt.Println("hello whats your name ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("hello ", input)
}
