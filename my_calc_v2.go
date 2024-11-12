package main

import (
	"errors"
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func getInput() (float64, float64, string, error) {
	var a, b float64
	var operator string

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		return 0, 0, "", errors.New("input is not a number")
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		return 0, 0, "", errors.New("input is not a number")
	}

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return 0, 0, "", fmt.Errorf("unknown operator %s", operator)
	}

	return a, b, operator, nil
}

func main() {
	a, b, operator, err := getInput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var result float64

	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result, err = divide(a, b)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	default:
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
