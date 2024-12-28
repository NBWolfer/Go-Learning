package main

// Import files from the same package
import (
	"fmt"
	basics "go-learning/packages/basics"
	binary_tree "go-learning/packages/binary_tree"
	single_linked_list "go-learning/packages/single_linked_list"
)

func main() {
	// UI for the user to choose between Calculator and Data Structures

	fmt.Println("1. Calculator")
	fmt.Println("2. Data Structures")
	fmt.Println("3. Binary Tree")
	fmt.Println("0. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	// Choose within switch case
	// Infinite loop until user chooses to exit
	for choice != 0 {
		switch choice {
		case 0:
			fmt.Println("Exiting...")
			choice = 0
		case 1:
			fmt.Printf("\033[H\033[2J")
			basics.Calculator()
		case 2:
			fmt.Printf("\033[H\033[2J")
			single_linked_list.LinkedListDemo()
		case 3:
			fmt.Printf("\033[H\033[2J")
			binary_tree.BinaryTreeDemo()
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println("To go back to the main menu, press any key")
		var key string
		fmt.Scanln(&key)
		fmt.Printf("\033[H\033[2J")
		fmt.Println("1. Calculator")
		fmt.Println("2. Data Structures")
		fmt.Println("3. Binary Tree")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
	}
}
