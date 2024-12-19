package main

import (
	"for_git/internal/console"
	"for_git/pkg/calculations"
	"for_git/pkg/logger"

	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	log := logger.New()
	consoleHandler := console.New(log)
	calculator := calculations.NewCalcualator(log)

	if len(os.Args) > 1 && os.Args[1] == "web" {
		// Start web server
		startWebServer(consoleHandler, calculator, log)
	} else {
		// Start CLI mode
		startCLI(consoleHandler, calculator, log)
	}
}
func startCLI(consoleHandler *console.Console, calculator *calculations.Calculator, log *logrus.Logger) {
	log.Info("Starting calculator in CLI mode")

	a, b, operator, err := consoleHandler.GetInput("", "", "")
	if err != nil {
		consoleHandler.DisplayError(err)
		return
	}

	result, err := calculator.Calculate(a, b, operator)
	if err != nil {
		consoleHandler.DisplayError(err)
		return
	}
	consoleHandler.DispayResults(result)
}
func startWebServer(consoleHandler *console.Console, calculator *calculations.Calculator, log *logrus.Logger) {
	log.Info("Starting calculator in web mode")

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET requests are supported", http.StatusMethodNotAllowed)
			return
		}

		// Parse query parameters
		aStr := r.URL.Query().Get("a")
		bStr := r.URL.Query().Get("b")
		operator := r.URL.Query().Get("operator")
		if operator == " " || operator == "+" {
			operator = "+"
		}

		a, b, operator, err := consoleHandler.GetInput(aStr, bStr, operator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := calculator.Calculate(a, b, operator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]any{
			"result": result,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Starting web server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
