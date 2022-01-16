package utils

import (
	"fmt"
)

func PrintOperations() {
	fmt.Printf("1. Create an account\n2. Delete an account\n3. Deposit money\n4. Withdraw money\n5. Get account history\n6. Exit the application\nEnter your choice: ")
}

func PrintHistory(serialNumber int,  transactionID string, typeOfOperation string, amount float64, timeStamp string) {

	fmt.Printf("|%6v|%40v|%10v|%10v|%30v|\n", serialNumber, transactionID, typeOfOperation, amount, timeStamp)

}
