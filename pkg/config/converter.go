package config

import "github.com/Anand078/moneytransfer/internal/model"

func ConvertToAccounts(balances []InitialBalance) []model.Account {
	var accounts []model.Account
	for _, b := range balances {
		accounts = append(accounts, model.Account{
			Name:    b.Name,
			Balance: b.Balance,
		})
	}
	return accounts
}
