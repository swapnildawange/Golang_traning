package operations

type bankCounterInterface interface {

	// Methods
	CreateAccount()
	DeleteAccount()
	DepositMoney()
	WithDrawMoney()
	Gethistory()
}

type AccountInfo struct {
	accID  string
	accountHolderName   string
	amount float64
}
