package presentation

import (
	"errors"
	"fmt"
)

func GetInput() (float64, float64, string, error) {
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

func DispayResults(result float64) {
	fmt.Printf("Result: %.2f\n", result)
}

func DisplayError(err error) {
	fmt.Println("Error:", err)
}
