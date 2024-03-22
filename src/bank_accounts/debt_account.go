package bank_accounts

import (
	"bank_account/src/bank_accounts/account"
)

type Debt_Account struct {
	BankAccount  account.Account
	InterestRate float64
}

func (debt Debt_Account) CalculateBalance(months int, monthly_payment float64) float64 {
	balance := debt.BankAccount.GetBalance()

	for i := 0; i < months; i++ {
		balance += balance*(debt.InterestRate/12) - monthly_payment
	}

	return balance
}
