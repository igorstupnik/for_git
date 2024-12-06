package calculations

import (
	"errors"

	"github.com/sirupsen/logrus"
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

func divide(a, b float64, log *logrus.Logger) (float64, error) {
	if b == 0 {
		err := errors.New("division by zero")
		log.WithError(err).Error("Error: division by zero")
		return 0, err
	}
	return a / b, nil
}

func Calculate(a, b float64, operator string, log *logrus.Logger) (float64, error) {
	log.WithFields(logrus.Fields{
		"a":        a,
		"b":        b,
		"operator": operator,
	}).Info("Calculation in prosess")
	switch operator {
	case "+":
		return add(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b, log)
	default:
		err := errors.New("invalid operator")
		log.WithError(err).Error("Error: invalid operator")
		return 0, err
	}
}
