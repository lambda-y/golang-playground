package main

import (
	"fmt"

	debt "bank_account/src/bank_accounts"
	investment "bank_account/src/bank_accounts"
	account "bank_account/src/bank_accounts/account"
)

func main() {
	fmt.Println("Hi")
	bank_Account := account.Account{Name: "BankAccount", Age: 26, Balance: 1000.0}
	new_debt_account := debt.Debt_Account{BankAccount: bank_Account, InterestRate: 0.05}
	new_investment_account := investment.Investment_Account{BankAccount: bank_Account, InterestRate: 0.05}

	fmt.Println(new_debt_account.CalculateBalance(6, 200.0))
	fmt.Println(new_investment_account.CalculateBalance(6, 200.0))
}
