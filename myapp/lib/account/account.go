package account

import (
	"errors"
	"fmt"
	"myapp/lib/user"
	"strconv"

	"time"
)

type AccountCreator struct{}
type Account struct {
	ID      int        `gorm:"primary_key;column:id"`
	Money   float64    `json:"Money" gorm:"column:account_money" validate:"required"`
	Name    string     `gorm:"column:account_name"`
	Logs    []Log      `gorm:"-"`
	OwnerID int        `gorm:"column:owner_id"`
	Owner   *user.User `gorm:"foreignkey:id;association_foreignkey:owner_id"`

	CreatedAt time.Time  `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime"`
}
type Log struct {
	ToAccountName string
	Amount        float64
	Total         float64
}

func (account *Account) Total() float64 {
	return account.Money
}

func (account *Account) SetAccountOwner(user *user.User) {
	account.Owner = user
}

func (account *Account) Split(amount1 float64, amount2 float64) (acc1 Account, acc2 Account) {
	var accountCreator1, accountCreator2 AccountCreator

	if account.Money >= amount1 {
		acc1 = accountCreator1.New(amount1)
		account.Withdraw(amount1)
	}
	if account.Money >= amount2 {
		acc2 = accountCreator2.New(amount2)
		account.Withdraw(amount2)
	}
	return acc1, acc2
}

func (account *Account) More(account2 Account) bool {
	if account.Total() > account2.Total() {
		return true
	}
	return false
}
func (account *Account) Less(account2 Account) bool {
	if account.Total() < account2.Total() {
		return true
	}
	return false
}

func (account *Account) TransferTo(account2 *Account, amount float64) bool {
	if account.Total() < amount {
		return false
	}
	account.Withdraw(amount)
	account2.Deposit(amount)
	account.Logs = append(account.Logs, Log{
		ToAccountName: account2.GetOwner(),
		Amount:        amount,
		Total:         account.Total(),
	})
	return true
}

func (account *Account) SetOwner(Name string) {
	account.Name = Name
}
func (account *Account) GetOwner() string {
	return account.Name
}
func (account *Account) GetOwnerUsername() string {
	return account.Owner.GetName()
}
func (account *Account) GetSummary() string {
	return fmt.Sprintf(account.Owner.GetName()+" has %.2f Baht", account.Total())
}

func (account *Account) GetTransferLog() []Log {
	return account.Logs
}

func (account *Account) Deposit(amount float64) {
	account.Money += amount
}

func (account *Account) Deposits(amounts ...interface{}) {
	for _, n := range amounts {
		if a, err := ConvertToFloat(n); err == nil {
			account.Deposit(a)
		} else {
			fmt.Println("On error", amounts, err.Error())
		}
	}
}

func (account *Account) Withdraw(amount float64) {
	account.Money -= amount
}

func (acc AccountCreator) New(amounts interface{}) Account {
	var account Account
	account.Deposits(amounts)
	return account
}

func ConvertToFloat(amount interface{}) (float64, error) {
	var result float64
	switch i := amount.(type) {
	case int:
		result = float64(i)
	case string:
		result, _ = strconv.ParseFloat(i, 64)
	case float64:
		result = i
	default:
		return 0.0, errors.New("error invalid type")
	}
	return result, nil
}
