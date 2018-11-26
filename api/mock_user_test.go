package api_test

import (
	"Demo_mysql/model"
	"time"
)

type mockUserRepository struct{}

func (mockUser mockUserRepository) CreateUser(model.User) (model.User, error) {
	fixedTime, _ := time.Parse("2006-Jan-22", "2018-Oct-25")
	return model.User{
		UserID:      "1",
		FristName:   "Smalldog",
		LastName:    "Adison",
		Address:     "123 californear",
		PhoneNumber: "092-3994-212",
		CreatedTime: fixedTime,
		UpdatedTime: fixedTime,
	}, nil
}
