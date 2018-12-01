package repository

import (
	"Demo_mysql/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(model.User) (model.User, error)
	GetAllUser() ([]model.User, error)
	GetUserById(string) (model.User, error)
	EditeUser(string) (model.User, error)
	DeleteUser(string) error
}

type RepositoryMysql struct {
	ConnectionDB *sqlx.DB
}
