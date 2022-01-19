package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func validateOperationNumber(userInput string) (int, bool, string) {
	isValide := false
	err := ""
	val, stringErr := strconv.Atoi(userInput)
	if stringErr != nil {
		fmt.Printf("Supplied value %s is not a number \nTry again : ", userInput)
	} else {

		if val > 0 && val <= 6 {
			isValide = true
		} else {
			isValide = false
			err = "value out of range"
		}
	}

	return val, isValide, err
}

func validateName(name string) (bool, string) {

	match, _ := regexp.MatchString("[A-Za-z]", name)
	if match {
		return true, ""

	} else {
		return false, "Enter correct name"
	}
}

func validateWithDrawAmount(total float64,amount float64) (bool, string) {

	if amount > 0 && amount <= total {
		return true, ""
	} else {
		return false, "Enter correct amount"
	}
}


