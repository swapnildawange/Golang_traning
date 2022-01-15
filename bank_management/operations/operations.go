package operations

import (
	bankUtils "bank/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func generateID() string {
	return uuid.NewString()
}

func CreateAccount() (AccountInfo, string) {
	accID := generateID()
	accountHolderName := bankUtils.GetAccountHolderName()
	amount := 0.0
	return AccountInfo{
		AccID:             accID,
		AccountHolderName: accountHolderName,
		Amount:            amount,
	}, accID
}

func DeleteAccount(customers map[string]AccountInfo, accID string) map[string]AccountInfo {
	delete(customers, accID)
	return customers
}

func DepositMoney(customers map[string]AccountInfo, accID string, money float64) map[string]AccountInfo {
	latestAmount := money + customers[accID].Amount
	customers[accID] = AccountInfo{customers[accID].AccID, customers[accID].AccountHolderName, latestAmount}
	return customers
}

func WithDrawMoney(customers map[string]AccountInfo, accID string, money float64) map[string]AccountInfo {
	latestAmount := customers[accID].Amount - money
	customers[accID] = AccountInfo{customers[accID].AccID, customers[accID].AccountHolderName, latestAmount}
	return customers
}

func AddToHistory(accID string, amount float64, operation string) AccountHistory {
	return AccountHistory{
		AccID:     accID,
		Amount:    amount,
		Operation: operation,
		TimeStamp: time.Now().Format("01-02-2006 15:04:05 Monday"),
	}
}

func Gethistory(name string, accID string, total float64, history map[string][]AccountHistory) {
	fmt.Printf("Name : %v\nAccount ID : %v\nTotal : %v\n", name, accID, total)
	// customerDetails := history[accID]
	fmt.Printf("|%6v|%40v|%10v|%10v|%30v|\n", "Sr.No", "Account ID", "Operation", "Amount", "TimeStamp")
	for index, transactionDetails := range history[accID] {
		bankUtils.PrintHistory(index+1, name, accID, transactionDetails.Operation, transactionDetails.Amount, transactionDetails.TimeStamp)
	}
}
