package console

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Console struct {
	logger *logrus.Logger
}

func New(logger *logrus.Logger) *Console {
	return &Console{logger: logger}
}

func (c *Console) GetInput() (float64, float64, string, error) {
	var a, b float64
	var operator string

	c.logger.Info("Requesting the first number")
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		c.logger.WithError(err).Error("Invalid input for the first number")
		return 0, 0, "", errors.New("invalid input for the first number")
	}
	c.logger.WithField("a", a).Info("First number input succesful")

	c.logger.Info(("Requesting the second number"))
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		c.logger.WithError(err).Error("Invalid input for the second number")
		return 0, 0, "", errors.New("invalid for the second number")
	}
	c.logger.WithField("b", b).Info("Second number input successful")

	c.logger.Info("Requesting the operator")
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	if isNotOperator(operator) {
		c.logger.WithField("operator", operator).Error("Invalid operator input")
		return 0, 0, "", fmt.Errorf("invalid operator %s", operator)
	}
	c.logger.WithField("operator", operator).Info("Operator input successful")

	return a, b, operator, nil
}

func (c *Console) DispayResults(result float64) {
	fmt.Printf("Result: %.2f\n", result)
	c.logger.WithField("result", result).Info("Displaying results")
}

func (c *Console) DisplayError(err error) {
	fmt.Println("Error:", err)
	c.logger.WithError(err).Debug("Displaying error")
}

func isNotOperator(s string) bool {
	return s != "+" && s != "-" && s != "*" && s != "/"
}
