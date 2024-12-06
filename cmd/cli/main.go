package main

import (
	"for_git/internal/console"
	"for_git/pkg/calculations"
	"for_git/pkg/logger"
)

func main() {

	logger := logger.NewLogger()
	console := &console.Console{Logger: logger}
	calculator := &calculations.Calculator{Logger: logger}

	logger.Info("Starting calculator program")

	a, b, operator, err := console.GetInput()
	if err != nil {
		logger.WithError(err).Error("Input error")
		console.DisplayError(err)
		return
	}

	result, err := calculator.Calculate(a, b, operator)
	if err != nil {
		console.DisplayError(err)
		return
	}
	console.DispayResults(result)

}
