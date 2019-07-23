package user

import (
	"strings"
	"time"
)

type User struct {
	ID    int    `gorm:"not_null;primary_key;column:id"`
	FName string `gorm:"column:firstname" validate:"required"`
	LName string `gorm:"column:lastname" validate:"required"`

	CreatedAt time.Time  `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime"`
}

func (User) TableName() string {
	return "user"
}

type UserOption func(user User) (User, error)

func withLower() UserOption {
	return func(user User) (User, error) {
		return User{
			FName: strings.ToLower(user.FName),
			LName: strings.ToLower(user.LName),
		}, nil
	}
}

func withAbvr(num int) UserOption {
	return func(user User) (User, error) {
		var result string
		l := strings.Split(user.GetName(), "")
		for i := 0; i < num; i++ {
			result += l[i]
		}
		return User{}, nil
	}
}

func (user User) GetName(userOption ...UserOption) string {
	for _, f := range userOption {
		user, _ = f(user)
	}
	return user.FName + " " + user.LName
}
