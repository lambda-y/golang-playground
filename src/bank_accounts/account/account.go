package account

type Account struct {
	Name    string
	Age     int
	Balance float64
}

func (acct Account) GetName() string {
	return acct.Name
}

func (acct Account) GetAge() int {
	return acct.Age
}

func (acct Account) GetBalance() float64 {
	return acct.Balance
}
