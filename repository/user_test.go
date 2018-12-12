package repository_test

import (
	"Demo_mysql/model"
	"Demo_mysql/repository"
	"fmt"
	"log"
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

func Test_CreateUser_Input_User_Shoul_Be_NewUser(t *testing.T) {
	urlSql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	db, err := sqlx.Connect("mysql", urlSql)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connect sucessfull")
	defer db.Close()
	userRepository := repository.RepositoryMysql{
		ConnectionDB: db,
	}

	fixedTime, _ := time.Parse("2006-Jan-22", "2018-Oct-25")
	newUser := model.User{
		FristName:   "Smalldog",
		LastName:    "Adison",
		Address:     "123 californear",
		PhoneNumber: "092-3994-212",
		CreatedTime: fixedTime,
		UpdatedTime: fixedTime,
	}

	lastedUser, err := userRepository.CreateUser(newUser)

	assert.Equal(t, nil, err)
	assert.Equal(t, newUser, lastedUser)
}
