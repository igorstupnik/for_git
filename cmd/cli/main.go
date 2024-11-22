package main

import (
	"for_git/internal/console"
	"for_git/pkg/calculations"
	"for_git/pkg/logger"
)

func main() {
	logger.Info("Calculator is running")
	a, b, operator, err := console.GetInput()
	if err != nil {
		console.DisplayError(err)
		logger.Error("Input error")
		return
	}
	result, err := calculations.Calculate(a, b, operator)
	if err != nil {
		console.DisplayError(err)
		logger.Error("Output error")
		return
	}
	console.DispayResults(result)
	logger.Info("Calculation completed")
}
