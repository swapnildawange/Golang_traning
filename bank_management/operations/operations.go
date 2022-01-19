package operations

import (
	bankUtils "bank/utils"
	"fmt"

	"github.com/google/uuid"
)

var accountHolders map[string]*AccountInfo = make(map[string]*AccountInfo)
var accountHistory map[string][]*TransactionDetails = make(map[string][]*TransactionDetails)

func generateID() string {
	return uuid.NewString()
}

func CreateAccount() (string, string) {
	accountHolderName := bankUtils.GetAccountHolderName()
	var account AccountInfo
	accID := account.Create(accountHolderName)
	accountHolders[accID] = &account

	return accountHolderName, accID
}

func DeleteAccount() {
	accID := GetAccountID()
	// var account AccountInfo
	// account.Delete(accID)
	delete(accountHolders, accID)
	fmt.Println("Account deleted successfully")
}

func DepositMoney() {
	accID := GetAccountID()
	amount := bankUtils.GetAmount()
	var account *AccountInfo = accountHolders[accID]
	account.Deposit(accID, amount)
	AddToHistory(accID, amount, "Deposit")

}

func WithDrawMoney() {
	accID := GetAccountID()
	var account *AccountInfo = accountHolders[accID]
	amount := bankUtils.GetWithDrawAmount(accountHolders[accID].Amount)
	account.Withdraw(accID, amount)
	AddToHistory(accID, amount, "Withdraw")
}

func AddToHistory(accID string, amount float64, operationType string) {
	var transaction TransactionDetails
	transaction.CreateTransaction(accID, amount, operationType)
	accountHistory[accID] = append(accountHistory[accID], &transaction)
}

func GetHistory() {
	accID := GetAccountID()
	fmt.Println("------------------------------------------------------------------------------------------------------")
	fmt.Printf("Account ID : %v\nName : %v\nTotal : %v\n", accID, accountHolders[accID].AccountHolderName, accountHolders[accID].Amount)
	fmt.Println("------------------------------------------------------------------------------------------------------")
	fmt.Printf("|%6v|%40v|%10v|%10v|%30v|\n", "Sr.No", "Transaction ID", "Operation", "Amount", "TimeStamp")
	for index, transactionDetails := range accountHistory[accID] {
		bankUtils.PrintHistory(index+1, transactionDetails.TransactionID, transactionDetails.Operation, transactionDetails.Amount, transactionDetails.TimeStamp)
	}
	fmt.Println("------------------------------------------------------------------------------------------------------")
}

func GetAccountID() string {

	var accID string
	for {

		fmt.Println("Enter account id : ")
		fmt.Scanf("%v", &accID)

		_, ok := accountHolders[accID]
		if ok {
			return accID

		} else {
			fmt.Println("Account id does not exist")
		}
	}

}
