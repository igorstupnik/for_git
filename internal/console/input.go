package console

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Console struct {
	Logger *logrus.Logger
}

func (c *Console) GetInput() (float64, float64, string, error) {
	var a, b float64
	var operator string

	c.Logger.Info("Requesting the first number")
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		c.Logger.WithError(err).Error("Invalid input for the first number")
		return 0, 0, "", errors.New("invalid input for the first number")
	}
	c.Logger.WithField("a", a).Info("First number input succesful")

	c.Logger.Info(("Requesting the second number"))
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		c.Logger.WithError(err).Error("Invalid input for the second number")
		return 0, 0, "", errors.New("invalid for the second number")
	}
	c.Logger.WithField("b", b).Info("Second number input successful")

	c.Logger.Info("Requesting the operator")
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	if isNotOperator(operator) {
		c.Logger.WithField("operator", operator).Error("Invalid operator input")
		return 0, 0, "", fmt.Errorf("invalid operator %s", operator)
	}
	c.Logger.WithField("operator", operator).Info("Operator input successful")

	return a, b, operator, nil
}

func (c *Console) DispayResults(result float64) {
	fmt.Printf("Result: %.2f\n", result)
	c.Logger.WithField("result", result).Info("Displaying results")
}

func (c *Console) DisplayError(err error) {
	fmt.Println("Error:", err)
	c.Logger.WithError(err).Debug("Displaying error")
}

func isNotOperator(s string) bool {
	return s != "+" && s != "-" && s != "*" && s != "/"
}
