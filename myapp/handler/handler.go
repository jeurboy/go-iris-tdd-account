package handler

import (
	"myapp/lib/account"
	"myapp/lib/user"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

var db *gorm.DB
var err error

// WithDB - use to auto attach db to current session
func WithDB(inDb *gorm.DB, f func(ctx iris.Context)) func(ctx iris.Context) {
	db = inDb
	return f
}

// GetUserByID - use for query user by id
func GetUserByID(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	var userModel user.User

	if userModel, err = getUserByID(db, id); err != nil {
		ctx.JSON(iris.Map{
			"message": "ID Undifined",
		})
	}

	ctx.JSON(iris.Map{
		"ID":    userModel.ID,
		"FName": userModel.FName,
		"LName": userModel.LName,
	})

	return
}

// PostUser - use for creating user
func PostUser(ctx iris.Context) {
	var user user.User
	ctx.ReadJSON(&user)

	if err := validator.New().Struct(user); err != nil {
		ctx.JSON(iris.Map{
			"message": "Error :" + err.Error(),
		})
		return
	}

	user.ID = getLatestUserID(db) + 1

	db.Create(&user)

	ctx.JSON(iris.Map{
		"status": "success",
		"ID":     user.ID,
	})
}

// PostAccount - use for creating user account
func PostAccount(ctx iris.Context) {
	var account account.Account
	ctx.ReadJSON(&account)

	if err := validator.New().Struct(account); err != nil {
		ctx.JSON(iris.Map{
			"message": "Error :" + err.Error(),
		})
		return
	}

	IDAccount := getLatestAccontID(db)
	account.ID = IDAccount + 1
	db.Create(&account)

	ctx.JSON(iris.Map{
		"status": "success",
		"ID":     account.ID,
	})
}

// MergeUserToAccount - use for set user into account
func MergeUserToAccount(ctx iris.Context) {
	accountID, _ := ctx.Params().GetInt("account_id")
	userID, _ := ctx.Params().GetInt("user_id")

	User, err := getUserByID(db, userID)
	if err != nil || User.ID == 0 {
		ctx.JSON(iris.Map{
			"message": "Error : Invalid User",
		})
		return
	}
	Account, err := getAccountByID(db, accountID)
	if err != nil || Account.ID == 0 {
		ctx.JSON(iris.Map{
			"message": "Error : Invalid Account",
		})
		return
	}
	Account.OwnerID = User.ID
	Account.Owner = &User

	db.Save(&Account)

	ctx.JSON(iris.Map{
		"status": "success",
	})
}
