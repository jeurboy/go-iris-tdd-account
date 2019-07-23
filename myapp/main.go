package main

import (
	"myapp/handler"
	"myapp/lib/account"
	"myapp/lib/user"

	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
)

// "github.com/davecgh/go-spew/spew"

var db *gorm.DB
var err error

func main() {
	fmt.Println("Server was started")

	username := "account"
	password := "account"
	hostname := "go-db"
	port := "3306"
	dbname := "account"

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC", username, password, hostname, port, dbname)
	db, err = gorm.Open("mysql", connString)

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.LogMode(true)

	// Migrate the schema
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&account.Account{})

	app := iris.Default()

	spew.Dump(handler.GetUserByID)

	app.Get("/user/{id:int}", handler.WithDB(db, handler.GetUserByID))
	app.Post("/user", handler.WithDB(db, handler.PostUser))
	app.Post("/account", handler.WithDB(db, handler.PostAccount))
	app.Post("/account/{account_id:int}/user/{user_id:int}", handler.WithDB(db, handler.MergeUserToAccount))

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}
