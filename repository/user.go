package repository

import (
	"Demo_mysql/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(model.User) (int64, error)
	GetAllUser() ([]model.User, error)
	GetUserById(string) (model.User, error)
	EditeUser(string) (model.User, error)
	DeleteUser(string) error
}

type RepositoryMysql struct {
	ConnectionDB *sqlx.DB
}

func (repository RepositoryMysql) CreateUser(newUser model.User) (int64, error) {
	statement := `INSERT INTO user (first_name,last_name,addess,phone_number,created_time,updated_time) VALUES (?,?,?,?,?,?)`
	tx := repository.ConnectionDB.MustBegin()
	resual := tx.MustExec(statement, newUser.FristName, newUser.LastName, newUser.Address, newUser.PhoneNumber, newUser.CreatedTime, newUser.UpdatedTime)
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return resual.LastInsertId()
}
