package main

import (
	bankOperation "bank/operations"
	bankUtils "bank/utils"
	"fmt"
)

func main() {

mainLoop:

	for {
		bankUtils.PrintOperations()
		selectedOperation := bankUtils.GetSelectedOperation()

		switch selectedOperation {
		case 1:
			// create account
			name, accID := bankOperation.CreateAccount() // create account and return account info
			fmt.Printf("Name : %v\nAccount ID : %v\n", name, accID)

		case 2:
			// delete account
			bankOperation.DeleteAccount()

		case 3:
			bankOperation.DepositMoney()

		case 4:
			bankOperation.WithDrawMoney()

		case 5:
			bankOperation.GetHistory()

		case 6:
			break mainLoop

		}
	}

}
