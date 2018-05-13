package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter the first number :")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number :")
	fmt.Scanln(&b)
	c := a + b
	fmt.Println(a, " + ", b, " = ", c)
}
