package bank_accounts

import (
	"bank_account/src/bank_accounts/account"
)

type Investment_Account struct {
	BankAccount  account.Account
	InterestRate float64
}

func (investment Investment_Account) CalculateBalance(months int, monthly_payment float64) float64 {
	balance := investment.BankAccount.GetBalance()

	for i := 0; i < months; i++ {
		balance += balance*(investment.InterestRate/12) + monthly_payment
	}

	return balance
}
