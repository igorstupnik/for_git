package main

import (
	"for_git/internal/console"
	"for_git/pkg/calculations"
)

func main() {
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
}
