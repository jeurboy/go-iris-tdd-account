package account

import (
	_ "github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"testing"
)

var accountCreator = AccountCreator{}

func TestCreate(t *testing.T) {
	account := accountCreator.New(100)

	assert := assert.New(t)
	assert.Equal(100.0, account.Total(), "Account has 100 Bath")
}

func TestCreateWith200(t *testing.T) {
	account := accountCreator.New(200)

	assert := assert.New(t)
	assert.Equal(200.0, account.Total(), "Account has 200 Bath")
}

func TestDeposit100ShouldGet200(t *testing.T) {
	account := accountCreator.New(100)
	account.Deposit(100)
	assert := assert.New(t)
	assert.Equal(200.0, account.Total(), "Account has 200 Bath")
}

func TestWithdraw50ShouldGet100(t *testing.T) {
	account := accountCreator.New(100)
	account.Withdraw(50)
	assert := assert.New(t)
	assert.Equal(50.0, account.Total(), "Account has 50 Bath")
}

func TestDepositFloatShouldGet200(t *testing.T) {
	account := accountCreator.New(100)
	account.Deposit(50.50)
	assert := assert.New(t)
	assert.Equal(150.50, account.Total(), "Account has 150.50 Bath")
}

func TestDepositFloatShouldGet(t *testing.T) {
	account := accountCreator.New(100)
	account.Deposits(50.50, 20, 30)

	assert := assert.New(t)
	assert.Equal(200.50, account.Total(), "Account has 200.50 Bath")

	account.Deposits(50.50)
	assert.Equal(251.0, account.Total(), "Account has 251 Bath")
}

func TestCreateWith200String(t *testing.T) {
	account := accountCreator.New("200")
	assert := assert.New(t)
	assert.Equal(200.0, account.Total(), "Account has 150.50 Bath")
}

func TestSplitOverLimitAccount(t *testing.T) {
	account := accountCreator.New(100.0)
	account1, account2 := account.Split(100.0, 50.0)

	assert := assert.New(t)
	assert.Equal(100.0, account1.Total(), "Account has 100 Bath")
	assert.Equal(0.0, account2.Total(), "Account has 0 Bath")
}

func TestCompareAccount(t *testing.T) {
	account1 := accountCreator.New(100.0)
	account2 := accountCreator.New(50.0)

	assert := assert.New(t)
	assert.True(account1.More(account2), "Account1 more than Account2")
	assert.True(account2.Less(account1), "Account2 less than Account1")
	assert.False(account2.More(account1), "Account1 more than Account2")
}

func TestTransferSuccessAccount(t *testing.T) {
	account1 := accountCreator.New(150.0)
	account2 := accountCreator.New(50.0)

	assert := assert.New(t)
	assert.True(account1.TransferTo(&account2, 50.0), "Account1 can transfer to Account2")
	assert.Equal(100.0, account1.Total(), "Account1 has 100 Bath")
	assert.Equal(100.0, account2.Total(), "Account2 has 100 Bath")
}

func TestSetAccountName(t *testing.T) {
	account := accountCreator.New(200.0)
	account.SetOwner("NEW")
	assert := assert.New(t)
	assert.Equal("NEW", account.GetOwner(), "Account owner is NEW")
}

func TestTransferLogAccount(t *testing.T) {
	account1 := accountCreator.New(150.0)
	account2 := accountCreator.New(50.0)

	account1.TransferTo(&account2, 70.0)

	assert := assert.New(t)
	Logs := account1.GetTransferLog()
	assert.Equal(70.0, Logs[0].Amount, "ok")
}

func TestTransferLogAccountWithFunk(t *testing.T) {
	account1 := accountCreator.New(150.0)
	account2 := accountCreator.New(50.0)

	account1.TransferTo(&account2, 70.0)
	account1.TransferTo(&account2, 20.0)
	account1.TransferTo(&account2, 10.0)

	assert := assert.New(t)
	Logs := account1.GetTransferLog()

	// Use funk mapping
	totals := funk.Map(Logs, func(x Log) float64 {
		return x.Amount
	})

	assert.Equal(100.0, funk.Sum(totals), "Total amount should be 100")
}
