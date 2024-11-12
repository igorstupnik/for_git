package main

import (
	"fmt"
	"os"
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

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Ошибка: деление на ноль!")
		return 0
	}
	return a / b
}

func getInput() (float64, float64, string) {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка: введено не число.")
		os.Exit(0)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка: введено не число.")
		os.Exit(0)
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Ошибка: неизвестный оператор ", operator)
		os.Exit(0)
	}

	return a, b, operator
}

func main() {
	a, b, operator := getInput()

	var result float64

	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		fmt.Println("Неизвестный оператор")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
