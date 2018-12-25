package repository_test

import (
	"Demo_mysql/model"
	"Demo_mysql/repository"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

const (
	username = "root"
	password = "sckshuhari"
	host     = "localhost"
	port     = "3360"
	database = "smalldogShop"
)

var urlSql = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
var db, _ = sqlx.Connect("mysql", urlSql)
var userRepository = repository.RepositoryMysql{
	ConnectionDB: db,
}

func Test_CreateUser_Input_User_Shoul_Be_NewUser(t *testing.T) {
	defer userRepository.ConnectionDB.Close()
	fixedTime, _ := time.Parse("2006-01-02", "2018-08-25")
	newUser := model.User{
		FristName:   "Smalldog",
		LastName:    "Adison",
		Address:     "123 californear",
		PhoneNumber: "092-3994-212",
		CreatedTime: fixedTime,
		UpdatedTime: fixedTime,
	}

	_, err := userRepository.CreateUser(newUser)

	assert.Equal(t, nil, err)
}

func Test_GetAllUser_Should_Be_Array_User(t *testing.T) {
	defer userRepository.ConnectionDB.Close()
	fixedTime, _ := time.Parse("2006-01-02", "2018-08-25")
	newUser := []model.User{
		{
			UserID:      "7",
			FristName:   "Smalldog",
			LastName:    "Adison",
			Address:     "123 californear",
			PhoneNumber: "092-3994-212",
			CreatedTime: fixedTime,
			UpdatedTime: fixedTime,
		},
	}

	users, err := userRepository.GetAllUser()

	assert.Equal(t, nil, err)
	assert.Equal(t, newUser, users)
}

func Test_GetUserByID_Input_ID_7_Should_Be_User_Smalldog(t *testing.T) {
	defer userRepository.ConnectionDB.Close()
	fixedTime, _ := time.Parse("2006-01-02", "2018-08-25")
	expectedUser := model.User{
		UserID:      "7",
		FristName:   "Smalldog",
		LastName:    "Adison",
		Address:     "123 californear",
		PhoneNumber: "092-3994-212",
		CreatedTime: fixedTime,
		UpdatedTime: fixedTime,
	}

	userID := 8
	users, err := userRepository.GetUserById(userID)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedUser, users)
}

func Test_EditeUser_Input_User_Name_Jone_Should_Be_User_Jone(t *testing.T) {
	defer userRepository.ConnectionDB.Close()
	fixedTime, _ := time.Parse("2006-01-02", "2018-12-24")
	user := model.User{
		FristName:   "Jone",
		LastName:    "Adison",
		Address:     "123 californear",
		PhoneNumber: "092-3994-212",
		UpdatedTime: fixedTime,
	}
	userID := 7

	actualUser, err := userRepository.EditeUser(userID, user)

	assert.Equal(t, nil, err)
	assert.True(t, actualUser == 1 || actualUser == 0)
}

func Test_DeleteUser_User_ID_7_Should_Be_Deleted(t *testing.T) {
	defer userRepository.ConnectionDB.Close()
	userID := 7

	actualUser, err := userRepository.DeleteUser(userID)

	assert.Equal(t, nil, err)
	assert.True(t, actualUser == 1 || actualUser == 0)
}
