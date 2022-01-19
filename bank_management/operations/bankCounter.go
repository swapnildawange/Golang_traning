package operations

import (
	"fmt"
	"time"
)

type BankServices interface {

	// Methods
	Create(name string) string
	Delete(accID string)
	Deposit(accID string, amount float64)
	Withdraw(accID string, amount float64)
	History()
	CreateTransaction(accID string, amount float64, operation string) string
}

type AccountInfo struct {
	AccID             string
	AccountHolderName string
	Amount            float64
}

type TransactionDetails struct {
	TransactionID string
	Amount        float64
	Operation     string
	TimeStamp     string
}

func (account *AccountInfo) Create(name string) string {
	account.AccID = generateID()
	account.AccountHolderName = name
	account.Amount = 0.0
	fmt.Println("Account created successfully")
	return account.AccID
}



func (account *AccountInfo) Deposit(accID string, amount float64) {
	account.AccID = accID
	// account.AccountHolderName = accountHolders[accID].AccountHolderName
	latestAmount := amount + account.Amount
	account.Amount = latestAmount
	fmt.Printf("Amount deposited successfully. Current balance is %v\n", account.Amount)
}
func (account *AccountInfo) Withdraw(accID string, amount float64) {
	latestAmount := account.Amount - amount
	account.AccID = accID
	// account.AccountHolderName = accountHolders[accID].AccountHolderName
	account.Amount = latestAmount
	fmt.Printf("Amount withdrawn successfully. Current balance is %v\n", account.Amount)
}

func (transaction *TransactionDetails) CreateTransaction(accID string, amount float64, operation string) string {

	transaction.TransactionID = generateID()
	transaction.Amount = amount
	transaction.Operation = operation
	transaction.TimeStamp = time.Now().Format("01-02-2006 15:04:05 Monday")
	return accID
}
