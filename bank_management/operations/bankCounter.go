package operations

import "time"

type bankCounterInterface interface {

	// Methods
	CreateAccount()
	DeleteAccount()
	DepositMoney()
	WithDrawMoney()
	Gethistory()
}

type AccountInfo struct {
	AccID  string
	AccountHolderName   string
	Amount float64
}

type typetime time.Time
type AccountHistory struct {
	AccID  string
	Amount float64
	Operation string
	TimeStamp string
}