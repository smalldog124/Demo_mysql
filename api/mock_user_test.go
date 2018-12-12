package api_test

import (
	"Demo_mysql/model"
	"time"
)

type mockUserRepository struct{}

func (mockUser mockUserRepository) CreateUser(model.User) (int64, error) {
	return 1, nil
}
func (mockUser mockUserRepository) GetAllUser() ([]model.User, error) {
	fixedTime, _ := time.Parse("2006-Jan-22", "2018-Oct-25")
	return []model.User{
		{
			UserID:      "1",
			FristName:   "Smalldog",
			LastName:    "Adison",
			Address:     "123 californear",
			PhoneNumber: "092-3994-212",
			CreatedTime: fixedTime,
			UpdatedTime: fixedTime,
		},
	}, nil
}

func (mockUser mockUserRepository) GetUserById(userID string) (model.User, error) {
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

func (mockUser mockUserRepository) EditeUser(string) (model.User, error) {
	fixedTime, _ := time.Parse("2006-Jan-22", "2018-Oct-25")
	return model.User{
		UserID:      "1",
		FristName:   "Lek",
		LastName:    "Adison",
		Address:     "123 californear",
		PhoneNumber: "098-4772-211",
		CreatedTime: fixedTime,
		UpdatedTime: fixedTime,
	}, nil
}

func (mockUser mockUserRepository) DeleteUser(string) error {
	return nil
}
