package repository

import (
	"Demo_mysql/model"
)

type UserRepository interface {
	CreateUser(model.User) (model.User, error)
}
