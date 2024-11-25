package main

import (
	"for_git/internal/console"
	"for_git/pkg/calculations"
	"for_git/pkg/logger"
)

func main() {
	appLogger := logger.NewStandardLogger()

	inputLogger := console.NewInputLogger(appLogger)

	inputLogger.GetInput()

	a, b, operator, err := console.GetInput()
	if err != nil {
		console.DisplayError(err)

		return
	}
	result, err := calculations.Calculate(a, b, operator)
	if err != nil {
		console.DisplayError(err)

		return
	}
	console.DispayResults(result)
	appLogger.Info("The application is completed")

}
