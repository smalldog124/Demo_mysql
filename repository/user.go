package repository

import (
	"Demo_mysql/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(model.User) (model.User, error)
	GeatAllUser() ([]model.User, error)
}

type RepositoryMysql struct {
	ConnectionDB *sqlx.DB
}
