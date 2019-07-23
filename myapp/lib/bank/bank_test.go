package bank

import (
	_ "github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	_ "github.com/thoas/go-funk"
	"myapp/lib/account"
	"myapp/lib/user"
	"testing"
)

var bankCreator = BankCreator{}
var accountCreator = account.AccountCreator{}

func TestCreateBank(t *testing.T) {
	bank := bankCreator.New(1)
	bank.SetName("SCB")

	assert := assert.New(t)
	assert.Equal("SCB", bank.GetName(), "Bank name is SCB")
}

func TestAddAccountToBank(t *testing.T) {
	account1 := accountCreator.New(200)
	account2 := accountCreator.New(200)

	bank := bankCreator.New(1)
	bank.AddAccount(account1)
	bank.AddAccount(account2)

	assert := assert.New(t)
	assert.Equal(2, len(bank.accounts), "Bank has 2 account")
}

func TestAccountWithUser(t *testing.T) {
	user := &user.User{
		FName: "New",
		LName: "Knocking",
	}
	account := accountCreator.New(200)
	account.SetAccountOwner(user)

	assert := assert.New(t)
	assert.Equal("New Knocking", user.GetName(), "Account owner is New Knocking")
	assert.Equal("New Knocking", account.GetOwnerUsername(), "Account owner is New Knocking")
	assert.Equal("New Knocking has 200.00 Baht", account.GetSummary(), "Account owner is New Knocking")
}
