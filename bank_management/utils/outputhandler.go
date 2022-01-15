package utils

import (
	"fmt"
)

func PrintOperations() {
	fmt.Println("1. Create an account\n2. Delete an account\n3. Deposit money\n4. Withdraw money\n5. Get account history\n6. Exit the application")
}

func PrintHistory(serialNumber int, name string, accID string, typeOfOperation string, amount float64, timeStamp string) {

	fmt.Printf("|%6v|%40v|%10v|%10v|%30v|\n", serialNumber, accID, typeOfOperation, amount, timeStamp)

}
