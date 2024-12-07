package calculations

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type Calculator struct {
	logger *logrus.Logger
}

func NewCalcualator(logger *logrus.Logger) *Calculator {
	return &Calculator{logger: logger}
}

func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.logger.WithFields(logrus.Fields{"a": a, "b": b, "result": result}).Info("Performed addition")
	return result
}

func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.logger.WithFields(logrus.Fields{"a": a, "b": b, "result": result}).Info("Performed substraction")
	return result
}

func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.logger.WithFields(logrus.Fields{"a": a, "b": b, "result": result}).Info("Performed multiplication")
	return result
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		err := errors.New("division by zero")
		c.logger.WithFields(logrus.Fields{"a": a, "b": b}).WithError(err).Error("Error in division")
		return 0, err
	}
	result := a / b
	c.logger.WithFields(logrus.Fields{"a": a, "b": b, "result": result}).Info("Performed division")
	return result, nil
}

func (c *Calculator) Calculate(a, b float64, operator string) (float64, error) {
	c.logger.WithFields(logrus.Fields{
		"a":        a,
		"b":        b,
		"operator": operator,
	}).Info("Starting calculation")
	switch operator {
	case "+":
		return c.Add(a, b), nil
	case "-":
		return c.Subtract(a, b), nil
	case "*":
		return c.Multiply(a, b), nil
	case "/":
		return c.Divide(a, b)
	default:
		err := errors.New("invalid operator")
		c.logger.WithFields(logrus.Fields{
			"a":        a,
			"b":        b,
			"operator": operator,
		}).WithError(err).Error("Error in calculation")
		return 0, err
	}
}
