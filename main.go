// main/main.go
package main

import (
	"for_git/business"
	"for_git/presentation"
)

func main() {
	a, b, operator, err := presentation.GetInput()
	if err != nil {
		presentation.DisplayError(err)
		return
	}
	result, err := business.Calculate(a, b, operator)
	if err != nil {
		presentation.DisplayError(err)
		return
	}
	presentation.DispayResults(result)
}
