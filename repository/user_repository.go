package repository

import "go-rest-api/model"

// interface, メソッド一覧
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
