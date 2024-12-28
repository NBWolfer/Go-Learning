package packages

import "fmt"

func Calculator() {
	// Terminal UI
	fmt.Println("Calculator")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Modulo")
	fmt.Println("6. Check Palindrome")
	fmt.Println("7. Factorial")
	fmt.Println("8. Fibonacci")
	fmt.Println("9. GCD")

	fmt.Println("0. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	var x, y int

	// Choose within switch case
	// Infinite loop until user chooses to exit
	for choice != 0 {
		fmt.Print("\033[H\033[2J")

		switch choice {
		case 0:
			fmt.Println("Exiting...")
			choice = 0
		case 1:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&x)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&y)
			fmt.Println("Result: ", add(x, y))
		case 2:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&x)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&y)
			fmt.Println("Result: ", subtract(x, y))
		case 3:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&x)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&y)
			fmt.Println("Result: ", multiply(x, y))
		case 4:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&x)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&y)
			fmt.Println("Result: ", divide(x, y))
		case 5:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&x)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&y)
			fmt.Println("Result: ", modulo(x, y))
		case 6:
			fmt.Print("Enter a string: ")
			var str string
			fmt.Scanln(&str)
			if checkPalindromeRecursive(str) {
				fmt.Println("Palindrome")
			} else {
				fmt.Println("Not Palindrome")
			}
		case 7:
			fmt.Print("Enter a number: ")
			var n int
			fmt.Scanln(&n)
			fmt.Println("Factorial: ", factorialRecursive(n))
		case 8:
			fmt.Print("Enter a number: ")

			var n int
			fmt.Scanln(&n)
			fmt.Print("\nFibonacci: ")
			fibonacciRecursiveSeries(n)
			fmt.Print("\n")
		case 9:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&x)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&y)
			fmt.Println("GCD: ", gcdRecursive(x, y))
		default:
			fmt.Println("Invalid choice")
		}
		// Clean terminal cls command wait for user input
		fmt.Println("Press Enter to continue...")
		var input string
		fmt.Scanln(&input)
		fmt.Println("Calculator")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Modulo")
		fmt.Println("6. Check Palindrome")
		fmt.Println("7. Factorial")
		fmt.Println("8. Fibonacci")
		fmt.Println("9. GCD")
		fmt.Println("0. Exit")
		fmt.Println("Choose option again: ")
		fmt.Scanln(&choice)
	}
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func modulo(x int, y int) int {
	return x % y
}

// Recursive functions samples
func checkPalindromeRecursive(str string) bool {
	if len(str) <= 1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}
	return checkPalindromeRecursive(str[1 : len(str)-1])
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

// Fibonacci recursive function - prints the fibonacci series up to n
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciRecursiveSeries(n int) {
	for i := 1; i < n; i++ {
		fmt.Printf("%d ", fibonacciRecursive(i))
	}
}

func gcdRecursive(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcdRecursive(y, x%y)
}
