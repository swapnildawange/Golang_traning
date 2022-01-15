package main

import (
	bankOperation "bank/operations"
	bankUtils "bank/utils"
	"fmt"
)

func main() {
	var customers map[string]bankOperation.AccountInfo = make(map[string]bankOperation.AccountInfo)

mainLoop:
	for {
		bankUtils.PrintOperations()
		selectedOperation := bankUtils.GetSelectedOperation()

		switch selectedOperation {
		case 1:
			accountInfo, accID := bankOperation.CreateAccount() // create account and return account info
			customers[accID] = accountInfo
			fmt.Println(accountInfo)
		case 2:
			accID := bankUtils.GetAccountID()
			customers = bankOperation.DeleteAccount(customers, accID)
		case 6:
			break mainLoop

		}
		fmt.Println(customers)
	}

}
