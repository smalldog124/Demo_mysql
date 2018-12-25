package repository

import (
	"Demo_mysql/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(model.User) (int64, error)
	GetAllUser() ([]model.User, error)
	GetUserById(int) (model.User, error)
	EditeUser(int, model.User) (model.User, error)
	DeleteUser(int) error
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

func (repository RepositoryMysql) GetAllUser() ([]model.User, error) {
	var user []model.User
	statement := `SELECT user_id,first_name,last_name,addess,phone_number,created_time,updated_time From user`
	err := repository.ConnectionDB.Select(&user, statement)
	return user, err
}

func (repository RepositoryMysql) GetUserById(userID int) (model.User, error) {
	var user model.User
	statement := `SELECT user_id,first_name,last_name,addess,phone_number,created_time,updated_time From user WHERE user_id=?`
	err := repository.ConnectionDB.Get(&user, statement, userID)
	return user, err
}

func (repository RepositoryMysql) EditeUser(userID int, user model.User) (int64, error) {
	statement := `UPDATE user SET first_name=?,last_name=?,addess=?,phone_number=?,updated_time=? WHERE user_id=?`
	tx := repository.ConnectionDB.MustBegin()
	resual := tx.MustExec(statement, user.FristName, user.LastName, user.Address, user.PhoneNumber, user.UpdatedTime, userID)
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return resual.RowsAffected()
}

func (repository RepositoryMysql) DeleteUser(userID int) (int64, error) {
	statement := `DELETE FROM user WHERE user_id=?`
	tx := repository.ConnectionDB.MustBegin()
	resual := tx.MustExec(statement, userID)
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return resual.RowsAffected()
}
