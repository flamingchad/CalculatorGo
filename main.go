package main

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
	if y == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return x / y
}

func Choice(choiceValue, x, y int) int {
	switch choiceValue {
	case 1:
		return Add(x, y)
	case 2:
		return Subtract(x, y)
	case 3:
		return Multiply(x, y)
	case 4:
		return Divide(x, y)
	default:
		fmt.Println("Invalid choice")
		return -1
	}
}

func main() {
	fmt.Println("------------------>    Calculator    <------------------")
	fmt.Println("Options: ")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	var choice int

	n, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error reading choice:", err)
		return
	}
	if n != 1 || choice < 1 || choice > 4 {
		fmt.Println("Invalid choice")
		return
	}

	var input1, input2 int

	fmt.Println("Enter the first number: ")
	_, err = fmt.Scan(&input1) //using _ here cause we want to ignore the count of the scanned items, instead only error.
	if err != nil {
		fmt.Println("Error reading first number:", err)
		return
	}

	fmt.Println("Enter the second number: ")
	_, err = fmt.Scan(&input2)
	if err != nil {
		fmt.Println("Error reading second number:", err)
		return
	}

	result := Choice(choice, input1, input2)
	fmt.Println("Result:", result)
}
