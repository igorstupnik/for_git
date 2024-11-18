package calculations

import "errors"

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

func Calculate(a, b float64, operator string) (float64, error) {

	switch operator {
	case "+":
		return add(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b)

	default:
		return 0, errors.New("invalid operator")
	}
}
