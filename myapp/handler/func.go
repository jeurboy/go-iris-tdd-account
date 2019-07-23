package handler

import (
	"myapp/lib/account"
	"myapp/lib/user"

	"github.com/jinzhu/gorm"
)

func getLatestUserID(db *gorm.DB) int {
	var user user.User
	db.Last(&user)

	return int(user.ID)
}

func getLatestAccontID(db *gorm.DB) int {
	var account account.Account
	db.Last(&account)

	return int(account.ID)
}

func getUserByID(db *gorm.DB, id int) (user.User, error) {
	var users user.User

	db.Where("id = ?", id).First(&users)
	return users, nil
}

func getAccountByID(db *gorm.DB, id int) (account.Account, error) {
	var account account.Account

	db.Where("id = ?", id).First(&account)
	return account, nil
}
