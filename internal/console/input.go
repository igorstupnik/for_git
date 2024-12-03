package console

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func GetInput(log *logrus.Logger) (float64, float64, string, error) {
	var a, b float64
	var operator string

	log.Info("Request on the first number")

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.WithError(err).Error("Wrong first number format")
		return 0, 0, "", errors.New("input is not a number")
	}

	log.Info(("Request on the second number"))
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		log.WithError(err).Error("Wrong second number format")
		return 0, 0, "", errors.New("input is not a number")
	}

	log.Info("Request on the operator")
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	if isNotOperator(operator) {
		log.WithField("Error with", operator).Info("Wrong format of operator")
		return 0, 0, "", fmt.Errorf("unknown operator %s", operator)
	}

	return a, b, operator, nil
}

func DispayResults(result float64, log *logrus.Logger) {
	log.WithField("result", result).Info("The result has displayed successfully")
	fmt.Printf("Result: %.2f\n", result)
}

func DisplayError(err error, log *logrus.Logger) {
	log.WithField("Error", err).Info("The error has displayed successfully")
	fmt.Println("Error:", err)
}

func isNotOperator(s string) bool {
	return s != "+" && s != "-" && s != "*" && s != "/"
}
