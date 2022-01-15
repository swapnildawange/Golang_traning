package operations

import (
	bankUtils "bank/utils"

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
		accID:             accID,
		accountHolderName: accountHolderName,
		amount:            amount,
	}, accID
}

func DeleteAccount(customers map[string]AccountInfo, accID string) map[string]AccountInfo {
	delete(customers, accID)
	return customers
}

func DepositMoney() {

}

func WithdrawMoney() {

}

func GetHistory() {

}
