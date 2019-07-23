package bank

import (
	_ "github.com/davecgh/go-spew/spew"
	"myapp/lib/account"
	_ "reflect"
)

type BankCreator struct{}
type Bank struct {
	id       int
	name     string
	accounts []account.Account
}

func (bankCreator BankCreator) New(id int) Bank {
	var bank Bank
	bank.id = id
	return bank
}

func (bank *Bank) SetName(name string) {
	bank.name = name
}
func (bank *Bank) GetName() string {
	return bank.name
}

func (bank *Bank) AddAccount(account1 account.Account) {
	bank.accounts = append(bank.accounts, account1)
}
