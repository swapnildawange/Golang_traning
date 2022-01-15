package utils

import (
	"fmt"
)

func GetSelectedOperation() int {
	var userInput string
	for {
		fmt.Scanf("%v", &userInput)
		val, isValide, err := validateOperationNumber(userInput)
		if isValide {
			return val
		} else {
			fmt.Println(err)
		}
	}

}

func GetAccountHolderName() string {
	var name string
	for {
		fmt.Println("Enter your name :")
		fmt.Scanf("%v", &name)
		isValid, err := validateName(name)
		if isValid {
			return name
		} else {
			fmt.Println(err)
		}
	}
}

func GetAccountID() string {
	var accID string
	fmt.Println("Enter account id : ")
	fmt.Scanf("%v", &accID)
	return accID
}
