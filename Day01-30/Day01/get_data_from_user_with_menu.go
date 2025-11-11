package main

import (
	"fmt"
)

func menu_options() int {
	fmt.Println("Menu")
	fmt.Println("Option 1: Add new program;")
	fmt.Println("Option 2: Select programs;")
	fmt.Println("Option 3: Edit a program;")
	fmt.Println("Option 4: Delete a program")

	var selection int
	fmt.Println("Enter your selection: ")
	fmt.Scan(&selection)
	fmt.Printf("You selected %d \n\n", selection)

	return selection
}

func menu_select(selection int) {
	switch selection {
	case 1:
		fmt.Println("Let's add a new program")
	case 2:
		fmt.Println("Let's select a program")
	case 3:
		fmt.Println("Let's edit a program")
	case 4:
		fmt.Println("Let's delete a program")
	default:
		fmt.Println("Invalid selection")
	}
}

func main() {
	menu_select(menu_options())
}
