package main

import (
	"fmt"
)

func sumNumbers(x, y int) int {
	return x + y
}

func greetPerson(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go functions!", name)
}

func main() {
	result := sumNumbers(5, 3)
	fmt.Printf("Sum of 5 and 3 is: %d\n", result)

	greeting := greetPerson("Alice")
	fmt.Println(greeting)
}
