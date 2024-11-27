package main

import (
	"for_git/internal/console"
	"for_git/pkg/calculations"
	"for_git/pkg/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	logger.Init()
	log := logger.Log

	log.Info("Need data from user")

	a, b, operator, err := console.GetInput()
	if err != nil {
		log.WithError(err).Error("Error with data input")
		return
	}

	log.WithFields(logrus.Fields{
		"a":        a,
		"b":        b,
		"operator": operator,
	}).Info("Start to calculate")

	result, err := calculations.Calculate(a, b, operator)
	if err != nil {
		log.WithError(err).Error("Error with calculation")
		return
	}
	log.WithField("result", result).Info("Calculation completed successfuly")
	console.DispayResults(result)

}
