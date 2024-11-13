//main/main.go
package main

import(
	"for_git/business"
	presentation
	"for_git/presentation"
)
func main(){
	a,b, operator, err := presentation.getInput
	if err != nil {
		presentation.displayError(err)
		return
	}
	result, err := business.calculate(a, b, operator)
	if err != nil {
		presentation.displayError(err)
		return
	}
	presentation.dispayResults(result)
}