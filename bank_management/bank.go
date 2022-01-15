package main

import (
	bankOperation "bank/operations"
	bankUtils "bank/utils"
	"fmt"
)

func main() {
	var customers map[string]bankOperation.AccountInfo = make(map[string]bankOperation.AccountInfo)
	var history map[string][]bankOperation.AccountHistory = make(map[string][]bankOperation.AccountHistory)

mainLoop:

	for {
		bankUtils.PrintOperations()
		selectedOperation := bankUtils.GetSelectedOperation()

		switch selectedOperation {
		// create account
		case 1:
			accountInfo, accID := bankOperation.CreateAccount() // create account and return account info
			customers[accID] = accountInfo
			fmt.Printf("Name : %v\nAccount ID : %v\n",accountInfo.AccountHolderName, accID)
			// delete account
		case 2:
			accID := bankUtils.GetAccountID()
			customers = bankOperation.DeleteAccount(customers, accID)
			// deposit money
		case 3:
			accID := bankUtils.GetAccountID()
			money := bankUtils.GetMoney()
			customers = bankOperation.DepositMoney(customers, accID, money)
			// add to history
			history[accID] = append(history[accID], bankOperation.AddToHistory(accID, money, "deposit"))
			// withdraw money
		case 4:
			accID := bankUtils.GetAccountID()
			for {
				money := bankUtils.GetMoneyToWithDraw()
				if money <= customers[accID].Amount {
					customers = bankOperation.WithDrawMoney(customers, accID, money)
					// add to history
					history[accID] = append(history[accID], bankOperation.AddToHistory(accID, money, "withdraw"))
					break
				} else {
					fmt.Println("Insufficient balance")
				}
			}
		case 5:
			accID := bankUtils.GetAccountID()
			name := customers[accID].AccountHolderName
			total:=customers[accID].Amount
			bankOperation.Gethistory(name,accID,total, history)
		case 6:
			break mainLoop

		}
	}

}
