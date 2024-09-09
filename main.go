package main

import "fmt"

func Add() int {
	fmt.Println("Enter two numbers to add: ")
	var x, y int
	n, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if n != 2 {
		fmt.Println("Invalid input")
		return 0
	}
	return x + y
}

func Subtract() int {
	fmt.Println("Enter two numbers to subtract: ")
	var x, y int
	n, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if n != 2 {
		fmt.Println("Invalid input")
		return 0
	}
	return x - y
}

func Multiply() int {
	fmt.Println("Enter two numbers to multiply:")
	var x, y int
	n, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if n != 2 {
		fmt.Println("Wrong input")
		return 0
	}
	return x * y
}

func Divide() int {
	fmt.Println("Enter two numbers to divide: ")
	var x, y int
	n, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if n != 2 {
		fmt.Println("Wrong input")
		return 0
	}
	return x / y
}

func Choice(choiceValue int) {
	switch choiceValue {
	case 1:
		fmt.Println(Add())

	case 2:
		fmt.Println(Subtract())

	case 3:
		fmt.Println(Multiply())

	case 4:
		fmt.Println(Divide())

	default:
		fmt.Println("Invalid choice")
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
		fmt.Println(err)
		return
	}
	if n > 4 || n < 0 {
		fmt.Println("Invalid input")
		return
	}
	Choice(choice)
}
