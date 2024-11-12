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
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil //возвращаем результат деления и nil для ошибки
}

func getInput() (float64, float64, string) {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка: введено не число.")
		return 0, 0, ""
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка: введено не число.")
		return 0, 0, ""
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Ошибка: неизвестный оператор ", operator)
		return 0, 0, ""
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
		result, _ = divide(a, b)
		_, err := divide(a, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
	default:
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
